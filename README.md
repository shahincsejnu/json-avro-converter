# json-avro-converter
This project is to convert from `JSON to Avro` format and `Avro to JSON`, it also support OCF file Read/Write.

## How to use this package in your code?
- import this package `"github.com/shahincsejnu/json-avro-converter/converters"` in your Go code, like:
  ```
  import (
	  "github.com/shahincsejnu/json-avro-converter/converters"
  )
  ```
- use respective functions in your code, like:
  ```
     converters.JsonToAvro()
     converters.AvroToJson()
     converters.OcfFileWrite()
     converters.OcfFileWrite()
  ```
  > provide related parameters of these functions too!

## Examples
- Check the [examples](./examples/) folder to see demo use of this project
- To run the examples go to the respective example folder and run `go run main.go`

## Features
- [x] conversion from JSON to Avro format
- [x] conversion from Avro to JSON
- [x] OCF file read
- [x] OCF file Write
- [ ] adding CLI for these conversion
- [ ] adding validator for validating the formats

## JSON & Avro format's example
### JSON format
```json
{
  "name": "shahin",
  "age": 20,
  "roll": "oka",
  "others": [
    {
      "name": "okaoka"
    }
  ]
}
```

### Avro
```
{
  "name": "MyClass",
  "type": "record",
  "namespace": "com.acme.avro",
  "fields": [
    {
      "name": "name",
      "type": "string"
    },
    {
      "name": "age",
      "type": "int"
    },
    {
      "name": "roll",
      "type": "string"
    },
    {
      "name": "others",
      "type": {
        "type": "array",
        "items": {
          "name": "others_record",
          "type": "record",
          "fields": [
            {
              "name": "name",
              "type": "string"
            }
          ]
        }
      }
    }
  ]
}
```


## References
- [goavro example](https://github.com/linkedin/goavro/blob/master/examples/165/main.go)
- [goavro package](https://pkg.go.dev/github.com/linkedin/goavro/v2#section-readme)

