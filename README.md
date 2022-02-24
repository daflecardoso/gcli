# gcli
Helper for conventional commits

## Instalation

```sh
bash < (curl -s https://raw.githubusercontent.com/daflecardoso/gcli/main/installer.sh)
```

## Setup

create a ```gcli.json``` file root project

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

## Update gcli

```sh
gcli --update
```

