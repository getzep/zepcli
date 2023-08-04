# Zep CLI Tool

This is a CLI tool for managing Zep. It's functionality is currently limited to generating authentication secrets and JWT tokens.

## Usage

### Generating a secret and JWT token

```bash
zepcli -i
```

## Installation

MacOS, Linux, and Windows binaries are available on the [releases page](https://github.com/getzep/zepcli/releases).

### From source

```bash
git clone https://github.com/getzep/zepcli.git
cd zepcli
make build
```

### With `go install`

```bash
go install github.com/getzep/zepcli@latest
```

