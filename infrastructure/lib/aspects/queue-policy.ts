import { Annotations, IAspect, Tokenization } from "aws-cdk-lib";
import { CfnBucket } from "aws-cdk-lib/aws-s3";
import { CfnQueue } from "aws-cdk-lib/aws-sqs";
import { IConstruct } from "constructs";

export class QueuePolicyChecker implements IAspect {
    public visit(node: IConstruct): void {
        if (node instanceof CfnQueue) {
            if (!node.kmsMasterKeyId) {
                Annotations.of(node).addError('Queue encryption is not enabled');
            }
        }
    }
}
