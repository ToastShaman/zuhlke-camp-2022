import { Command } from "commander";
import { VirtualEnvironment } from "../virtual-environment";
import { NewApiKey } from "./generate-api-key";
import axios from 'axios';

export function NewPingCommand(venvs: VirtualEnvironment[]): Command {
    const ping = new Command("ping");

    venvs.forEach(venv => {
        ping
            .command(venv.name)
            .action(NewPingAction(venv));
    });

    return ping
}

function NewPingAction(venv: VirtualEnvironment): () => any {
    return async () => {
        const [outputs, token] = await NewApiKey(venv);
        const results: any[] = []

        for await (const url of [outputs.apigwhelloapiurl, outputs.apigwworldapiurl]) {
            try {
                const response = await axios.get(url, {
                    headers: {
                        Authorization: `Bearer ${token}`
                    }
                });
                results.push({ URL: url, Status: response.statusText });
            } catch (error) {
                if (axios.isAxiosError(error)) {
                    results.push({ URL: url, Status: error.response?.statusText ?? '-1' });
                } else {
                    results.push({ URL: url, Status: error });
                }
            }
        }

        console.table(results);
    };
}