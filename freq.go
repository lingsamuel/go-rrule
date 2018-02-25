package rrule

type Freq int8

const (
	NoRecurrence = Freq(-1)
	Daily        = Freq(3)
	Weekly       = Freq(4)
	Monthly      = Freq(5)
	Yearly       = Freq(6)
)

var (
	freqMap = map[Freq]string{
		0:       "SECONDLY",
		1:       "MINUTELY",
		2:       "HOURLY",
		Daily:   "DAILY",
		Weekly:  "WEEKLY",
		Monthly: "MONTHLY",
		Yearly:  "YEARLY",
	}
)

func (f *Freq) String() string {
	return freqMap[*f]
}
