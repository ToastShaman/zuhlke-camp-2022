import { CfnOutput, RemovalPolicy } from "aws-cdk-lib";
import { BlockPublicAccess, Bucket, BucketEncryption } from "aws-cdk-lib/aws-s3";
import { Construct } from "constructs";

export interface SecureBucketProps {
    removalPolicy?: RemovalPolicy
}

export class SecureBucket extends Bucket {
    public readonly output: CfnOutput;

    constructor(scope: Construct, id: string, props?: SecureBucketProps) {
        super(scope, id, {
            encryption: BucketEncryption.S3_MANAGED,
            enforceSSL: true,
            versioned: true,
            blockPublicAccess: BlockPublicAccess.BLOCK_ALL,
            removalPolicy: props?.removalPolicy ?? RemovalPolicy.DESTROY,
        });
        
        this.output = new CfnOutput(scope, `s3.${id}.name`, {
            value: this.bucketName
        });
    }
}