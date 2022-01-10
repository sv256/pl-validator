package pl_validator

import (
	"errors"
	"strconv"
)

type Helper struct {
}

func (s *Helper) digit(num string, start int, end int) (int, error) {
	if start == end {
		return -1, errors.New("start and end value cannot be the same")
	}
	if start > end {
		return -1, errors.New("start value cannot be greater then the end one")
	}
	if start < 0 || end < 0 {
		return -1, errors.New("start end end value cannot be less than 0")
	}
	numStr := num[start:end]
	intVar, err := strconv.Atoi(numStr)
	return intVar, err
}
