package converters

import (
	"encoding/json"
	"strings"

	"github.com/linkedin/goavro/v2"
)

func OcfFileRead(ocfFileContents string) ([]byte, error) {
	// Reading OCF data
	ocfReader, err := goavro.NewOCFReader(strings.NewReader(ocfFileContents))
	if err != nil {
		return nil, err
	}

	var mp []interface{}

	for ocfReader.Scan() {
		record, err := ocfReader.Read()
		if err != nil {
			return nil, err
		}
		mp = append(mp, record)
	}

	jsnByt, err := json.Marshal(mp)
	if err != nil {
		return nil, err
	}

	return jsnByt, nil
}
