import { CfnOutput } from "aws-cdk-lib";
import { Key, KeySpec, KeyUsage } from "aws-cdk-lib/aws-kms";
import { Construct } from "constructs";

export class SigningKey extends Key {
    public readonly outputArn: CfnOutput;
    public readonly outputId: CfnOutput;

    constructor(scope: Construct, id: string) {
        super(scope, id, {
            enableKeyRotation: false,
            keySpec: KeySpec.ECC_NIST_P256,
            keyUsage: KeyUsage.SIGN_VERIFY
        });

        this.outputArn = new CfnOutput(scope, `kms.${id}.arn`, {
            value: this.keyArn
        });

        this.outputId = new CfnOutput(scope, `kms.${id}.id`, {
            value: this.keyId
        });
    }
}
