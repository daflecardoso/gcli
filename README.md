## gcli

Helper for conventional commits, see more: https://www.conventionalcommits.org/en/v1.0.0/

![Architecture](https://github.com/daflecardoso/gcli/blob/main/example.png)

A single native binary — no Node, no Go, no runtime required to run it.

## Installation

**Linux / macOS**

```sh
curl -fsSL https://raw.githubusercontent.com/daflecardoso/gcli/main/install.sh | sh
```

**Windows (PowerShell)**

```powershell
iwr -useb https://raw.githubusercontent.com/daflecardoso/gcli/main/install.ps1 | iex
```

Both scripts detect your OS/architecture, download the matching binary from the
[latest release](https://github.com/daflecardoso/gcli/releases/latest), and put it on your PATH.

## Setup

Create a `gcli.json` file in the root of your project:

```json
{
  "name": "name",
  "color": "#1476FF",
  "showTutorial": false,
  "scopes": [
    "home",
    "products",
    "foo",
    "bar"
  ]
}
```

## Usage

```sh
gcli
```

It walks you through commit type, scope, message, and optional breaking change,
then runs `git add .`, `git commit`, and `git push` for you.

## Update gcli

```sh
gcli --update
```

This re-runs the install script to fetch the latest release.

## Building from source

Requires Go 1.21+.

```sh
go build -o gcli ./cmd/gcli
```

## Releasing

Tag a version and push it; CI (GitHub Actions + [GoReleaser](https://goreleaser.com))
builds binaries for linux/macOS/windows (amd64/arm64) and publishes them as a GitHub
Release, which `install.sh`/`install.ps1` always pull from `latest`.

```sh
git tag v1.0.0
git push origin v1.0.0
```
