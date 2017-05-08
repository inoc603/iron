package iron

import (
	"fmt"
	"testing"
)

type Block struct {
	Str    string        `json:"str,omitempty"`
	Int    int           `json:"int"`
	Array  []interface{} `json:"array"`
	Struct interface{}
}

func TestFlattern(t *testing.T) {
	obj := Block{
		Str:   "L1",
		Int:   1,
		Array: []interface{}{1, 2, 3},
		Struct: Block{
			Str:   "L2",
			Int:   2,
			Array: []interface{}{2, "string"},
			Struct: map[string]int{
				"haha": 1,
			},
		},
	}
	flat := Flatten(obj)
	fmt.Println(flat)
}
