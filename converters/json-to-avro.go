package converters

import (
	"github.com/linkedin/goavro/v2"
)

func JsonToAvro(jsn []byte, avroSchema string) ([]byte, error) {
	// Create a Codec to translate between a byte slice of either binary or textual Avro data and native Go data
	codec, err := goavro.NewCodec(avroSchema)
	if err != nil {
		return nil, err
	}

	// Convert textual Avro data (in Avro JSON format) to native Go form
	native, _, err := codec.NativeFromTextual(jsn)
	if err != nil {
		return nil, err
	}

	// Convert native Go form to binary Avro data
	return codec.BinaryFromNative(nil, native)
}
