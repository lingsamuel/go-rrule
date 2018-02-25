package rrule

import (
	"errors"
	"fmt"
	"time"
)

func AbsInt(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

// TODO get negative order

func GetOrderInMonth(t time.Time) int {
	deltaDays := t.Day() - 1
	deltaWeeks := (deltaDays / 7) + 1
	return deltaWeeks
}

func GetOrderInYear(t time.Time) int {
	firstDayInYear := time.Date(t.Year(), 1, 1, 0, 0, 0, 0, t.Location())

	deltaDays := t.Sub(firstDayInYear) / time.Hour / 24
	deltaWeeks := (int(deltaDays) / 7) + 1
	return deltaWeeks
}

func GetNegativeOrderInMonth(t time.Time) int {
	nextMonth := time.Date(t.Year(), t.Month()+1, 1, 0, 0, 0, 0, t.Location())
	thisMonthLastDate := nextMonth.AddDate(0, 0, -1)

	deltaDays := thisMonthLastDate.Sub(t) / time.Hour / 24
	deltaWeeks := (int(deltaDays) / 7) + 1
	return -deltaWeeks
}

func GetNegativeOrderInYear(t time.Time) int {
	nextYear := time.Date(t.Year()+1, 1, 1, 0, 0, 0, 0, t.Location())
	thisYearLastDate := nextYear.AddDate(0, 0, -1)

	deltaDays := thisYearLastDate.Sub(t) / time.Hour / 24
	deltaWeeks := (int(deltaDays) / 7) + 1
	return -deltaWeeks
}

func Errorf(format string, args ...interface{}) error {
	return errors.New(fmt.Sprintf(format, args...))
}
