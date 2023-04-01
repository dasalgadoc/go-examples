package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/PaesslerAG/jsonpath"
)

func main() {
	v := interface{}(nil)

	json.Unmarshal([]byte(`{
		"welcome":{
				"message":["Good Morning", "Hello World!"]
			}
		}`), &v)

	welcome, err := jsonpath.Get("$.welcome.message[1]", v)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(welcome)

}
