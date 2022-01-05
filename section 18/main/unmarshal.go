package main

import (
	"encoding/json"
	"fmt"
)

type human struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {
	s := `[{"First":"James","Last":"Bond","Age":42},{"First":"Alan","Last":"Walker","Age":23}]`
	bs := []byte(s)
	fmt.Printf("%T\n", bs)
	var humans []human
	err := json.Unmarshal(bs, &humans)
	if err != nil {
		fmt.Println(err)
	}
	for i, v := range humans {
		fmt.Println(i, ":", v.First, v.Last, "is", v.Age, "years old.")
	}
}

//[{"First":"James","Last":"Bond","Age":42},{"First":"Alan","Last":"Walker","Age":23}]
