package rrule

import (
	"time"
)

// RRule Iterator iterate over time based on RRule
// First is self
type RecurrenceRuleIterator struct {
	RecurrenceRule *RecurrenceRule
	Start          time.Time
	End            time.Time
}

// Clone will deep copy current iterator
func (iter *RecurrenceRuleIterator) Clone() *RecurrenceRuleIterator {
	return &RecurrenceRuleIterator{
		Start:          iter.Start,
		End:            iter.End,
		RecurrenceRule: (*RecurrenceRule).Clone(iter.RecurrenceRule),
	}
}

func (iter *RecurrenceRuleIterator) Next() *RecurrenceRuleIterator {
	if iter.RecurrenceRule == nil {
		return nil
	}
	until, freq, interval := iter.RecurrenceRule.Until, iter.RecurrenceRule.Freq, iter.RecurrenceRule.Interval

	if interval <= 0 {
		interval = 1
	}

	next := iter.Clone()
	dur := iter.End.Sub(iter.Start)

	if freq == Daily {
		// DAILY can directly add
		next.Start = iter.Start.AddDate(0, 0, interval*1)
		next.End = iter.End.AddDate(0, 0, interval*1)
	} else if freq == Weekly {
		if iter.RecurrenceRule.ByDay == nil {
			next.Start = iter.Start.AddDate(0, 0, interval*7)
			next.End = iter.End.AddDate(0, 0, interval*7)
		} else {
			// WEEKLY will ignore the ORDER of BYDAY
			// in RFC, it's an error(p41)
			// but here just ignore this error
			found := false
			for offset := 1; offset <= 7; offset++ {
				newStart := iter.Start.AddDate(0, 0, offset)
				if Contains(iter.RecurrenceRule.ByDay, newStart.Weekday()) {
					next.Start = iter.Start.AddDate(0, 0, interval*offset)
					next.End = iter.End.AddDate(0, 0, interval*offset)
					found = true
					break
				}
			}
			if !found {
				return nil
			}
		}
	} else if freq == Monthly {
		if iter.RecurrenceRule.ByDay == nil {
			next.Start = iter.Start.AddDate(0, interval*1, 0)
			next.End = iter.End.AddDate(0, interval*1, 0)
		} else {
			// first check if this month has valid occurrence
			// then check if the order is in 1-4

			thisMonth := iter.Start.Month()
			newStart := iter.Start

			found := false
			// make sure this month is empty
			for {
				newStart = newStart.AddDate(0, 0, 1)
				if newStart.Month() != thisMonth {
					newStart = time.Date(iter.Start.Year(), iter.Start.Month(),
						1, iter.Start.Hour(), iter.Start.Minute(), iter.Start.Second(),
						iter.Start.Nanosecond(), iter.Start.Location())
					newStart = newStart.AddDate(0, interval*1, 0)
					thisMonth = newStart.Month()
					break
				}
				if Satisfy(iter.RecurrenceRule.ByDay, newStart, Monthly) {
					next.Start = newStart
					next.End = newStart.Add(dur)
					found = true
					break
				}
			}

			if !found {
				// add FREQ
				//next.Start = iter.Start.AddDate(0, interval*1, 0)
				//next.End = iter.End.AddDate(0, interval*1, 0)

				for {
					if newStart.Month() != thisMonth {
						thisMonth = newStart.Month()
						// if next month not found, the RRule cannot find another satisfied case
						return nil
					}
					if Satisfy(iter.RecurrenceRule.ByDay, newStart, Monthly) {
						next.Start = newStart
						next.End = newStart.Add(dur)
						break
					}
					newStart = newStart.AddDate(0, 0, 1)
				}
			}
		}
	} else if freq == Yearly {
		if iter.RecurrenceRule.ByDay == nil {
			next.Start = iter.Start.AddDate(interval*1, 0, 0)
			next.End = iter.End.AddDate(interval*1, 0, 0)
		} else {
			thisYear := iter.Start.Year()
			newStart := iter.Start

			found := false
			for {
				newStart = newStart.AddDate(0, 0, 1)
				if newStart.Year() != thisYear {
					newStart = time.Date(iter.Start.Year(), 1, 1,
						iter.Start.Hour(), iter.Start.Minute(), iter.Start.Second(),
						iter.Start.Nanosecond(), iter.Start.Location())
					newStart = newStart.AddDate(interval*1, 0, 0)
					thisYear = newStart.Year()
					break
				}
				if Satisfy(iter.RecurrenceRule.ByDay, newStart, Yearly) {
					next.Start = newStart
					next.End = newStart.Add(dur)
					found = true
					break
				}
			}

			if !found {
				// add FREQ
				//next.Start = iter.Start.AddDate(interval*1, 0, 0)
				//next.End = iter.End.AddDate(interval*1, 0, 0)

				for {
					if newStart.Year() != thisYear {
						thisYear = newStart.Year()
						return nil
					}
					if Satisfy(iter.RecurrenceRule.ByDay, newStart, Yearly) {
						next.Start = newStart
						next.End = newStart.Add(dur)
						break
					}
					newStart = newStart.AddDate(0, 0, 1)
				}
			}
		}
	} else {
		return nil
	}

	// only check COUNT > 0
	// equals 0 as infinity
	if next.RecurrenceRule.Count > 0 {
		next.RecurrenceRule.Count -= 1

		// because start time is first occurrence
		// so check COUNT after sub 1
		// RFC 5545 page 40
		if next.RecurrenceRule.Count <= 0 {
			return nil
		}
	}

	// Although RFC says that COUNT and UNTIL is conflict
	// but we can just check it, I don't know why they have to conflict
	// In this implement, occurrence should Satisfy UNTIL and COUNT both
	if until != nil {
		if next.Start.After(*until) {
			return nil
		}
	}

	return next
}
