package utils

func ReplaceFloatIfNotNil(newValue *float64, oldValue float64) float64 {
	if newValue == nil {
		return oldValue
	}

	return *newValue
}

func ReplaceIntIfNotNil(newValue *int, oldValue int) int {
	if newValue == nil {
		return oldValue
	}

	return *newValue
}
