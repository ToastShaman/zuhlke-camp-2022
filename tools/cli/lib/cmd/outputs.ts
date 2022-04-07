import { Command } from "commander";
import { StackOutputsFor } from "../aws/cloudformation";
import { VirtualEnvironment } from "../virtual-environment";

export function NewOutputsCommand(venvs: VirtualEnvironment[]): Command {
    const outputs = new Command("outputs");

    venvs.forEach(venv => {
        outputs
            .command(venv.name)
            .action(NewOutputAction(venv));
    });

    return outputs
}

function NewOutputAction(venv: VirtualEnvironment): () => any {
    return async () => {
        const outputs = await StackOutputsFor(venv);

        console.table([
            { name: 'hello', endpoint: outputs.apigwhelloapiurl },
            { name: 'world', endpoint: outputs.apigwworldapiurl },
        ]);
    };
}