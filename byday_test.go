package rrule_test

import (
	"time"

	"strings"

	"github.com/lingsamuel/go-rrule"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("RecurrenceRule", func() {
	It("Should be success to parse and toString", func() {

		checkFunc := func(rule string, order int, weekday time.Weekday, shouldSuccess bool) {
			result, err := rrule.ParseByDay(rule)
			if shouldSuccess {
				Expect(err).To(BeNil())
				Expect(result.OrderWeek).To(Equal(order))
				Expect(result.Weekday).To(Equal(weekday))
				Expect(rrule.ByDayListToString([]*rrule.ByDay{result})).
					To(Equal(strings.TrimPrefix(strings.TrimPrefix(rule, "+"), "0")))
			} else {
				Expect(err).NotTo(BeNil())
			}
		}
		checkFunc("2MO", 2, time.Monday, true)
		checkFunc("-2MO", -2, time.Monday, true)
		checkFunc("+2MO", 2, time.Monday, true)
		checkFunc("0MO", 0, time.Monday, true)
		checkFunc("MO", 0, time.Monday, true)
		checkFunc("88MO", 0, 0, false)
		checkFunc("2ZZ", 0, 0, false)
	})
})
