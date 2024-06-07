package main

import (
	"fmt"
	"html"
	"io"
	"os"
	"path/filepath"
	"strings"

	"golang.design/x/clipboard"

	"github.com/buger/jsonparser"
	"github.com/kendfss/but"
	lib "github.com/kendfss/jsol"
	"github.com/kendfss/mandy"
	"github.com/kendfss/pipe"
)

var (
	name string

	cli = mandy.NewCommand(name, mandy.ExitOnError)

	getArg, putArg, decodeArg bool
	queryArg, jsonArg, sepArg string
)

func init() {
	cli.Format = "%s [options] [src]+"
	cli.URL = filepath.Join("https://github.com/kendfss", name)

	cli.Bool(&getArg, "get", false, "get from clipboard", true)
	cli.Bool(&putArg, "put", false, "put result on clipboard", true)
	cli.Bool(&decodeArg, "decode", false, "decode input (restore quotes and ampersands)", true)
	cli.String(&sepArg, "sep", "/", "separator-string to split query path", true)
	cli.String(&queryArg, "query", "", "string to join arguments together", true)
}

func main() {
	but.Must(clipboard.Init())
	but.Must(cli.Parse())

	if data := pipe.Get(); len(data) == 0 {
		if cli.HelpNeeded() {
			fmt.Fprintln(os.Stderr, cli.Usage())
			os.Exit(1)
		}

		if getArg {
			jsonArg = string(clipboard.Read(0))
		} else {
			if jsonArg == "" {
				switch len(cli.Args()) {
				default:
					jsonArg = strings.Join(cli.Args(), " ")

				case 0:
					buf, err := io.ReadAll(os.Stdin)
					lib.Must(err)

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
			jsonArg = html.UnescapeString(jsonArg)
		}

		if len(queryArg) > 0 {
			out, _, _, err := jsonparser.Get([]byte(jsonArg), strings.Split(queryArg, sepArg)...)
			but.Exif(err != nil)
			jsonArg = string(out)
		}
		fmt.Println(string(lib.MustPrettify(jsonArg)))

		if putArg {
			clipboard.Write(0, []byte(lib.Format(jsonArg)))
		}
	} else {
		if len(queryArg) > 0 {
			out, _, _, err := jsonparser.Get(data, strings.Split(queryArg, sepArg)...)
			but.Exif(err != nil)
			data = out
		}

		fmt.Println(string(lib.MustPrettify(data)))
	}
}

func getErr(to *error) {
	*to = recover().(error)
	if *to != nil {
		clipboard.Write(0, []byte((*to).Error()))
	}
}

func getPipe() []byte {
	var data []byte
	info, err := os.Stdin.Stat()
	but.Must(err)

	if info.Size() > 0 {
		data, err = io.ReadAll(os.Stdin)
		but.Must(err)
	}

	return data
}
