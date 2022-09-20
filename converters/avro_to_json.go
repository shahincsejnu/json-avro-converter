package converters

import (
	"github.com/linkedin/goavro/v2"
)

func AvroToJson(avroBinary []byte, avroSchema string) ([]byte, error) {
	// Create a Codec to translate between a byte slice of either binary or textual Avro data and native Go data
	codec, err := goavro.NewCodec(avroSchema)
	if err != nil {
		return nil, err
	}

	// Convert binary Avro data back to native Go form
	native, _, err := codec.NativeFromBinary(avroBinary)
	if err != nil {
		return nil, err
	}

	return codec.TextualFromNative(nil, native)
}
