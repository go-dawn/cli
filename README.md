# Dawn Cli
<p>
  <a href="https://pkg.go.dev/github.com/go-dawn/cli?tab=doc">
    <img src="https://img.shields.io/badge/%F0%9F%93%9A%20godoc-pkg-00ACD7.svg?color=00ACD7&style=flat">
  </a>
  <a href="https://goreportcard.com/report/github.com/go-dawn/cli">
    <img src="https://img.shields.io/badge/%F0%9F%93%9D%20goreport-A%2B-75C46B">
  </a>
  <a href="https://codecov.io/gh/go-dawn/cli">
    <img src="https://codecov.io/gh/go-dawn/cli/branch/main/graph/badge.svg?token=50B581R7EE"/>
  </a>
  <a href="https://github.com/go-dawn/cli/actions?query=workflow%3ASecurity">
    <img src="https://img.shields.io/github/workflow/status/gofiber/fiber/Security?label=%F0%9F%94%91%20gosec&style=flat&color=75C46B">
  </a>
  <a href="https://github.com/go-dawn/cli/actions?query=workflow%3ATest">
    <img src="https://img.shields.io/github/workflow/status/gofiber/fiber/Test?label=%F0%9F%A7%AA%20tests&style=flat&color=75C46B">
  </a>
  <a>
    <img src="https://counter.gofiber.io/badge/go-dawn/cli">
  </a>
</p>

Dawn Command Line Interface

# Installation
```bash
go get -u github.com/go-dawn/cli/dawn
```

# Commands
## dawn
### Synopsis

```
       __
   ___/ /__ __    _____    dawn-cli v0.0.x
 ~/ _  / _ '/ |/|/ / _ \~  For the opinionated lightweight framework dawn
~~\_,_/\_,_/|__,__/_//_/~~ Visit https://github.com/go-dawn/dawn for detail
 ~~~  ~~ ~~~~~~~~~ ~~~~~~  (c) since 2020 by kiyonlin@gmail.com
```
### Options

```
  -h, --help   help for dawn
```

## dawn dev
### Synopsis

Rerun the dawn project if watched files changed

```
dawn dev [flags]
```

### Options

```
  -d, --delay duration          delay to trigger rerun (default 1s)
  -D, --exclude_dirs strings    ignore these directories (default [assets,tmp,vendor,node_modules])
  -F, --exclude_files strings   ignore these files
  -e, --extensions strings      file extensions to watch (default [go,tmpl,tpl,html])
  -h, --help                    help for dev
  -r, --root string             root path for watch, all files must be under root (default ".")
  -t, --target string           target path for go build (default ".")
```

## dawn generate
### Synopsis

Generate boilerplate code with different commands

```
dawn generate [flags]
dawn generate [command]
```

### Available Commands

```
module      Generate a new dawn module
```

### Options

```
  -h, --help                    help for dev
```

## dawn generate module
### Synopsis

Generate a new dawn module

```
dawn generate module NAME [flags]
```

### Examples

```
  dawn module hello
    Generates a new module dir named hello includes boilerplate code

```

### Options

```
  -h, --help                    help for dev
```

## dawn new
### Synopsis

Generate a new dawn web/application project

```
dawn new PROJECT [module name] [flags]
```

### Examples

```
  dawn new dawn-demo
    Generates a web project with go module name dawn-demo

  dawn new dawn-demo github.com/go-dawn/dawn-demo
    Specific the go module name

  dawn new dawn-demo --app
    Generate an application project

```

### Options

```
  -h, --help              help for new
  -t, --template string   basic|complex (default "basic")
```

## dawn upgrade
### Synopsis

Upgrade Dawn cli if a newer version is available

```
dawn upgrade [flags]
```

### Options

```
  -h, --help   help for upgrade
```

## dawn version
### Synopsis

Print the local and released version number of dawn

```
dawn version [flags]
```

### Options

```
  -h, --help   help for version
```
