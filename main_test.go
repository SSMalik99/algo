package main

import (
	"testing"
)

func TestDigitInNum(t *testing.T) {
	ints := map[int]int{
		1234:      4,
		1234567:   7,
		90:        2,
		78:        2,
		123456789: 9,
		12345678:  8,
		1:         1,
	}
	for key, val := range ints {
		numOfDigit := DigitInNumber(key)
		if numOfDigit != val {
			t.Errorf("Expected : %v,,,, got : %v", val, numOfDigit)
		}
		
	}
}

func TestReverseNumber(t *testing.T) {
	ints := map[int]int{
		1234:    4321,
		1234567:   7654321,
		90:        9,
		78:        87,
		123456789: 987654321,
		12345678:  87654321,
		1:         1,
	}

	for key, val := range ints {
		reverseNum := reverseNumber(key)

		if reverseNum != val {
			t.Errorf("expected reverse %v , got %v", val, reverseNum)
		}

	}
}

func TestFindNumberInList(t *testing.T) {
	
	// To Do will do when i got the testcase
	t.Errorf("You don't have enough test cases!!!")
}