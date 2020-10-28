# Dawn Cli
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
