package rrule

import (
	"strconv"
	"strings"
	"time"
)

func Parse(rrule string, loc *time.Location) (*RecurrenceRule, error) {
	if loc == nil {
		loc = time.UTC
	}
	prefix := "RRULE:"
	if !strings.HasPrefix(rrule, prefix) || rrule == prefix {
		return nil, Errorf("Empty or wrong RRule.")
	}

	if !strings.Contains(rrule, "FREQ=") {
		return nil, Errorf("FREQ does not exist!")
	}

	rrule = strings.TrimPrefix(rrule, prefix)
	res := &RecurrenceRule{}
	for _, kv := range strings.Split(rrule, ";") {
		parts := strings.Split(kv, "=")
		if len(parts) != 2 {
			continue
		}
		k := strings.ToUpper(parts[0])
		v := parts[1]
		if k == "FREQ" {
			if len(v) == 0 {
				return nil, Errorf("FREQ value does not exist!")
			}
			freq, ok := StringToFreqMap[v]
			if ok {
				res.Freq = freq
			} else {
				return nil, Errorf("FREQ value %v is invalid.", freq)
			}
		}
		if k == "UNTIL" {
			var (
				t   time.Time
				err error
			)

			if len(v) == 8 {
				layout := "20060102"
				t, err = time.ParseInLocation(layout, v, loc)
			} else {
				layout := "20060102T150405Z"
				t, err = time.Parse(layout, v)
			}

			if err != nil {
				return nil, err
			}
			res.Until = t
		}
		if k == "BYDAY" {
			res.ByDay = ParseByDayList(v)
		}
		if k == "COUNT" {
			val, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				return nil, err
			}
			res.Count = val
		}
		if k == "INTERVAL" {
			val, err := strconv.ParseInt(v, 10, 0)
			if err != nil {
				return nil, err
			}
			res.Interval = int(val)
		}
	}

	return res, nil
}
