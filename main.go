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

const avroSchema = `
	{
	  "type": "record",
	  "name": "test_schema",
	  "fields": [
		{
		  "name": "time",
		  "type": "long"
		},
		{
		  "name": "customer",
		  "type": "string"
		}
	  ]
	}`

var arrayOfEntries = []byte(`
		[
			{
				"time":     1617104831727,
				"customer": "customer1"
			},
			{
				"time":     1717104831727,
				"customer": "customer2"
			}
		]
	`)

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
		avroToJson()
	} else if conversionType == "JSON_TO_AVRO" {
		jsonToAvro()
	} else {
		log.Println(conversionType, "is not supported yet!!")
	}

	ocfFileWrite()
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

func jsonToAvro() {
	avroBinary, err := converters.JsonToAvro(jsn, loginEventAvroSchema)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("avroBinary :", string(avroBinary))
	fmt.Printf("binary: %#v", avroBinary)
}

func ocfFileRead() {
	jsn, err := converters.OcfFileRead("")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(string(jsn))
}

func ocfFileWrite() {
	ocfFileContents, err := converters.OcfFileWrite(avroSchema, arrayOfEntries)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(ocfFileContents)
}
