import { CfnOutput } from "aws-cdk-lib";
import { Queue, QueueEncryption, QueueProps } from "aws-cdk-lib/aws-sqs";
import { Construct } from "constructs";

export class SecureQueue extends Queue {
    public readonly outputArn: CfnOutput;
    public readonly outputName: CfnOutput;
    public readonly outputUrl: CfnOutput;

    constructor(scope: Construct, id: string, props?: QueueProps) {
        super(scope, id, {
            encryption: props?.encryption ?? QueueEncryption.KMS_MANAGED
        });

        this.outputArn = new CfnOutput(scope, `sqs.${id}.arn`, {
            value: this.queueArn
        });

        this.outputName = new CfnOutput(scope, `sqs.${id}.name`, {
            value: this.queueName
        });

        this.outputUrl = new CfnOutput(scope, `sqs.${id}.url`, {
            value: this.queueUrl
        });
    }
}
