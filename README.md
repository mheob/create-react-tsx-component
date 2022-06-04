oclif-hello-world
=================

oclif example Hello World CLI

[![oclif](https://img.shields.io/badge/cli-oclif-brightgreen.svg)](https://oclif.io)
[![Version](https://img.shields.io/npm/v/oclif-hello-world.svg)](https://npmjs.org/package/oclif-hello-world)
[![CircleCI](https://circleci.com/gh/oclif/hello-world/tree/main.svg?style=shield)](https://circleci.com/gh/oclif/hello-world/tree/main)
[![Downloads/week](https://img.shields.io/npm/dw/oclif-hello-world.svg)](https://npmjs.org/package/oclif-hello-world)
[![License](https://img.shields.io/npm/l/oclif-hello-world.svg)](https://github.com/oclif/hello-world/blob/main/package.json)

<!-- toc -->
* [Usage](#usage)
* [Commands](#commands)
<!-- tocstop -->
# Usage
<!-- usage -->
```sh-session
$ npm install -g create-react-tsx-component
$ create-react-tsx-component COMMAND
running command...
$ create-react-tsx-component (--version)
create-react-tsx-component/0.0.0 darwin-x64 node-v18.0.0
$ create-react-tsx-component --help [COMMAND]
USAGE
  $ create-react-tsx-component COMMAND
...
```
<!-- usagestop -->
# Commands
<!-- commands -->
* [`create-react-tsx-component hello PERSON`](#create-react-tsx-component-hello-person)
* [`create-react-tsx-component hello world`](#create-react-tsx-component-hello-world)
* [`create-react-tsx-component help [COMMAND]`](#create-react-tsx-component-help-command)
* [`create-react-tsx-component plugins`](#create-react-tsx-component-plugins)
* [`create-react-tsx-component plugins:install PLUGIN...`](#create-react-tsx-component-pluginsinstall-plugin)
* [`create-react-tsx-component plugins:inspect PLUGIN...`](#create-react-tsx-component-pluginsinspect-plugin)
* [`create-react-tsx-component plugins:install PLUGIN...`](#create-react-tsx-component-pluginsinstall-plugin-1)
* [`create-react-tsx-component plugins:link PLUGIN`](#create-react-tsx-component-pluginslink-plugin)
* [`create-react-tsx-component plugins:uninstall PLUGIN...`](#create-react-tsx-component-pluginsuninstall-plugin)
* [`create-react-tsx-component plugins:uninstall PLUGIN...`](#create-react-tsx-component-pluginsuninstall-plugin-1)
* [`create-react-tsx-component plugins:uninstall PLUGIN...`](#create-react-tsx-component-pluginsuninstall-plugin-2)
* [`create-react-tsx-component plugins update`](#create-react-tsx-component-plugins-update)

## `create-react-tsx-component hello PERSON`

Say hello

```
USAGE
  $ create-react-tsx-component hello [PERSON] -f <value>

ARGUMENTS
  PERSON  Person to say hello to

FLAGS
  -f, --from=<value>  (required) Whom is saying hello

DESCRIPTION
  Say hello

EXAMPLES
  $ oex hello friend --from oclif
  hello friend from oclif! (./src/commands/hello/index.ts)
```

_See code: [dist/commands/hello/index.ts](https://github.com/mheob/create-react-tsx-component/blob/v0.0.0/dist/commands/hello/index.ts)_

## `create-react-tsx-component hello world`

Say hello world

```
USAGE
  $ create-react-tsx-component hello world

DESCRIPTION
  Say hello world

EXAMPLES
  $ oex hello world
  hello world! (./src/commands/hello/world.ts)
```

## `create-react-tsx-component help [COMMAND]`

Display help for create-react-tsx-component.

```
USAGE
  $ create-react-tsx-component help [COMMAND] [-n]

ARGUMENTS
  COMMAND  Command to show help for.

FLAGS
  -n, --nested-commands  Include all nested commands in the output.

DESCRIPTION
  Display help for create-react-tsx-component.
```

_See code: [@oclif/plugin-help](https://github.com/oclif/plugin-help/blob/v5.1.10/src/commands/help.ts)_

## `create-react-tsx-component plugins`

List installed plugins.

```
USAGE
  $ create-react-tsx-component plugins [--core]

FLAGS
  --core  Show core plugins.

DESCRIPTION
  List installed plugins.

EXAMPLES
  $ create-react-tsx-component plugins
```

_See code: [@oclif/plugin-plugins](https://github.com/oclif/plugin-plugins/blob/v2.0.11/src/commands/plugins/index.ts)_

## `create-react-tsx-component plugins:install PLUGIN...`

Installs a plugin into the CLI.

```
USAGE
  $ create-react-tsx-component plugins:install PLUGIN...

ARGUMENTS
  PLUGIN  Plugin to install.

FLAGS
  -f, --force    Run yarn install with force flag.
  -h, --help     Show CLI help.
  -v, --verbose

DESCRIPTION
  Installs a plugin into the CLI.

  Can be installed from npm or a git url.

  Installation of a user-installed plugin will override a core plugin.

  e.g. If you have a core plugin that has a 'hello' command, installing a user-installed plugin with a 'hello' command
  will override the core plugin implementation. This is useful if a user needs to update core plugin functionality in
  the CLI without the need to patch and update the whole CLI.

ALIASES
  $ create-react-tsx-component plugins add

EXAMPLES
  $ create-react-tsx-component plugins:install myplugin 

  $ create-react-tsx-component plugins:install https://github.com/someuser/someplugin

  $ create-react-tsx-component plugins:install someuser/someplugin
```

## `create-react-tsx-component plugins:inspect PLUGIN...`

Displays installation properties of a plugin.

```
USAGE
  $ create-react-tsx-component plugins:inspect PLUGIN...

ARGUMENTS
  PLUGIN  [default: .] Plugin to inspect.

FLAGS
  -h, --help     Show CLI help.
  -v, --verbose

DESCRIPTION
  Displays installation properties of a plugin.

EXAMPLES
  $ create-react-tsx-component plugins:inspect myplugin
```

## `create-react-tsx-component plugins:install PLUGIN...`

Installs a plugin into the CLI.

```
USAGE
  $ create-react-tsx-component plugins:install PLUGIN...

ARGUMENTS
  PLUGIN  Plugin to install.

FLAGS
  -f, --force    Run yarn install with force flag.
  -h, --help     Show CLI help.
  -v, --verbose

DESCRIPTION
  Installs a plugin into the CLI.

  Can be installed from npm or a git url.

  Installation of a user-installed plugin will override a core plugin.

  e.g. If you have a core plugin that has a 'hello' command, installing a user-installed plugin with a 'hello' command
  will override the core plugin implementation. This is useful if a user needs to update core plugin functionality in
  the CLI without the need to patch and update the whole CLI.

ALIASES
  $ create-react-tsx-component plugins add

EXAMPLES
  $ create-react-tsx-component plugins:install myplugin 

  $ create-react-tsx-component plugins:install https://github.com/someuser/someplugin

  $ create-react-tsx-component plugins:install someuser/someplugin
```

## `create-react-tsx-component plugins:link PLUGIN`

Links a plugin into the CLI for development.

```
USAGE
  $ create-react-tsx-component plugins:link PLUGIN

ARGUMENTS
  PATH  [default: .] path to plugin

FLAGS
  -h, --help     Show CLI help.
  -v, --verbose

DESCRIPTION
  Links a plugin into the CLI for development.

  Installation of a linked plugin will override a user-installed or core plugin.

  e.g. If you have a user-installed or core plugin that has a 'hello' command, installing a linked plugin with a 'hello'
  command will override the user-installed or core plugin implementation. This is useful for development work.

EXAMPLES
  $ create-react-tsx-component plugins:link myplugin
```

## `create-react-tsx-component plugins:uninstall PLUGIN...`

Removes a plugin from the CLI.

```
USAGE
  $ create-react-tsx-component plugins:uninstall PLUGIN...

ARGUMENTS
  PLUGIN  plugin to uninstall

FLAGS
  -h, --help     Show CLI help.
  -v, --verbose

DESCRIPTION
  Removes a plugin from the CLI.

ALIASES
  $ create-react-tsx-component plugins unlink
  $ create-react-tsx-component plugins remove
```

## `create-react-tsx-component plugins:uninstall PLUGIN...`

Removes a plugin from the CLI.

```
USAGE
  $ create-react-tsx-component plugins:uninstall PLUGIN...

ARGUMENTS
  PLUGIN  plugin to uninstall

FLAGS
  -h, --help     Show CLI help.
  -v, --verbose

DESCRIPTION
  Removes a plugin from the CLI.

ALIASES
  $ create-react-tsx-component plugins unlink
  $ create-react-tsx-component plugins remove
```

## `create-react-tsx-component plugins:uninstall PLUGIN...`

Removes a plugin from the CLI.

```
USAGE
  $ create-react-tsx-component plugins:uninstall PLUGIN...

ARGUMENTS
  PLUGIN  plugin to uninstall

FLAGS
  -h, --help     Show CLI help.
  -v, --verbose

DESCRIPTION
  Removes a plugin from the CLI.

ALIASES
  $ create-react-tsx-component plugins unlink
  $ create-react-tsx-component plugins remove
```

## `create-react-tsx-component plugins update`

Update installed plugins.

```
USAGE
  $ create-react-tsx-component plugins update [-h] [-v]

FLAGS
  -h, --help     Show CLI help.
  -v, --verbose

DESCRIPTION
  Update installed plugins.
```
<!-- commandsstop -->
