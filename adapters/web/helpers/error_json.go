package helpers

import "encoding/json"

func JsonError(msg string) []byte {

	errorStruct := struct {
		Message string `json:"message"`
	}{
		Message: msg,
	}

	json, err := json.Marshal(errorStruct)

	if err != nil {
		return []byte(err.Error())
	}

	return json
}
