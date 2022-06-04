import { Command } from '@oclif/core';
import { capitalCase } from 'change-case';
import type { ListQuestion } from 'inquirer';
import { prompt } from 'inquirer';

import Component from './component';
import Page from './page';

const variants = ['component', 'page', 'hook'] as const;
type Variants = typeof variants[number];

type VariantChoice = { name: string; value: Variants };

const variantChoices = (): VariantChoice[] =>
  variants.map((variant) => ({
    name: capitalCase(variant),
    value: variant,
  }));

const generateQuestion: ListQuestion<{ generate: Variants }> = {
  type: 'list',
  name: 'generate',
  message: 'What do you want to generate?',
  choices: variantChoices,
};

export default class Generate extends Command {
  static description = 'Generate a component of you choice';

  async run(): Promise<void> {
    const userResponse = await prompt([generateQuestion]);

    switch (userResponse.generate) {
      case 'component':
        Component.run();
        break;
      case 'page':
        Page.run();
        break;
      default:
        break;
    }
  }

  // static flags = {
  //   // flag with a value (-n, --name=VALUE)
  //   name: Flags.string({ char: 'n', description: 'name to print' }),
  //   // flag with no value (-f, --force)
  //   force: Flags.boolean({ char: 'f' }),
  // };
  //
  // static args = [{ name: 'file' }];
  //
  // public async run(): Promise<void> {
  //   const { args, flags } = await this.parse(Generate);
  //
  //   const name = flags.name ?? 'world';
  //   this.log(
  //     `hello ${name} from /Users/a.boehm/dev/tools/Generate-react-tsx-component/src/commands/Generate.ts`
  //   );
  //   if (args.file && flags.force) {
  //     this.log(`you input --force and --file: ${args.file}`);
  //   }
  // }
}
