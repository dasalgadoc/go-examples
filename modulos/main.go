package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/PaesslerAG/jsonpath"
	"github.com/shopspring/decimal"
)

type ClientResponse map[string]interface{}

func main() {
	testMarshall()
}

func testMarshall() {
	var clientResponse ClientResponse

	pi := decimal.NewFromFloat(3.14)
	rateString, _ := pi.Float64()
	fmt.Printf("%T - %v\n", pi, pi)
	s := `{"valor_pi": %.2f}`
	b := []byte(fmt.Sprintf(s, rateString))

	err := json.Unmarshal(b, &clientResponse)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(clientResponse)
}

func testJson() {
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
