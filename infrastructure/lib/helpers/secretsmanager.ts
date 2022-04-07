
import { CfnOutput } from "aws-cdk-lib";
import { Secret, SecretProps } from "aws-cdk-lib/aws-secretsmanager";
import { Construct } from "constructs";

export class ApiKeySecret extends Secret {
    public readonly output: CfnOutput;

    constructor(scope: Construct, id: string, props?: SecretProps) {
        super(scope, id, {
            description: "Secret for creating signed JWT tokens",
            generateSecretString: {
                passwordLength: 64
            }
        });

        this.output = new CfnOutput(scope, `sm.${id}.name`, {
            value: this.secretName
        });
    }
}