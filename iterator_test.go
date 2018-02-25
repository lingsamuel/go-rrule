package rrule_test

import (
	"time"

	"github.com/lingsamuel/go-rrule"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("RecurrenceRule", func() {
	Describe("RecurrenceRuleIterator", func() {
		It("Should return nil if rrule is nil", func() {
			now := time.Now()
			iterator := &rrule.RecurrenceRuleIterator{
				Start: now,
				End:   now.Add(time.Hour),
			}
			Expect(iterator.Next()).To(BeNil())
		})
		It("Should respect until", func() {
			start := time.Now()
			end := start.Add(time.Hour)
			until := start.AddDate(0, 0, 2)
			iterator := &rrule.RecurrenceRuleIterator{
				Start: start,
				End:   end,
				RecurrenceRule: &rrule.RecurrenceRule{
					Freq:  rrule.Daily,
					Until: &until,
				},
			}
			iterator = iterator.Next()
			Expect(iterator.Start.Equal(start.AddDate(0, 0, 1))).To(BeTrue())
			Expect(iterator.End.Equal(end.AddDate(0, 0, 1))).To(BeTrue())
			iterator = iterator.Next()
			Expect(iterator.Start.Equal(start.AddDate(0, 0, 2))).To(BeTrue())
			Expect(iterator.End.Equal(end.AddDate(0, 0, 2))).To(BeTrue())
			iterator = iterator.Next()
			Expect(iterator).To(BeNil())
		})
		It("Should respect weekly & byday", func() {
			start := time.Date(2017, 9, 13, 0, 0, 0, 0, time.UTC) // WE
			end := start.Add(time.Hour)
			iterator := &rrule.RecurrenceRuleIterator{
				Start: start,
				End:   end,
				RecurrenceRule: &rrule.RecurrenceRule{
					Freq:  rrule.Weekly,
					ByDay: rrule.ParseByDayList("TU,WE"),
				},
			}
			iterator = iterator.Next()
			Expect(iterator.Start.Equal(time.Date(2017, 9, 19, 0, 0, 0, 0, time.UTC))).To(BeTrue())
			Expect(iterator.End.Equal(time.Date(2017, 9, 19, 1, 0, 0, 0, time.UTC))).To(BeTrue())
			iterator = iterator.Next()
			Expect(iterator.Start.Equal(time.Date(2017, 9, 20, 0, 0, 0, 0, time.UTC))).To(BeTrue())
			Expect(iterator.End.Equal(time.Date(2017, 9, 20, 1, 0, 0, 0, time.UTC))).To(BeTrue())
			iterator = iterator.Next()
			Expect(iterator.Start.Equal(time.Date(2017, 9, 26, 0, 0, 0, 0, time.UTC))).To(BeTrue())
			Expect(iterator.End.Equal(time.Date(2017, 9, 26, 1, 0, 0, 0, time.UTC))).To(BeTrue())
		})
		It("Should respect monthly & byday", func() {
			start := time.Date(2017, 9, 13, 0, 0, 0, 0, time.UTC)
			end := start.Add(time.Hour)
			iterator := &rrule.RecurrenceRuleIterator{
				Start: start,
				End:   end,
				RecurrenceRule: &rrule.RecurrenceRule{
					Freq:  rrule.Monthly,
					ByDay: rrule.ParseByDayList("1WE"),
				},
			}
			iterator = iterator.Next()
			Expect(iterator.Start.Equal(time.Date(2017, 10, 4, 0, 0, 0, 0, time.UTC))).To(BeTrue())
			Expect(iterator.End.Equal(time.Date(2017, 10, 4, 1, 0, 0, 0, time.UTC))).To(BeTrue())
			iterator = iterator.Next()
			Expect(iterator.Start.Equal(time.Date(2017, 11, 1, 0, 0, 0, 0, time.UTC))).To(BeTrue())
			Expect(iterator.End.Equal(time.Date(2017, 11, 1, 1, 0, 0, 0, time.UTC))).To(BeTrue())

			start = time.Date(2017, 9, 21, 0, 0, 0, 0, time.UTC)
			end = start.Add(time.Hour)
			iterator = &rrule.RecurrenceRuleIterator{
				Start: start,
				End:   end,
				RecurrenceRule: &rrule.RecurrenceRule{
					Freq:  rrule.Monthly,
					ByDay: rrule.ParseByDayList("-2TU"),
				},
			}
			iterator = iterator.Next()
			Expect(iterator.Start.Equal(time.Date(2017, 10, 24, 0, 0, 0, 0, time.UTC))).To(BeTrue())
			Expect(iterator.End.Equal(time.Date(2017, 10, 24, 1, 0, 0, 0, time.UTC))).To(BeTrue())
			iterator = iterator.Next()
			Expect(iterator.Start.Equal(time.Date(2017, 11, 21, 0, 0, 0, 0, time.UTC))).To(BeTrue())
			Expect(iterator.End.Equal(time.Date(2017, 11, 21, 1, 0, 0, 0, time.UTC))).To(BeTrue())
		})
		It("Should respect count", func() {
			start := time.Date(2017, 9, 13, 0, 0, 0, 0, time.UTC)
			end := start.Add(time.Hour)
			iterator := &rrule.RecurrenceRuleIterator{
				Start: start,
				End:   end,
				RecurrenceRule: &rrule.RecurrenceRule{
					Freq:  rrule.Yearly,
					Count: 3,
				},
			}
			iterator = iterator.Next()
			Expect(iterator.Start.Equal(time.Date(2018, 9, 13, 0, 0, 0, 0, time.UTC))).To(BeTrue())
			Expect(iterator.End.Equal(time.Date(2018, 9, 13, 1, 0, 0, 0, time.UTC))).To(BeTrue())
			iterator = iterator.Next()
			Expect(iterator.Start.Equal(time.Date(2019, 9, 13, 0, 0, 0, 0, time.UTC))).To(BeTrue())
			Expect(iterator.End.Equal(time.Date(2019, 9, 13, 1, 0, 0, 0, time.UTC))).To(BeTrue())
			iterator = iterator.Next()
			Expect(iterator).To(BeNil())
		})
		//It("Should return correct HasIntersection", func() {
		//	start := time.Date(2017, 9, 13, 0, 0, 0, 0, time.UTC)
		//	end := start.Add(time.Hour)
		//	iterator := &RecurrenceRuleIterator{
		//		Start: start,
		//		End:   end,
		//		RecurrenceRule: &RecurrenceRule{
		//			Freq: Daily,
		//		},
		//	}
		//	Expect(iterator.HasIntersection(TimeRange{
		//		Left:       time.Date(2017, 9, 13, 0, 0, 0, 0, time.UTC),
		//		Right:      time.Date(2017, 10, 13, 0, 0, 0, 0, time.UTC),
		//		LeftClosed: true,
		//	})).To(BeTrue())
		//	Expect(iterator.HasIntersection(TimeRange{
		//		Left:       time.Date(2017, 9, 13, 1, 0, 0, 0, time.UTC),
		//		Right:      time.Date(2017, 10, 13, 0, 0, 0, 0, time.UTC),
		//		LeftClosed: true,
		//	})).To(BeFalse())
		//	// [2017-09-13 00:00:00Z, 2017-09-13 01:00:00Z) ∩ [2017-09-13 00:01:00Z, 2017-10-13 00:00:00Z) = ∅
		//	Expect(iterator.HasIntersection(TimeRange{
		//		Left:       time.Date(2017, 9, 13, 1, 0, 0, 1, time.UTC),
		//		Right:      time.Date(2017, 10, 13, 0, 0, 0, 0, time.UTC),
		//		LeftClosed: true,
		//	})).To(BeFalse())
		//	// [2017-09-13 00:00:00Z, 2017-09-13 01:00:00Z) ∩ (2017-09-13 00:01:00Z, 2017-10-13 00:00:00Z) = ∅
		//	Expect(iterator.HasIntersection(TimeRange{
		//		Left:  time.Date(2017, 9, 13, 1, 0, 0, 0, time.UTC),
		//		Right: time.Date(2017, 10, 13, 0, 0, 0, 0, time.UTC),
		//	})).To(BeFalse())
		//})
		//It("Should return correct next within range", func() {
		//	start := time.Date(2017, 9, 13, 0, 0, 0, 0, time.UTC)
		//	end := start.Add(time.Hour)
		//	iterator := &RecurrenceRuleIterator{
		//		Start: start,
		//		End:   end,
		//		RecurrenceRule: &RecurrenceRule{
		//			Freq: Daily,
		//		},
		//	}
		//	i1 := iterator.NextWithinRange(TimeRange{
		//		Left:       time.Date(2017, 9, 13, 0, 0, 0, 0, time.UTC),
		//		Right:      time.Date(2017, 10, 13, 0, 0, 0, 0, time.UTC),
		//		LeftClosed: true,
		//	})
		//	Expect(i1.Start.Equal(time.Date(2017, 9, 14, 0, 0, 0, 0, time.UTC))).To(BeTrue())
		//	Expect(i1.End.Equal(time.Date(2017, 9, 14, 1, 0, 0, 0, time.UTC))).To(BeTrue())
		//	i2 := iterator.NextWithinRange(TimeRange{
		//		Left:       time.Date(2017, 10, 13, 0, 0, 0, 0, time.UTC),
		//		Right:      time.Date(2017, 11, 13, 0, 0, 0, 0, time.UTC),
		//		LeftClosed: true,
		//	})
		//	Expect(i2.Start.Equal(time.Date(2017, 10, 13, 0, 0, 0, 0, time.UTC))).To(BeTrue())
		//	Expect(i2.End.Equal(time.Date(2017, 10, 13, 1, 0, 0, 0, time.UTC))).To(BeTrue())
		//})
		//It("Should return correct first within range", func() {
		//	start := time.Date(2017, 9, 13, 0, 0, 0, 0, time.UTC)
		//	end := start.Add(time.Hour)
		//	iterator := &RecurrenceRuleIterator{
		//		Start: start,
		//		End:   end,
		//		RecurrenceRule: &RecurrenceRule{
		//			Freq: Daily,
		//		},
		//	}
		//	i1 := iterator.FirstWithinRange(TimeRange{
		//		Left:  time.Date(2017, 9, 13, 1, 0, 0, 0, time.UTC),
		//		Right: time.Date(2017, 10, 13, 0, 0, 0, 0, time.UTC),
		//	})
		//	Expect(i1.Start.Equal(time.Date(2017, 9, 14, 0, 0, 0, 0, time.UTC))).To(BeTrue())
		//	Expect(i1.End.Equal(time.Date(2017, 9, 14, 1, 0, 0, 0, time.UTC))).To(BeTrue())
		//	i2 := iterator.FirstWithinRange(TimeRange{
		//		Left:  time.Date(2017, 10, 13, 0, 0, 0, 0, time.UTC),
		//		Right: time.Date(2017, 11, 13, 0, 0, 0, 0, time.UTC),
		//	})
		//	Expect(i2.Start.Equal(time.Date(2017, 10, 13, 0, 0, 0, 0, time.UTC))).To(BeTrue())
		//	Expect(i2.End.Equal(time.Date(2017, 10, 13, 1, 0, 0, 0, time.UTC))).To(BeTrue())
		//})

		It("Should be able to generate with interval", func() {
			start := time.Date(2017, 9, 13, 0, 0, 0, 0, time.UTC)
			end := start.Add(time.Hour)
			iterator := &rrule.RecurrenceRuleIterator{
				Start: start,
				End:   end,
				RecurrenceRule: &rrule.RecurrenceRule{
					Freq:     rrule.Daily,
					Interval: 2,
				},
			}
			iterator = iterator.Next()
			Expect(iterator.Start.Equal(time.Date(2017, 9, 15, 0, 0, 0, 0, time.UTC))).To(BeTrue())
			Expect(iterator.End.Equal(time.Date(2017, 9, 15, 1, 0, 0, 0, time.UTC))).To(BeTrue())
			iterator = iterator.Next()
			Expect(iterator.Start.Equal(time.Date(2017, 9, 17, 0, 0, 0, 0, time.UTC))).To(BeTrue())
			Expect(iterator.End.Equal(time.Date(2017, 9, 17, 1, 0, 0, 0, time.UTC))).To(BeTrue())
			iterator = iterator.Next()
			Expect(iterator.Start.Equal(time.Date(2017, 9, 19, 0, 0, 0, 0, time.UTC))).To(BeTrue())
			Expect(iterator.End.Equal(time.Date(2017, 9, 19, 1, 0, 0, 0, time.UTC))).To(BeTrue())
		})
		It("Should be able to generate with interval and count", func() {
			start := time.Date(2017, 9, 13, 0, 0, 0, 0, time.UTC)
			end := start.Add(time.Hour)
			iterator := &rrule.RecurrenceRuleIterator{
				Start: start,
				End:   end,
				RecurrenceRule: &rrule.RecurrenceRule{
					Freq:     rrule.Daily,
					Interval: 2,
					Count:    3,
				},
			}
			iterator = iterator.Next()
			Expect(iterator.Start.Equal(time.Date(2017, 9, 15, 0, 0, 0, 0, time.UTC))).To(BeTrue())
			Expect(iterator.End.Equal(time.Date(2017, 9, 15, 1, 0, 0, 0, time.UTC))).To(BeTrue())
			iterator = iterator.Next()
			Expect(iterator.Start.Equal(time.Date(2017, 9, 17, 0, 0, 0, 0, time.UTC))).To(BeTrue())
			Expect(iterator.End.Equal(time.Date(2017, 9, 17, 1, 0, 0, 0, time.UTC))).To(BeTrue())
			iterator = iterator.Next()
			Expect(iterator).To(BeNil())
		})
		It("Should be able to generate with interval and byday", func() {
			start := time.Date(2017, 1, 1, 0, 0, 0, 0, time.UTC) // TU
			end := start.Add(time.Hour)
			iterator := &rrule.RecurrenceRuleIterator{
				Start: start,
				End:   end,
				RecurrenceRule: &rrule.RecurrenceRule{
					Freq:     rrule.Monthly,
					ByDay:    rrule.ParseByDayList("1MO,2MO,3MO,4MO"),
					Interval: 2,
				},
			}

			iterator = iterator.Next()
			Expect(iterator.Start.Equal(time.Date(2017, 1, 2, 0, 0, 0, 0, time.UTC))).To(BeTrue())
			Expect(iterator.End.Equal(time.Date(2017, 1, 2, 1, 0, 0, 0, time.UTC))).To(BeTrue())
			iterator = iterator.Next()
			Expect(iterator.Start.Equal(time.Date(2017, 1, 9, 0, 0, 0, 0, time.UTC))).To(BeTrue())
			Expect(iterator.End.Equal(time.Date(2017, 1, 9, 1, 0, 0, 0, time.UTC))).To(BeTrue())
			iterator = iterator.Next()
			Expect(iterator.Start.Equal(time.Date(2017, 1, 16, 0, 0, 0, 0, time.UTC))).To(BeTrue())
			Expect(iterator.End.Equal(time.Date(2017, 1, 16, 1, 0, 0, 0, time.UTC))).To(BeTrue())
			iterator = iterator.Next()
			Expect(iterator.Start.Equal(time.Date(2017, 1, 23, 0, 0, 0, 0, time.UTC))).To(BeTrue())
			Expect(iterator.End.Equal(time.Date(2017, 1, 23, 1, 0, 0, 0, time.UTC))).To(BeTrue())

			iterator = iterator.Next()
			Expect(iterator.Start.Equal(time.Date(2017, 3, 6, 0, 0, 0, 0, time.UTC))).To(BeTrue())
			Expect(iterator.End.Equal(time.Date(2017, 3, 6, 1, 0, 0, 0, time.UTC))).To(BeTrue())
			iterator = iterator.Next()
			Expect(iterator.Start.Equal(time.Date(2017, 3, 13, 0, 0, 0, 0, time.UTC))).To(BeTrue())
			Expect(iterator.End.Equal(time.Date(2017, 3, 13, 1, 0, 0, 0, time.UTC))).To(BeTrue())
			iterator = iterator.Next()
			Expect(iterator.Start.Equal(time.Date(2017, 3, 20, 0, 0, 0, 0, time.UTC))).To(BeTrue())
			Expect(iterator.End.Equal(time.Date(2017, 3, 20, 1, 0, 0, 0, time.UTC))).To(BeTrue())
			iterator = iterator.Next()
			Expect(iterator.Start.Equal(time.Date(2017, 3, 27, 0, 0, 0, 0, time.UTC))).To(BeTrue())
			Expect(iterator.End.Equal(time.Date(2017, 3, 27, 1, 0, 0, 0, time.UTC))).To(BeTrue())
		})

		// TODO more complex test case
	})
})
