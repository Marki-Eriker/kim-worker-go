package validator

import "fmt"

func (v *Validator) MaxLength(field, value string, condition int) bool {
	if _, ok := v.Errors[field]; ok {
		return false
	}

	if len(value) > condition {
		v.Errors[field] = fmt.Sprintf("%s must be shorter than (%d) characters long", field, condition)
		return false
	}

	return true
}
