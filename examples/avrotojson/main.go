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

	avroToJson()
}

func avroToJson() {
	avroBinary, err := converters.JsonToAvro(jsn, loginEventAvroSchema)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	jsnByt, err := converters.AvroToJson(avroBinary, loginEventAvroSchema)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("jsnByt:", string(jsnByt))
}
