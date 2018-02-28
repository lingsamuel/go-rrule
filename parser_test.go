package rrule_test

import (
	"time"

	. "github.com/lingsamuel/go-rrule"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("RRule Parser Test", func() {
	It("Should return correct result", func() {
		layout := "20060102T150405Z"
		t, err := time.Parse(layout, "20180711T070000Z")
		Expect(err).To(BeNil())
		rrule, err := Parse("RRULE:FREQ=WEEKLY;UNTIL=20180711T070000Z;BYDAY=TU,WE;INTERVAL=2;COUNT=3", nil)

		Expect(rrule.Freq).To(Equal(ParseFreq("WEEKLY")))
		Expect(rrule.Until.Unix()).To(Equal(t.Unix()))
		Expect(rrule.Until.Unix()).To(Equal(int64(1531292400)))
		Expect(rrule.Interval).To(Equal(2))
		Expect(rrule.Count).To(Equal(int64(3)))
		Expect(rrule.ByDay).To(Equal(ParseByDayList("TU,WE")))
	})

	It("Should return nil when err", func() {
		errTest := func(rruleSrc string) {
			rrule, err := Parse(rruleSrc, nil)
			Expect(rrule).To(BeNil())
			Expect(err).NotTo(BeNil())
		}

		errTest("")
		errTest("INVALID")
		errTest("RRULE:")
		errTest("RRULE:FREQ=R")
		errTest("RRULE:FREQ=DAILY,")
		errTest("RRULE:FREQ=")
	})
})
