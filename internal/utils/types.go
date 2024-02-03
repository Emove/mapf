package utils

import (
	"strconv"
	"strings"
)

const (
	Bool   = "bool"
	Int    = "int"
	Float  = "float"
	String = "string"
)

func IsValidBuiltinType(valueType string) bool {
	valueType = strings.ToLower(valueType)
	if valueType == Bool {
		return true
	} else if valueType == Int {
		return true
	} else if valueType == Float {
		return true
	} else if valueType == String {
		return true
	}
	return false
}

func IsValidValue(valueType, value string) bool {
	if !IsValidBuiltinType(valueType) {
		return false
	}
	valueType, value = strings.ToLower(valueType), strings.ToLower(value)
	if valueType == Bool {
		return isValidBoolValue(value)
	} else if valueType == Int {
		return isValidIntValue(value)
	} else if valueType == Float {
		return isValidFloatValue(value)
	}
	return true
}

func isValidBoolValue(value string) bool {
	_, err := strconv.ParseBool(value)
	if err != nil {
		return false
	}
	return true
}

func isValidIntValue(value string) bool {
	_, err := strconv.Atoi(value)
	if err != nil {
		return false
	}
	return true
}

func isValidFloatValue(value string) bool {
	_, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return false
	}
	return true
}
