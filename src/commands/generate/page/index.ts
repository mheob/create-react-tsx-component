import { Command } from '@oclif/core';

export default class extends Command {
  static description = 'Create a React Page';

  static examples = ['$ oex hello friend --from oclif'];

  static flags = {};

  static args = [];

  async run(): Promise<void> {
    this.log('Hello from the Page! (./src/commands/page/index.ts)');
  }
}
