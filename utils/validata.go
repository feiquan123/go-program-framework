package utils

import (
	"errors"
)

// validate args, exclude bool type
func Validate(args ...interface{}) error {
	for _, arg := range args {
		// type judge
		switch arg.(type) {
		case int:
			if value, ok := arg.(int); ok && value == 0 {
				return errors.New("int values is 0")
			}
		case string:
			if value, ok := arg.(string); ok && len(value) == 0 {
				return errors.New("string value length is 0")
			}
		case int64:
			if value, ok := arg.(int64); ok && value == 0 {
				return errors.New("int64 value is 0")
			}
		case float64:
			if value, ok := arg.(float64); ok && value == 0 {
				return errors.New("float64 value is 0.0")
			}
		case []int:
			if value, ok := arg.([]int); ok && len(value) == 0 {
				return errors.New("[]int value is 0")
			}
		case []string:
			if value, ok := arg.([]string); ok && len(value) == 0 {
				return errors.New("[]string value is nil")
			}
		case map[string]string:
			if len(arg.(map[string]string)) == 0 {
				return errors.New("map[string]string is nil")
			}
		default:
			return errors.New("unSupport type")
		}
	}
	return nil
}