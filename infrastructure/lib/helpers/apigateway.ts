import { CfnOutput } from "aws-cdk-lib";
import { LambdaRestApi, LambdaRestApiProps } from "aws-cdk-lib/aws-apigateway";
import { Construct } from "constructs";
import { VirtualEnvironment } from "../virtual-environment";

export interface ProxyingApiGatewayProps extends LambdaRestApiProps {
    virtualEnv: VirtualEnvironment
}

export class ProxyingApiGateway extends LambdaRestApi {
    public readonly output: CfnOutput;

    constructor(scope: Construct, id: string, props: ProxyingApiGatewayProps) {
        super(scope, id, {
            ...props,
            description: props.virtualEnv.name,
        });

        this.output = new CfnOutput(scope, `apigw.${id}.url`, {
            value: this.url
        });
    }
}