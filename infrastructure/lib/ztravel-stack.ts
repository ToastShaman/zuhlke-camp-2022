import { Stack, StackProps } from 'aws-cdk-lib';
import { Construct } from 'constructs';
import { VirtualEnvironment } from './virtual-environment';
import { GoFunction } from './helpers/go-function';
import { SecureBucket } from './helpers/s3';
import { ProxyingApiGateway } from './helpers/apigateway';
import { ApiKeySecret } from './helpers/secretsmanager';
import { SigningKey } from './helpers/kms';
import { SecureQueue } from './helpers/sqs';
import { SqsEventSource } from 'aws-cdk-lib/aws-lambda-event-sources';

export interface ZTravelStackProps extends StackProps {
  virtualEnv: VirtualEnvironment,
  maintenanceModeEnabled: boolean
}

export class ZTravelStack extends Stack {
  public readonly venv: VirtualEnvironment;

  public readonly submissions: SecureBucket;
  public readonly helloFn: GoFunction;
  public readonly worldFn: GoFunction;
  public readonly pigtailFn: GoFunction;
  public readonly helloApi: ProxyingApiGateway;
  public readonly worldApi: ProxyingApiGateway;
  public readonly apiKeySecret: ApiKeySecret;
  public readonly signingKey: SigningKey;
  public readonly pigtailQueue: SecureQueue;

  constructor(scope: Construct, id: string, props: ZTravelStackProps) {
    super(scope, id, props);

    this.venv = props.virtualEnv;

    this.apiKeySecret = new ApiKeySecret(this, 'api-key');

    this.signingKey = new SigningKey(this, 'signing-key');

    this.submissions = new SecureBucket(this, 'submissions');

    this.helloFn = new GoFunction(this, "hello-proc", {
      handler: 'hello',
    });

    this.worldFn = new GoFunction(this, "world-proc", {
      handler: 'world',
    });

    this.pigtailFn = new GoFunction(this, "pigtail-proc", {
      handler: 'pigtail',
    });

    this.pigtailQueue = new SecureQueue(this, "pigtail-queue");

    this.helloApi = new ProxyingApiGateway(this, `hello-api`, {
      handler: this.helloFn,
      virtualEnv: this.venv
    });

    this.worldApi = new ProxyingApiGateway(this, `world-api`, {
      handler: this.worldFn,
      virtualEnv: this.venv
    });

    [this.helloFn, this.worldFn].forEach(fn => {
      fn.addEnvironment("BUCKET_NAME", this.submissions.bucketName);
      fn.addEnvironment("API_KEY_SM_NAME", this.apiKeySecret.secretName);
      fn.addEnvironment("SIGNING_KEY_KMS_ID", this.signingKey.keyId);
      fn.addEnvironment("QUEUE_URL", this.pigtailQueue.queueUrl);
      fn.addEnvironment("MAINTENANCE_MODE_ENABLED", String(props.maintenanceModeEnabled))
      this.submissions.grantReadWrite(fn);
      this.apiKeySecret.grantRead(fn);
      this.signingKey.grant(fn, "kms:Sign");
      this.pigtailQueue.grantSendMessages(fn);
    });

    this.pigtailQueue.grantConsumeMessages(this.pigtailFn);

    this.pigtailFn.addEventSource(
      new SqsEventSource(this.pigtailQueue, { batchSize: 10 })
    );
  }
}
