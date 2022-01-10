package pl_validator

import "testing"

func TestHelper_digit(t *testing.T) {
	helper := Helper{}
	result1, err1 := helper.digit("123456780", 2, 4)
	result2, err2 := helper.digit("123456780", 0, 8)
	result3, err3 := helper.digit("123456780", 0, 0)
	result4, err4 := helper.digit("123456780", 1, 0)
	result5, err5 := helper.digit("123456780", 1, 0)
	if err1 != nil {
		t.Errorf("Expected %v, got %v", result1, err1)
	}
	if result1 != 34 {
		t.Errorf("Expected %v, got %v", 34, result1)
	}
	if err2 != nil {
		t.Errorf("Expected %v, got %v", result2, err2)
	}
	if result2 != 12345678 {
		t.Errorf("Expected %v, got %v", 12345678, result2)
	}
	if err3 == nil {
		expectedError := "start and end value cannot be the same"
		if err3.Error() != expectedError {
			t.Errorf("Expected %v, got %v", expectedError, err3)
		}
		if result3 != -1 {
			t.Errorf("Expected %v, got %v", -1, result3)
		}
	}
	if err4 == nil {
		expectedError := "start value cannot be greater then the end one"
		if err4.Error() != expectedError {
			t.Errorf("Expected %v, got %v", expectedError, err4)
		}
		if result4 != -1 {
			t.Errorf("Expected %v, got %v", -1, result4)
		}
	}
	if err5 == nil {
		expectedError := "start end end value cannot be less than 0"
		if err5.Error() != expectedError {
			t.Errorf("Expected %v, got %v", expectedError, err5)
		}
		if result5 != -1 {
			t.Errorf("Expected %v, got %v", -1, result5)
		}
	}
}
