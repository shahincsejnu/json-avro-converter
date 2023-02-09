package main

import (
	"fmt"

	"github.com/shahincsejnu/json-avro-converter/converters"
)

const loginEventAvroSchema = `
	{
		"type": "record", 
		"name": "LoginEvent", 
		"fields": 
		[
			{
				"name": "Username", 
				"type": "string"
			},
			{
				"name": "Password",
				"type": "string"
			}
		]
	}`

var jsn = []byte(`
		{
			"Username": "Sahadat Hossain",
			"Password": "oka"
		}	
    `)

func main() {
	fmt.Println("Let's start!")

	jsonToAvro()
}

func jsonToAvro() {
	avroBinary, err := converters.JsonToAvro(jsn, loginEventAvroSchema)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("avroBinary :", string(avroBinary))
	fmt.Printf("binary: %#v", avroBinary)
}
