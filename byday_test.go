package rrule_test

import (
	"testing"
	"github.com/lingsamuel/go-rrule"
)

func TestParseByDay(t *testing.T) {
	byDayList := map[string]bool{
		"2MO":  true,
		"-2MO": true,
		"+2MO": true,
		"0MO":  true,
		"88MO": false,
		"MO":   true,
		"8ZZ":  false,
	}
	for byDay, shouldSuccess := range byDayList {
		result, err := rrule.ParseByDay(byDay)
		if err != nil && shouldSuccess {
			t.Errorf("Should Success: ByDay: %v, Err: %v\n", byDay, err)
		} else if err == nil && !shouldSuccess {
			t.Errorf("Should Fail: ByDay: %v: Got: %v, %v\n", byDay, result.OrderWeek, result.Weekday)
		}
	}
}
