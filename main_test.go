package main

import "testing"

func TestAddOneNumber(t *testing.T) {
	actualResult := Add(9)
	var expectedResult = 9

	if actualResult != expectedResult {
		t.Fatalf("Expected %d, but got %d", expectedResult, actualResult)
	}
}

func TestAddFourNumbers(t *testing.T) {
	actualResult := Add(1, 2, 3, 4)
	var expectedResult = 10

	if actualResult != expectedResult {
		t.Fatalf("Expected %d, but got %d", expectedResult, actualResult)
	}
}

func TestSubOneNumber(t *testing.T) {
	actualResult := Sub(9)
	var expectedResult = 9

	if actualResult != expectedResult {
		t.Fatalf("Expected %d, but got %d", expectedResult, actualResult)
	}
}

func TestSubFourNumbers(t *testing.T) {
	actualResult := Sub(10, 2, 3, 4)
	var expectedResult = 1

	if actualResult != expectedResult {
		t.Fatalf("Expected %d, but got %d", expectedResult, actualResult)
	}
}
