package pl_validator

import (
	"errors"
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

	transform := func(accumulator int, entry int, idx int) int {
		return accumulator + entry*(times[idx%4])
	}

	digits, _ := helper.stringToIntArray(input, "11")
	lastDigit := digits[10]
	digits = digits[:len(digits)-1]
	control := 0
	for index, element := range digits {
		if index == 0 {
			control = control + element*(times[index%4])
		} else {
			control = digits[index-1] + element*(times[index%4])
		}
	}
	result, err := helper.reduce(digits, 0, transform)
	if err != nil {
		return errors.New(err.Error())
	}

	control = result.(int) % 10
	if control == 0 {
		if lastDigit != 0 {
			return errors.New("wrong pesel controlsum")
		}
	} else {
		if 10-control != lastDigit {
			return errors.New("wrong pesel controlsum")
		}
	}

	return nil
}
