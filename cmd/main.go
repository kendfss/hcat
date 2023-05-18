package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/kendfss/but"
	lib "github.com/kendfss/hcat"
	"github.com/kendfss/mandy"
)

var (
	name string
	cli  = mandy.NewCommand(name, mandy.ExitOnError)

	rawArg bool

	dashes = strings.Repeat("-", 40)
	sep    = fmt.Sprintf("%s %s %s", dashes, eof, dashes)
)

const (
	eof = "END OF DOCUMENT"
)

func init() {
	cli.Bool(&rawArg, "raw", rawArg, "do not prettify the output", true)
	// cli.URL = filepath.Join("https://github.com/kendfss", name)
	cli.Format = "%s [options] [urls...]"
}

func main() {
	cli.MustParse()

	if cli.HelpNeeded() {
		cli.PrintHelp()
	}

	data, err := io.ReadAll(os.Stdin)
	but.Must(err)

	if data == nil || len(cli.Args()) > 0 {
		for i, arg := range cli.Args() {
			if i > 0 {
				println(sep)
			}
			data = lib.Read(lib.Scrape(arg))
			if !rawArg {
				data = lib.Prettify(data)
			}
			fmt.Println(string(data))
		}
	} else {
		if !rawArg {
			data = lib.Prettify(data)
		}

		fmt.Println(string(data))
	}
}
