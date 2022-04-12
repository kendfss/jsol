package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"golang.design/x/clipboard"

	"github.com/kendfss/jsol"
	"github.com/kendfss/mandy"
)

var (
	name string

	cli = mandy.NewCommand(name, mandy.ExitOnError)

	getArg, putArg, decodeArg bool
	jsonArg, joinerArg        string
)

func init() {
	cli.Format = "%s src"
	cli.URL = filepath.Join("https://github.com/kendfss", name)

	cli.Bool(&getArg, "get", false, "get from clipboard", true)
	cli.Bool(&putArg, "put", false, "put result on clipboard", true)
	cli.Bool(&decodeArg, "decode", false, "decode input (restore quotes and ampersands)", true)
	cli.String(&joinerArg, "joiner", " ", "string to join arguments together", true)
}

func main() {
	jsol.Must(clipboard.Init())
	jsol.Must(cli.Parse())

	if cli.HelpWanted() {
		fmt.Fprintln(os.Stderr, cli.Usage())
		os.Exit(1)
	}

	if getArg {
		jsonArg = string(clipboard.Read(0))
	} else {
		if jsonArg == "" {
			switch len(cli.Args()) {
			default:
				jsonArg = strings.Join(cli.Args(), joinerArg)

			case 0:
				buf, err := io.ReadAll(os.Stdin)
				jsol.Must(err)

				jsonArg = string(buf)

			}
		}
	}

	if len(jsonArg) == 0 {
		fmt.Fprintln(os.Stderr, fmt.Errorf("%s: no input", name))
		fmt.Fprintln(os.Stderr, cli.Usage())
		os.Exit(1)
	}

	if decodeArg {
		jsonArg = strings.ReplaceAll(jsonArg, "&quot;", `"`)
		jsonArg = strings.ReplaceAll(jsonArg, "&amp;", `&`)
	}

	fmt.Println(string(jsol.Prettify(jsonArg)))

	if putArg {
		clipboard.Write(0, []byte(jsol.Format(jsonArg)))

	}
}

func getErr(to *error) {
	*to = recover().(error)
	if *to != nil {
		clipboard.Write(0, []byte((*to).Error()))
	}
}
