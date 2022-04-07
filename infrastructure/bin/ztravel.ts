#!/usr/bin/env node

import 'source-map-support/register';
import * as cdk from 'aws-cdk-lib';
import { ZTravelStack } from '../lib/ztravel-stack';
import { Aspects, CfnParameter, Tags } from 'aws-cdk-lib';
import { VirtualEnvironment } from '../lib/virtual-environment';
import { BucketPolicyChecker } from '../lib/aspects/bucket-policy';
import { QueuePolicyChecker } from '../lib/aspects/queue-policy';

const app = new cdk.App();

VirtualEnvironment.DEV
  .concat(VirtualEnvironment.PRE_PROD)
  .forEach(venv => {
    const id = venv.withPrefix("ztravel");

    const stack = new ZTravelStack(app, id.name, {
      virtualEnv: id,
      maintenanceModeEnabled: false,
      env: {
        account: process.env.CDK_DEPLOY_ACCOUNT || process.env.CDK_DEFAULT_ACCOUNT,
        region: process.env.CDK_DEPLOY_REGION || process.env.CDK_DEFAULT_REGION
      }
    });

    const owner = new CfnParameter(stack, "owner", {
      type: "String",
      minLength: 1,
      description: "The name of owner.",
    });

    Tags.of(stack).add('Owner', owner.valueAsString);
    Tags.of(stack).add('Environment', id.name);

    Aspects.of(stack).add(new BucketPolicyChecker());
    Aspects.of(stack).add(new QueuePolicyChecker());
  });
