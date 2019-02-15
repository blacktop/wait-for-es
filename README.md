# wait-for-es

[![GoDoc](https://godoc.org/github.com/blacktop/wait-for-es?status.svg)](https://godoc.org/github.com/blacktop/wait-for-es) [![License](http://img.shields.io/:license-mit-blue.svg)](http://doge.mit-license.org)

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
    "io/ioutil"
    "log"

    "github.com/blacktop/wait-for-es"
    "github.com/pkg/errors"
)

func main() {
    dat, err := ioutil.ReadFile("compressed.bin")
    if err != nil {
        log.Fatal(errors.Wrap(err, "failed to read compressed file"))
    }

    decompressed := wait-for-es.Decompress(dat)
    err = ioutil.WriteFile("compressed.bin.decompressed", decompressed, 0644)
    if err != nil {
        log.Fatal(errors.Wrap(err, "failed to decompress file"))
    }
}
```

## CLI

### Install

Download **wait-for-es** from [releases](https://github.com/blacktop/wait-for-es/releases)

### Usage

Wait-for-it to be ready

```bash
$ wait-for-es --addr localhost:9200 --timeout 60

...
```

## License

MIT Copyright (c) 2019 blacktop
