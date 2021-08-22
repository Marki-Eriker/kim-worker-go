package model

import (
	"github.com/marki-eriker/kim-worker-go/internal/validator"
)

func (i UserCreateInput) Validate() (bool, map[string]string) {
	v := validator.New()

	v.Required("email", i.Email)
	v.IsEmail("email", i.Email)

	v.Required("password", i.Password)
	v.MinLength("password", i.Password, 6)

	v.Required("fullName", i.FullName)
	v.MinLength("fullName", i.FullName, 3)

	return v.IsValid(), v.Errors
}

func (i UserUpdateMeInput) Validate() (bool, map[string]string) {
	v := validator.New()

	if i.Email != nil {
		v.IsEmail("email", *i.Email)
	}

	if i.FullName != nil {
		v.MinLength("fullName", *i.FullName, 3)
	}

	return v.IsValid(), v.Errors
}
