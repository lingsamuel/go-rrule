package rrule_test

import (
	"github.com/lingsamuel/go-rrule"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("FreqParse", func() {
	It("Should parse", func() {
		Expect(rrule.ParseFreq("YEARLY")).To(Equal(rrule.Yearly))
		Expect(rrule.ParseFreq("")).To(Equal(rrule.NoRecurrence))
		Expect(rrule.ParseFreq("Invalid")).To(Equal(rrule.NoRecurrence))
	})
})
