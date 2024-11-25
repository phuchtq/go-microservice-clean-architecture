package helper

import "encoding/json"

func ToJson(data interface{}) ([]byte, error) {
	return json.Marshal(data)
}

func ConvertJsonToModel[T any](jsonData string) *T {
	var res T

	if json.Unmarshal([]byte(jsonData), &res) != nil {
		return nil
	}

	return &res
}

func ConvertModelToString(data interface{}) string {
	jsonData, err := ToJson(data)

	if err != nil {
		return ""
	}

	return string(jsonData)
}
