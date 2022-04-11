package jsol

import (
	"encoding/json"

	pj "github.com/hokaccha/go-prettyjson"
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
	must(json.Unmarshal([]byte(arg), &m))
	data, err := json.MarshalIndent(m, "", "\t")
	must(err)

	return string(data)
}

func parseMap(arg string) string {
	m := make(map[interface{}]interface{}) //
	must(json.Unmarshal([]byte(arg), &m))
	data, err := json.MarshalIndent(m, "", "\t")
	must(err)

	return string(data)
}

func Prettify(obj any) (data []byte) {
	var err error
	switch obj.(type) {
	case string:
		data = []byte(obj.(string))
	case []byte:
		data = obj.([]byte)
	default:
		data, err = json.Marshal(obj)
		must(err)
	}

	data, err = pj.Format([]byte(data))
	must(err)

	return data
}
