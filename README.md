# Silkworm

Silkworm is a WordPress plugin update ticket creation tool. It's meant to bridge the gap between [Platypus](https://github.com/nausicaan/platypus.git) and [Bowerbird](https://github.com/nausicaan/bowerbird.git), reading the output from *Platypus* and triggering *Bowerbird* with information recieved and tickets created.

## Prerequisite

- Googles' [Go language](https://go.dev) installed to enable building executables from source code.

## Build

From the root folder containing *main.go*, use the command that matches your environment:

### Windows & Mac:

```console
go build -o [name] main.go
```

### Linux:

```console
GOOS=linux GOARCH=amd64 go build -o [name] main.go
```

## Run

```console
./[program] [vendor/plugin]:[version]
```

Example:

```console
./silkworm wpackagist-plugin/all-in-one-seo-pack:4.4.1
```

## License
Code is distributed under [The Unlicense](https://github.com/nausicaan/free/blob/main/LICENSE.md) and is part of the Public Domain.
