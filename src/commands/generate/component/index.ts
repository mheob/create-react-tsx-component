import { Command, Flags } from '@oclif/core';

import Page from '../page';

export default class extends Command {
  static description = 'Create a React Component';

  static examples = ['$ oex hello friend --from oclif'];

  static flags = {
    dest: Flags.string({ char: 'd', description: 'Where to create the file(s)?' }),
    story: Flags.boolean({ char: 's', description: 'Generate storybook story file?' }),
    test: Flags.boolean({ char: 't', description: 'Generate tests file?' }),
    'dry-run': Flags.boolean({ description: 'Dry run only?' }),
  };

  static args = [];

  async run(): Promise<void> {
    this.log('Hello from the Component! (./src/commands/component/index.ts)');
    Page.run();
  }
}
