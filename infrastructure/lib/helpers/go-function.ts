import { CfnOutput, Duration } from "aws-cdk-lib";
import { Metric } from "aws-cdk-lib/aws-cloudwatch";
import { Runtime, Code, Function } from "aws-cdk-lib/aws-lambda";
import { Construct } from "constructs";
import { join } from "path";

export interface GoFunctionProps {
    handler: string
}

export class GoFunction extends Function {
    public readonly output: CfnOutput;
    public readonly failureRate: Metric;
    public readonly latency: Metric;
    public readonly incovations: Metric;

    constructor(scope: Construct, id: string, props: GoFunctionProps) {
        super(scope, id, {
            runtime: Runtime.GO_1_X,
            handler: props.handler,
            code: Code.fromAsset(join(__dirname, "../../../lambdas/ztravel"), {
                bundling: {
                    image: Runtime.GO_1_X.bundlingImage,
                    command: ['make', 'bundling'],
                    environment: {
                        'GOOS': 'linux',
                        'GOARCH': 'amd64'
                    },
                    user: 'root'
                }
            })
        });

        this.output = new CfnOutput(scope, `lambda.${id}.name`, {
            value: this.functionName
        });

        this.failureRate = this.metricErrors({
            statistic: 'avg',
            period: Duration.minutes(5)
        });

        this.latency = this.metricDuration({
            statistic: "p99.00",
            period: Duration.minutes(5)
        });

        this.incovations = this.metricInvocations({
            period: Duration.minutes(5)
        });
    }
};