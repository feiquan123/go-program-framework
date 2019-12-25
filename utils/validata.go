package utils

import "fmt"

// validate args, exclude bool type
func Validate(args ...interface{}) error {
	for _, arg := range args {
		// type judge
		switch arg.(type) {
		case int:
			if value, ok := arg.(int); ok && value == 0 {
				return fmt.Errorf("int values is 0")
			}
		case string:
			if value, ok := arg.(string); ok && len(value) == 0 {
				return fmt.Errorf("string value length is 0")
			}
		case int64:
			if value, ok := arg.(int64); ok && value == 0 {
				return fmt.Errorf("int64 value is 0")
			}
		case float64:
			if value, ok := arg.(float64); ok && value == 0 {
				return fmt.Errorf("float64 value is 0.0")
			}
		case []int:
			if value, ok := arg.([]int); ok && len(value) == 0 {
				return fmt.Errorf("[]int value is 0")
			}
		case []string:
			if value, ok := arg.([]string); ok && len(value) == 0 {
				return fmt.Errorf("[]string value is nil")
			}
		default:
			return fmt.Errorf("unSupport type")
		}
	}
	return nil
}
