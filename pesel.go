package pl_validator

import "errors"

type PeselValidator struct {
	Validator
}

func (s *PeselValidator) validate(input uint64) error {

	count := checkDigitsCount(input)
	if count < 11 {
		return errors.New("pesel number is too short")
	}
	if count > 11 {
		return errors.New("pesel number is too long")
	}
	return nil
}

func checkDigitsCount(input uint64) int {
	count := 0
	for input != 0 {
		input /= 10
		count += 1
	}
	return count
}
