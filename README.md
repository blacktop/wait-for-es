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

```bash
$ wait-for-es --help

Wait until Elasticsearch become available

Usage:
  wait-for-es [flags]

Flags:
      --address string   elasticsearch address (default "http://localhost:9200")
      --config string    config file (default is $HOME/.wait-for-es.yaml)
  -H, --healthy          wait until cluster health is green
  -h, --help             help for wait-for-es
      --timeout int      timeout (default is 60) (default 60)
  -V, --verbose          verbose output
```

```bash
$ wait-for-es --address http://localhost:9200 --timeout 60 --healthy --verbose

INFO[0000] ===> trying to connect to elasticsearch
DEBU[0000] attempting to PING to: http://localhost:9200
DEBU[0000]  * could not connect to elasticsearch (sleeping for 1 second)
DEBU[0001] attempting to PING to: http://localhost:9200
DEBU[0001]  * could not connect to elasticsearch (sleeping for 1 second)
..SNIP...
DEBU[0021] cluster status is "green"
DEBU[0021] elasticSearch connection successful           cluster=elasticsearch code=200 url="http://localhost:9200" version=7.0.0
INFO[0021] elasticsearch came online after 21 seconds
```

## License

MIT Copyright (c) 2019 blacktop
