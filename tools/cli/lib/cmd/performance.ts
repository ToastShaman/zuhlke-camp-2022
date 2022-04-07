import { Command } from "commander";
import { writeFileSync } from "fs";
import { VirtualEnvironment } from "../virtual-environment";
import { NewApiKey } from "./generate-api-key";
const { spawn } = require('child_process');

export function NewPerfCommand(venvs: VirtualEnvironment[]): Command {
    const perf = new Command("perf");

    venvs.forEach(venv => {
        perf
            .command(venv.name)
            .action(NewPerfAction(venv));
    });

    return perf
}

function NewPerfAction(venv: VirtualEnvironment): () => any {
    return async () => {
        const [outputs, token] = await NewApiKey(venv);

        writeFileSync('../performance/.env', `ZHOST=${outputs.apigwhelloapiurl}\nZTOKEN=${token}`);

        const child = spawn("docker run --rm -i --env-file=.env grafana/k6 run - <simple.js", [], {
            shell: true,
            cwd: '../performance',
            stdio: 'inherit'
        });

        process.on('SIGINT', () => {
            child.kill('SIGINT');
        })

        child.on('close', (code: any) => {
            console.log(`child process close all stdio with code ${code}`);
        });

        child.on('exit', (code: any) => {
            console.log(`child process exited with code ${code}`);
        });
    };
}