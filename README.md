[![GoDoc](https://godoc.org/gopkg.in/romanyx/mdopen.v1?status.svg)](https://godoc.org/gopkg.in/romanyx/mdopen.v1)
[![Build Status](https://travis-ci.org/romanyx/mdopen.png)](https://travis-ci.org/romanyx/mdopen)
[![Go Report Card](https://goreportcard.com/badge/github.com/romanyx/mdopen)](https://goreportcard.com/report/github.com/romanyx/mdopen)

# mdopen

Allows to view markdown files in the default browser. For more details, see the API [documentation](https://godoc.org/gopkg.in/romanyx/mdopen.v1).

## CLI usage

Install:

```bash
go get gopkg.in/romanyx/mdopen.v1/cmd/mdopen
```

Create a markdown file:

```bash
echo "# Hello from markdown" > hello.md
```

View it in the default browser as html:

```bash
mdopen hello.md
```

You will see:

![Example](https://monosnap.com/image/1erjc9khEuyB3fHSr1qQaJE5BYhzPC.png)

## API usage

Install:

```bash
go get gopkg.in/romanyx/mdopen.v1
```

``` go
package main

import "gopkg.in/romanyx/mdopen.v1"

func main() {
    f := strings.NewReader("# Hello from markdown")

    opnr := mdopen.New()
    if err := opnr.Open(f); err != nil {
        log.Fatal(err)
    }
}
```

## Contributing

Please feel free to submit issues, fork the repository and send pull requests!
