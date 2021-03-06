package json

import (
	"bytes"
	"encoding/json"
)

const (
	empty = ""
	tab   = "\t"
)

func PrettyJson(data interface{}) (string, error) {
	buffer := new(bytes.Buffer)
	encoder := json.NewEncoder(buffer)
	encoder.SetIndent(empty, tab)

	err := encoder.Encode(data)
	if err != nil {
		return empty, err
	}
	return buffer.String(), nil
}

func CompactJSON(in string) string {
	var out bytes.Buffer
	if err := json.Compact(&out, []byte(in)); err != nil {
		return in
	}
	return out.String()
}
