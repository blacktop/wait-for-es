# wait-for-es

[![GoDoc](https://godoc.org/github.com/blacktop/wait-for-es?status.svg)](https://godoc.org/github.com/blacktop/wait-for-es) [![License](http://img.shields.io/:license-mit-blue.svg)](http://doge.mit-license.org) [![Github All Releases](https://img.shields.io/github/downloads/blacktop/wait-for-es/total.svg)](https://github.com/blacktop/wait-for-es/releases/latest) [![GitHub release](https://img.shields.io/github/release/blacktop/wait-for-es.svg)](https://github.com/blacktop/wait-for-es/releases)

> Wait until Elasticsearch become available.

---

## Why 🤔

When testing Elasticsearch in an automated way you have to wait for it to be up and ready to accept connections.

## Install

```bash
go get github.com/blacktop/wait-for-es
```

## Example

```golang
import (
    "context"
    "log"
    "time"

    w4e "github.com/blacktop/wait-for-es"
    "github.com/pkg/errors"
)

func main() {
    connCtx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
    defer cancel()

    w := w4e.WaitForEs{
        URL:     address,
        Timeout: timeout,
    }
    err := w.WaitForConnection(connCtx, timeout)
    if err != nil {
        log.Fatal(errors.Wrap(err, "failed to connect to elasticsearch"))
    }
}
```

## CLI

### Install

Download **wait-for-es** from [releases](https://github.com/blacktop/wait-for-es/releases)

#### macOS

Via [homebrew](https://brew.sh)

```bash
$ brew install blacktop/tap/wait-for-es
```

### Usage

Wait-for-it to be ready

```bash
$ wait-for-es --address http://localhost:9200 --timeout 60

INFO[0000] ===> trying to connect to elasticsearch
INFO[0015] elasticsearch came online after 15 seconds
```

## License

MIT Copyright (c) 2019 blacktop
