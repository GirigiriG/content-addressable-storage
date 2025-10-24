package utils

import (
	"encoding/json"
	"fmt"
)

func GetStructLength(data ...any) int {
	contentLen := 0
	for _, item := range data {

		jsonBytes, err := json.Marshal(item)
		if err != nil {
			fmt.Println("error occurred marshaling struct data", err)
			return 0
		}

		contentLen += len(jsonBytes)
	}
	return contentLen
}

func ParsePayloadToJSON(data ...any) []byte {
	dataBtye, err := json.Marshal(data)
	if err != nil {
		fmt.Println("error parsing struct to json string", err)
		return nil
	}

	return dataBtye
}
