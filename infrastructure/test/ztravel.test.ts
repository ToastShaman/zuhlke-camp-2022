import * as cdk from 'aws-cdk-lib';
import { Template } from 'aws-cdk-lib/assertions';
import { Runtime } from 'aws-cdk-lib/aws-lambda';
import { VirtualEnvironment } from '../lib/virtual-environment';
import { ZTravelStack } from '../lib/ztravel-stack';

describe("ZTravel Stack", () => {
    let app: cdk.App, stack: ZTravelStack, template: Template;
    const venv = VirtualEnvironment.CI

    beforeAll(() => {
        app = new cdk.App();

        stack = new ZTravelStack(app, `ztravel-${venv}`, {
            virtualEnv: venv,
            maintenanceModeEnabled: false
        });

        template = Template.fromStack(stack);
    });

    test("Creates Hello Function", () => {
        template.hasResource('AWS::Lambda::Function', {
            Properties: {
                Handler: 'hello',
                Runtime: Runtime.GO_1_X.name
            }
        });
    });

    test("Creates World Function", () => {
        template.hasResource('AWS::Lambda::Function', {
            Properties: {
                Handler: 'world',
                Runtime: Runtime.GO_1_X.name
            }
        });
    });

    test("Has all necessary infrastructure", () => {
        template.resourceCountIs('AWS::Lambda::Function', 3)
        template.resourceCountIs('AWS::ApiGateway::RestApi', 2)
    });

    test("Creates an API GW for the Hello API", () => {
        template.hasResource('AWS::ApiGateway::RestApi', {
            Properties: {
                Name: `hello-api`
            }
        });
    });

    test("Creates an API GW for the World API", () => {
        template.hasResource('AWS::ApiGateway::RestApi', {
            Properties: {
                Name: `world-api`
            }
        });
    });

    test("Creates S3 Bucket", () => {
        template.resourceCountIs('AWS::S3::Bucket', 1)
    });
});
