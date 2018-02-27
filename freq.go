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
	FreqToStringMap = map[Freq]string{
		0:       "SECONDLY",
		1:       "MINUTELY",
		2:       "HOURLY",
		Daily:   "DAILY",
		Weekly:  "WEEKLY",
		Monthly: "MONTHLY",
		Yearly:  "YEARLY",
	}

	StringToFreqMap = map[string]Freq{
		"SECONDLY": 0,
		"MINUTELY": 1,
		"HOURLY":   2,
		"DAILY":    Daily,
		"WEEKLY":   Weekly,
		"MONTHLY":  Monthly,
		"YEARLY":   Yearly,
	}
)

func (f *Freq) String() string {
	return FreqToStringMap[*f]
}
