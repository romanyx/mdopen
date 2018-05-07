[![GoDoc](https://godoc.org/github.com/romanyx/mdopen?status.svg)](https://godoc.org/github.com/romanyx/mdopen)
[![Build Status](https://travis-ci.org/romanyx/mdopen.png)](https://travis-ci.org/romanyx/mdopen)
[![Go Report Card](https://goreportcard.com/badge/github.com/romanyx/mdopen)](https://goreportcard.com/report/github.com/romanyx/mdopen)

# mdopen

Allows to view markdown files in the default browser. For more details, see the API [documentation](https://godoc.org/github.com/romanyx/mdopen).

## CLI usage

Install:

```bash
go get github.com/romanyx/mdopen/cmd/mdopen
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
go get github.com/romanyx/mdopen
```

``` go
package main

import "github.com/romanyx/mdopen"

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
