[![Build Status](https://github.com/abtris/ga-badge/actions/workflows/node.js.yml/badge.svg)](https://github.com/abtris/ga-badge/actions)
# Badge Generator

- naive implementation
- only for Github Actions
- no support

## App

![](screenshot.png)

## CLI

```bash
$ gab

NAME:
   gab - A new cli application

USAGE:
   gab [global options] command [command options] [arguments...]

COMMANDS:
   create, c  Create github action badge from URL
   help, h    Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help (default: false)
```

and

```bash
$ /gab create https://github.com/abtris/ga-badge/actions/workflows/node.js.yml/badge.svg
[![Build Status](https://github.com/abtris/ga-badge/actions/workflows/node.js.yml/badge.svg)](https://github.com/abtris/ga-badge/actions)
```
