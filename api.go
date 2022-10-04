package jsol

import (
	"encoding/json"

	pj "github.com/hokaccha/go-prettyjson"

	"github.com/kendfss/but"
)

func Format(arg string) (out string) {
	defer func() {
		if recover() != nil || out == "" {
			out = parseMap(arg)
		}
	}()

	out = parseSlice(arg)
	return
}

func parseSlice(arg string) string {
	m := make([]interface{}, 0) //
	but.Must(json.Unmarshal([]byte(arg), &m))
	data, err := json.MarshalIndent(m, "", "\t")
	but.Must(err)

	return string(data)
}

func parseMap(arg string) string {
	m := make(map[interface{}]interface{}) //
	but.Must(json.Unmarshal([]byte(arg), &m))
	data, err := json.MarshalIndent(m, "", "\t")
	but.Must(err)

	return string(data)
}

func Prettify(obj any) (data []byte, err error) {
	var tmpData []byte
	switch obj.(type) {
	case string:
		data = []byte(obj.(string))
	case []byte:
		data = obj.([]byte)
	default:
		tmpData, err = json.MarshalIndent(obj, "", "  ")
		if err != nil {
			return nil, err
		}
		data = tmpData
	}

	data, err = pj.Format([]byte(data))
	if err != nil {
		return nil, err
	}

	return data, err
}

func MustPrettify(obj any) []byte {
	out, err := Prettify(obj)
	if err != nil {
		panic(err)
	}
	return out
}
