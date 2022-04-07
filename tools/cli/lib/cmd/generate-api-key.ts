import { Command } from "commander";
import { sign } from "jsonwebtoken";
import { Outputs, StackOutputsFor } from "../aws/cloudformation";
import { GetSecretValue } from "../aws/secretsmanager";
import { VirtualEnvironment } from "../virtual-environment";

export function NewGenerateApiKeyCommand(venvs: VirtualEnvironment[]): Command {
    const apikey = new Command("apikey");

    venvs.forEach(venv => {
        apikey
            .command(venv.name)
            .action(NewGenerateApiKeyAction(venv));
    });

    return apikey
}

function NewGenerateApiKeyAction(venv: VirtualEnvironment): () => any {
    return async () => {
        const [outputs, token] = await NewApiKey(venv);
        console.table([
            { name: 'hello', endpoint: outputs.apigwhelloapiurl },
            { name: 'world', endpoint: outputs.apigwworldapiurl },
        ]);

        console.table([{ token }]);
    };
}

export async function NewApiKey(venv: VirtualEnvironment): Promise<[Outputs, string]> {
    const outputs = await StackOutputsFor(venv);
    const key = await GetSecretValue(outputs.smapikeyname);
    const token = sign({ ephemeral: true }, key, { expiresIn: '5h' });
    return [outputs, token];
}
