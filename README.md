# Silkworm

Silkworm is a WordPress plugin update ticket creation tool. It's meant to bridge the gap between [Platypus](https://github.com/nausicaan/platypus.git) and [Bowerbird](https://github.com/nausicaan/bowerbird.git), reading the output from *Platypus* and triggering *Bowerbird* with information recieved and tickets created.

![Silkworm](cocoons.webp)

## Prerequisite

Googles' [Go language](https://go.dev) installed to enable building executables from source code.

A `jira.json` file containing your API URL and Bearer token to enable ticket creation:

``` go
{
    "base": "Jira Issue base URL",
    "token": "Jira Bearer Token"
}
```

## Build

From the root folder the `go` files, use the command that matches your environment:

### Windows & Mac

``` console
go build -o [name] .
```

### Linux

``` console
GOOS=linux GOARCH=amd64 go build -o [name] .
```

## Run

``` console
[program] [optional flag]
```

Example:

``` console
silkworm -h
```

Output:

``` console
Usage:
  [program] [flag]

Example:
  Adding your path to file if necessary, run:
    silkworm

Additional Options:
  -h, --help 		Help Information
  -v, --version 	Display App Version

Help:
  For more information go to:
    https://github.com/nausicaan/silkworm.git
```

## License
Code is distributed under [The Unlicense](https://github.com/nausicaan/free/blob/main/LICENSE.md) and is part of the Public Domain.
