package utils

import "errors"

// StringToBool converts string to bool.
func StringToBool(value string) (bool, error) {
	switch value {
	case "true":
		return true, nil
	case "false":
		return false, nil
	default:
		return false, errors.New("value must be either true or false")
	}
}
