package pl_validator

import (
	"errors"
	"strconv"
	"strings"
)

type PeselValidator struct {
	Validator
}

// https://obywatel.gov.pl/pl/dokumenty-i-dane-osobowe/czym-jest-numer-pesel
func (s *PeselValidator) validate(input string) error {

	helper := Helper{}

	times := []int{1, 3, 7, 9}

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

	var inputAsArray []int

	for pos, char := range input {
		Int, err := strconv.Atoi(char)

		append(inputAsArray, Int)
	}
	inputAsArray := strings.Split(input, "")

	return nil
}
