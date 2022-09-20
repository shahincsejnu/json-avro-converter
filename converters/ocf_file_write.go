package converters

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/linkedin/goavro/v2"
)

func OcfFileWrite(avroSchema string, ocfData []byte) (string, error) {
	// Writing OCF data
	var ocfFileContents bytes.Buffer
	writer, err := goavro.NewOCFWriter(goavro.OCFConfig{
		W:      &ocfFileContents,
		Schema: avroSchema,
	})
	if err != nil {
		return "", err
	}

	// convert ocfData (json format) to Go native form
	var m []map[string]interface{}
	err = json.Unmarshal(ocfData, &m)
	if err != nil {
		return "", err
	}

	err = writer.Append(m)
	if err != nil {
		return "", err
	}

	fmt.Println("ocfFileContents", ocfFileContents.String())

	return ocfFileContents.String(), nil
}
