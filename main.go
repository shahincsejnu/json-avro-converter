package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/shahincsejnu/json-avro-converter/converters"
	"log"
	"os"
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

	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
		return
	}

	conversionType := os.Getenv("CONVERSION_TYPE")

	if conversionType == "AVRO_TO_JSON" {

	} else if conversionType == "JSON_TO_AVRO" {
		avroBinary, err := converters.JsonToAvro(jsn, loginEventAvroSchema)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("avroBinary :", string(avroBinary))
	} else {
		log.Println(conversionType, "is not supported yet!!")
	}
}
