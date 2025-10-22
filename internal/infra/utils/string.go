package utils

import "strconv"

func GetStringValueOrDefault(value, def string) string {
	if value != "" {
		return value
	}
	return def
}

func StringToBool(value string, def bool) bool {
	res, err := strconv.ParseBool(value)
	if err != nil {
		return def
	}
	return res
}

func CheckForNil(value interface{}) string {
	if value == nil {
		return ""
	}
	return value.(string)
}

func ReplaceStringIfNotNil(newValue *string, oldValue string) string {
	if newValue == nil {
		return oldValue
	}

	return *newValue
}
