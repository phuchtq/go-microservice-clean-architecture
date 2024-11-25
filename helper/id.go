package helper

import (
	model_types "architecture_template/constants/modelTypes"
	"fmt"
)

func GenerateId(entity string, n int) string {
	prefix, format := getFormat(entity)

	if prefix == "" || format == "" {
		return ""
	}

	return prefix + fmt.Sprintf(format, n+1)
}

func getFormat(entity string) (string, string) {
	var prefix string
	var format string

	switch entity {
	case model_types.USER_TYPE:
		prefix = "U"
		format = model_types.USER_FORMAT

	case model_types.ROLE_TYPE:
		prefix = "R"
		format = model_types.ROLE_FORMAT
	}

	return prefix, format
}
