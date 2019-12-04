package main

import (
	"testing"
)

func TestHasDoubleDigit(t *testing.T) {
	if !HasDoubleDigit(111111) {
		t.Error("failed")
	}
	if !HasDoubleDigit(112345) {
		t.Error("failed")
	}
	if !HasDoubleDigit(123345) {
		t.Error("failed")
	}
	if !HasDoubleDigit(123455) {
		t.Error("failed")
	}
	if HasDoubleDigit(123456) {
		t.Error("failed")
	}
}
func TestHasDoubleDigitNotMore(t *testing.T) {
	if HasDoubleDigitNotMore(111111) {
		t.Error("failed")
	}
	if HasDoubleDigitNotMore(111112) {
		t.Error("failed")
	}
	if !HasDoubleDigitNotMore(112345) {
		t.Error("failed")
	}
	if !HasDoubleDigitNotMore(123345) {
		t.Error("failed")
	}
	if !HasDoubleDigitNotMore(133344) {
		t.Error("failed")
	}
	if !HasDoubleDigitNotMore(112222) {
		t.Error("failed")
	}
	if !HasDoubleDigitNotMore(111122) {
		t.Error("failed")
	}
	if HasDoubleDigitNotMore(123456) {
		t.Error("failed")
	}
	if !HasDoubleDigitNotMore(112233) {
		t.Error("failed")
	}
	if HasDoubleDigitNotMore(123444) {
		t.Error("failed")
	}
}

func TestHasIncreasingDigitOnly(t *testing.T) {
	if !HasIncreasingDigitOnly(111111) {
		t.Error("failed")
	}
	if !HasIncreasingDigitOnly(112345) {
		t.Error("failed")
	}
	if !HasIncreasingDigitOnly(123456) {
		t.Error("failed")
	}
	if HasIncreasingDigitOnly(123450) {
		t.Error("failed")
	}
	if HasIncreasingDigitOnly(121450) {
		t.Error("failed")
	}
}

func TestCountNumberOfPasswordBetween(t *testing.T) {
	if CountNumberOfPasswordBetween(111111, 111122) != 10 {
		t.Error("failed")
	}
	if CountNumberOfPasswordBetween(112344, 112345) != 2 {
		t.Error("failed")
	}
	if CountNumberOfPasswordBetween(123456, 123456) != 0 {
		t.Error("failed")
	}
	if CountNumberOfPasswordBetween(123450, 123455) != 1 {
		t.Error("failed")
	}
}
