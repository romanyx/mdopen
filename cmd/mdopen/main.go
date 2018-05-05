package main

import (
	"log"
	"os"

	"github.com/pkg/errors"
	"github.com/romanyx/mdopen"
	"gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "t",
			Value: "github",
			Usage: "Set template",
		},
	}
	app.Name = "mdopen"
	app.Usage = "View markdown files in default browser"
	app.Action = func(c *cli.Context) error {
		var filenames []string
		if c.NArg() > 0 {
			filenames = c.Args()
		}

		options := []mdopen.Option{
			templateOption(c.String("t")),
		}

		opnr := mdopen.New(options...)
		for _, fName := range filenames {
			f, err := os.Open(fName)
			if err != nil {
				return errors.Wrap(err, "failed to open file")
			}

			if err := opnr.Open(f); err != nil {
				return errors.Wrap(err, "open markdown failed")
			}
		}

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func templateOption(t string) mdopen.Option {
	return mdopen.GithubTemplate()
}
