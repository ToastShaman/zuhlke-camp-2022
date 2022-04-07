import { Rule, Schedule } from "aws-cdk-lib/aws-events";
import { LambdaFunction } from "aws-cdk-lib/aws-events-targets";
import { IFunction } from "aws-cdk-lib/aws-lambda";
import { Construct } from "constructs";

export interface EventRuleProps {
    schedule: Schedule
    target: IFunction
    enabled: boolean
}

export class EventRule {
    public readonly schedule: Schedule;
    public readonly target: LambdaFunction;
    public readonly rule: Rule;

    constructor(scope: Construct, id: string, props: EventRuleProps) {
        this.schedule = props.schedule;
        this.target = new LambdaFunction(props.target);
        this.rule =  new Rule(scope,id, {
            enabled: props.enabled,
            schedule: this.schedule,
            targets: [this.target]
        });
    }
}
