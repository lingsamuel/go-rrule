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
	freqToStringMap = map[Freq]string{
		0:       "SECONDLY",
		1:       "MINUTELY",
		2:       "HOURLY",
		Daily:   "DAILY",
		Weekly:  "WEEKLY",
		Monthly: "MONTHLY",
		Yearly:  "YEARLY",
	}

	stringToFreqMap = map[string]Freq{
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
	return freqToStringMap[*f]
}

// TODO CHECK NoRecurrence usage
// ParseFreq input and return RecurrenceFreq, will return NoRecurrence if input is invalid or empty string
func ParseFreq(input string) Freq {
	freq, ok := stringToFreqMap[input]
	if ok {
		return freq
	}
	return NoRecurrence
}
