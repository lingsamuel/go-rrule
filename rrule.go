package rrule

import "time"

// In general, Freq defines the time section first
// Such as, FREQ=YEARLY will generate every year

// Secondly, interval is applied.
// Such as, INTERVAL=2 will generate every 2 year

// Thirdly, ByX(s) is(are) applied.
// Such as, BYDAY=SU, will generate all Sunday every 2 year

// Overall, count or until limit the generation


// Beware, Start time will be the first occurrence
// And before generate next FREQ, should check if there are occurrence Satisfy RRule in this FREQ

type RecurrenceRule struct {
	// Required
	Freq Freq

	// UNTIL and COUNT can only have 1
	Until time.Time // nil means forever
	Count int64      // 0 means non

	// BYX
	ByDay ByDayList

	// Optional
	Interval int
}

// Clone will deep copy current rrule reference
func (rrule *RecurrenceRule) Clone() *RecurrenceRule {
	if rrule == nil {
		return nil
	}
	r := *rrule
	return &r
}
