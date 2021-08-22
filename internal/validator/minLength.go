package validator

import "fmt"

func (v *Validator) MinLength(field, value string, condition int) bool {
	if _, ok := v.Errors[field]; ok {
		return false
	}

	if len(value) < condition {
		v.Errors[field] = fmt.Sprintf("%s must be at least (%d) characters long", field, condition)
		return false
	}

	return true
}
