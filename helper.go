package pl_validator

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
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

func (s *Helper) stringToIntArray(input string, size string) ([]int, error) {
	a := strings.Split(input, "")
	n, err := strconv.Atoi(size) // int 32bit
	b := make([]int, n)
	for i, v := range a {
		b[i], err = strconv.Atoi(v)
		if err != nil {
			return nil, errors.New("problem with data conversion")
		}
	}
	return b, err
}

// https://github.com/bastianrob/arrayutil/blob/bc57aac1ad39a96006bf13b9996e50b1894f9972/array.go
func (s *Helper) reduce(arr interface{}, initialValue interface{}, transform interface{}) (interface{}, error) {
	arrV := reflect.ValueOf(arr)
	kind := arrV.Kind()
	if kind != reflect.Slice && kind != reflect.Array {
		return nil, fmt.Errorf("Input value is not an array")
	}

	if transform == nil {
		return nil, fmt.Errorf("Transform function cannot be nil")
	}

	tv := reflect.ValueOf(transform)
	if tv.Kind() != reflect.Func {
		return nil, fmt.Errorf("Transform argument must be a function")
	}

	accV := reflect.ValueOf(initialValue)
	for i := 0; i < arrV.Len(); i++ {
		entry := arrV.Index(i)
		tfResults := tv.Call([]reflect.Value{accV, entry, reflect.ValueOf(i)})
		accV = tfResults[0]
	}

	return accV.Interface(), nil
}
