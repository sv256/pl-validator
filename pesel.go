package pl_validator

import (
	"errors"
)

type PeselValidator struct {
	Validator
}

func (s *PeselValidator) validate(input string) error {

	helper := Helper{}

	count := len(input)
	if count < 11 {
		return errors.New("pesel number is too short")
	}
	if count > 11 {
		return errors.New("pesel number is too long")
	}

	//year, _ := helper.digit(input, 0, 2)
	month, _ := helper.digit(input, 2, 4)
	day, _ := helper.digit(input, 4, 6)

	if month == 0 || month%20 > 12 {
		return errors.New("incorrect month. (range 1-12)")
	}
	if day < 1 || day > 31 {
		return errors.New("incorrect day. (range 1-31)")
	}

	return nil
}
