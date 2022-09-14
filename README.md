# json-avro-converter
This project is to convert from `JSON to Avro` format and `Avro to JSON`, it's implemented using Golang.


## Features
- [x] conversion from JSON to Avro format
- [x] conversion from Avro to JSON
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

