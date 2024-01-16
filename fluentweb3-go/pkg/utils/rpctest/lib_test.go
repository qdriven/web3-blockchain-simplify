package rpctest

import (
	"fmt"
	"reflect"
	"testing"

	"gotest.tools/assert"
)

type studentsOrNames struct {
	Names    []string
	Students []student
}
type student struct {
	Name     string
	Treasure *[]byte `json:"treasure,omitempty" testomit:"false"`
	Nicon    *string `json:"nicon" testomit:"true"`
}

type teacher struct {
	Name     string
	Students []studentsOrNames
}

var table = []struct {
	raw    interface{}
	expect string
}{
	{
		raw: teacher{
			Name: "sophia",
			Students: []studentsOrNames{
				{
					Students: []student{
						{
							Name:     "jack",
							Treasure: nil,
							Nicon:    nil,
						},
					},
				},
			},
		},
		expect: `{"Name":"sophia","Students":[{"Names":null,"Students":[{"Name":"jack","treasure":null}]}]}`,
	},
	{
		raw: map[string]interface{}{
			"Name": "sophia",
			"Students": []interface{}{
				map[string]interface{}{
					"Name":     "jack",
					"treasure": nil,
					"nicon":    nil,
				},
			},
		},
		expect: `{"Name":"sophia","Students":[{"Name":"jack","nicon":null,"treasure":null}]}`,
	},
}

func TestJsonMarshalForTest(t *testing.T) {
	for _, item := range table {
		j, e := JsonMarshalForRpcTest(item.raw, false)
		if e != nil {
			panic(e)
		}
		fmt.Printf("marshal result: %s", j)
		assert.Equal(t, item.expect, string(j))
	}
}

func TestA(t *testing.T) {
	fmt.Printf("%v", reflect.TypeOf(nil))
	v, e := convertType(nil, interface{}(nil))
	if e != nil {
		panic(e)
	}
	fmt.Printf("%v\n", v)
}
