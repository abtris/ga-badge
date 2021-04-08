[![Build Status](https://github.com/abtris/ga-badge/actions/workflows/node.js.yml/badge.svg)](https://github.com/abtris/ga-badge/actions)
# Badge Generator

- naive implementation
- only for Github Actions
- no support

## App

![](screenshot.png)

## CLI

### Install

```
brew install abtris/tap/gab
```

### Usage

Help with simple one command `create`

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

and you can use command `create` with URL

```bash
$ gab create https://github.com/abtris/ga-badge/actions/workflows/node.js.yml/badge.svg
[![Build Status](https://github.com/abtris/ga-badge/actions/workflows/node.js.yml/badge.svg)](https://github.com/abtris/ga-badge/actions)
```

or specify more options

```
$ gab help create
NAME:
   gab create - Create github action badge from URL

USAGE:
   gab create [command options] [arguments...]

OPTIONS:
   --url value, -u value
   --branch value, -b value  (default: "master")
   --label value, -l value   (default: "Build Status")
```

second command make fast init template for my used templates stored in `cli/templates`.

```
$ gab help init
NAME:
   main init - Create github workflow file

USAGE:
   main init [command options] [arguments...]

OPTIONS:
   --lang value, -l value  (default: "node.js")
```
