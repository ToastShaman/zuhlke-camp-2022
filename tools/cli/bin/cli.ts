#!/usr/bin/env npx ts-node

import { Command } from "commander";
import { NewGenerateApiKeyCommand } from "../lib/cmd/generate-api-key";
import { NewOutputsCommand } from "../lib/cmd/outputs";
import { NewPerfCommand } from "../lib/cmd/performance";
import { NewPingCommand } from "../lib/cmd/ping";
import { VirtualEnvironment } from "../lib/virtual-environment";

(async function () {
  const program = new Command();

  program
    .command("generate")
    .addCommand(NewGenerateApiKeyCommand(VirtualEnvironment.DEV))

  program
    .addCommand(NewOutputsCommand(VirtualEnvironment.DEV))

  program
    .addCommand(NewPingCommand(VirtualEnvironment.DEV))

  program
    .addCommand(NewPerfCommand(VirtualEnvironment.DEV))

  program.parse();
})();
