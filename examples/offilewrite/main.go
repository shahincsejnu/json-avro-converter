package main

import (
	"fmt"

	"github.com/shahincsejnu/json-avro-converter/converters"
)

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

func main() {
	fmt.Println("Let's start!")

	ocfFileWrite()
}

func ocfFileWrite() {
	ocfFileContents, err := converters.OcfFileWrite(avroSchema, arrayOfEntries)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(ocfFileContents)
}
