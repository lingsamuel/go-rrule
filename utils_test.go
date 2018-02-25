package rrule_test

import (
	"testing"
	"time"
	"github.com/lingsamuel/go-rrule"
)

func TestGetOrderInMonth(t *testing.T) {
	testFunc := func(year, month, day, right int) {
		date := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
		result := rrule.GetOrderInMonth(date)
		if !(result == right) {
			t.Errorf("Should success: %v got order: %v, should be %v", date, result, right)
		}
	}

	testFunc(2018, 2, 18, 3)
	testFunc(2018, 3, 1, 1)
	testFunc(2018, 3, 25, 4)
	testFunc(2018, 3, 31, 5)
}

func TestGetOrderInYear(t *testing.T) {
	testFunc := func(year, month, day, right int) {
		date := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
		result := rrule.GetOrderInYear(date)
		if !(result == right) {
			t.Errorf("Should success: %v got order: %v, should be %v", date, result, right)
		}
	}

	testFunc(2018, 2, 18, 7)
	testFunc(2018, 3, 1, 9)
	testFunc(2018, 3, 25, 12)
	testFunc(2018, 3, 31, 13)
}

func TestGetNegativeOrderInMonth(t *testing.T) {
	testFunc := func(year, month, day, right int) {
		date := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
		result := rrule.GetNegativeOrderInMonth(date)
		if !(result == right) {
			t.Errorf("Should success: %v got order: %v, should be %v", date, result, right)
		}
	}

	testFunc(2018, 2, 25, -1)
	testFunc(2018, 2, 24, -1)
	testFunc(2018, 2, 23, -1)
	testFunc(2018, 2, 22, -1)
	testFunc(2018, 2, 21, -2)
	testFunc(2018, 2, 18, -2)
	testFunc(2018, 3, 1, -5)
	testFunc(2018, 3, 25, -1)
	testFunc(2018, 3, 31, -1)
}

func TestGetNegativeOrderInYear(t *testing.T) {
	testFunc := func(year, month, day, right int) {
		date := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
		result := rrule.GetNegativeOrderInYear(date)
		if !(result == right) {
			t.Errorf("Should success: %v got order: %v, should be %v", date, result, right)
		}
	}

	testFunc(2018, 12, 30, -1)
	testFunc(2018, 12, 31, -1)
	testFunc(2018, 12, 1, -5)
	testFunc(2018, 11, 26, -6)
	testFunc(2018, 1, 2, -52)
	testFunc(2018, 1, 1, -53)
}
