import { CloudFormationClient, DescribeStacksCommand, ListStacksCommand, StackStatus } from "@aws-sdk/client-cloudformation";
import { VirtualEnvironment } from "../virtual-environment";

export interface Outputs {
    s3submissionsname: string
    apigwworldapiurl: string
    lambdaworldprocname: string
    smapikeyname: string
    lambdahelloprocname: string
    apigwhelloapiurl: string
}

export async function StackOutputsFor(venv: VirtualEnvironment): Promise<Outputs> {
    const client = new CloudFormationClient({ region: 'eu-west-2' });

    const stacks = await client.send(new ListStacksCommand({
        StackStatusFilter: [StackStatus.CREATE_COMPLETE, StackStatus.UPDATE_COMPLETE]
    }));

    const stack = stacks.StackSummaries?.find(s => s.StackName === `ztravel-${venv}`)
    if (!stack) {
        throw Error(`ztravel-${venv} could not be found`);
    }

    const desc = await client.send(new DescribeStacksCommand({
        StackName: stack.StackName
    }));
    if (!desc.Stacks) {
        throw Error(`ztravel-${venv} could not be described`);
    }

    const mapped = desc.Stacks[0]?.Outputs?.map(o => [o.OutputKey, o.OutputValue])!
    return Object.fromEntries(mapped);
}