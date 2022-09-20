package converters

import (
	"encoding/json"
	"fmt"
	"github.com/linkedin/goavro/v2"
	"strings"
)

func OcfFileRead(ocfFileContents string) ([]byte, error) {
	// Reading OCF data
	ocfReader, err := goavro.NewOCFReader(strings.NewReader(ocfFileContents))
	if err != nil {
		return nil, err
	}
	fmt.Println("Records in OCF File")

	var mp []interface{}

	for ocfReader.Scan() {
		record, err := ocfReader.Read()
		if err != nil {
			return nil, err
		}
		fmt.Println("record", record)
		mp = append(mp, record)
	}

	fmt.Println(mp)

	jsnByt, err := json.Marshal(mp)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(jsnByt))

	return jsnByt, nil
}
