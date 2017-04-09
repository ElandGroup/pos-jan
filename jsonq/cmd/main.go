package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/jmoiron/jsonq"
)

func main() {
	j := `{
	"foo": 1,
	"bar": 2,
	"test": "Hello, world!",
	"baz": 123.1,
	"array": [
		{"foo": 1},
		{"bar": 2},
		{"baz": 3}
	],
	"subobj": {
		"foo": 1,
		"subarray": [1,2,3],
		"subsubobj": {
			"bar": 2,
			"baz": 3,
			"array": ["hello", "world"]
		}
	},
	"bool": true
}`

	data := map[string]interface{}{}
	dec := json.NewDecoder(strings.NewReader(j))
	dec.Decode(&data)
	jq := jsonq.NewQuery(data)
	// data["foo"] -> 1
	// fmt.Println(jq.Int("foo"))

	// // data["subobj"]["subarray"][1] -> 2
	// fmt.Println(jq.Int("subobj", "subarray", "1"))

	// // data["subobj"]["subarray"]["array"][0] -> "hello"
	// fmt.Println(jq.String("subobj", "subsubobj", "array", "0"))

	// // data["subobj"] -> map[string]interface{}{"subobj": ...}
	// fmt.Println(jq.Object("subobj"))
	fmt.Println(jq.ArrayOfObjects("array"))
}
