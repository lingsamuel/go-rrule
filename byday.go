package rrule

import (
	"time"

	"unicode"

	"strings"
)

var weekdayToByDayMap = map[time.Weekday]string{
	time.Sunday:    "SU",
	time.Monday:    "MO",
	time.Tuesday:   "TU",
	time.Wednesday: "WE",
	time.Thursday:  "TH",
	time.Friday:    "FR",
	time.Saturday:  "SA",
}

var byDayToWeekdayMap = map[string]time.Weekday{
	"SU": time.Sunday,
	"MO": time.Monday,
	"TU": time.Tuesday,
	"WE": time.Wednesday,
	"TH": time.Thursday,
	"FR": time.Friday,
	"SA": time.Saturday,
}

type ByDay struct {
	OrderWeek int // 1 to 53
	Weekday   time.Weekday
}

type ByDayList = []*ByDay

func (byDay *ByDay) validate() (bool, error) {
	orderWeek := byDay.OrderWeek
	if orderWeek < 0 {
		orderWeek = -orderWeek
	}
	if !((orderWeek >= 1 && orderWeek <= 53) || orderWeek == 0) {
		return false, Errorf("Got invalid BYDAY order week: %v", byDay.OrderWeek)
	}

	_, ok := weekdayToByDayMap[byDay.Weekday]

	if !ok {
		return false, Errorf("Got invalid BYDAY weekday: %v", byDay.Weekday)
	}

	return true, nil
}

func ParseByDayList(byDayListSrc string) ByDayList {
	byDaySlice := strings.Split(byDayListSrc, ",")
	var byDayList ByDayList

	for _, v := range byDaySlice {
		byDay, err := ParseByDay(v)
		if err == nil {
			byDayList = append(byDayList, byDay)
		} else {
			//  ignore err byday
		}
	}

	if len(byDayList) == 0 {
		return nil
	}
	return byDayList
}

func ParseByDay(byDaySrc string) (*ByDay, error) {
	var byDay = new(ByDay)
	var orderWeek int
	var weekday string

	// Remove useless +
	if []rune(byDaySrc)[0] == '+' {
		byDaySrc = string(([]rune(byDaySrc))[1:])
	}

	// Parse ByDay
	byDayRunes := []rune(byDaySrc)
	weekdayIndex := 0

	// Parse Negative Number
	if byDayRunes[0] == '-' {
		if unicode.IsDigit(byDayRunes[1]) {
			// strconv returns an err, I hate it
			orderWeek = int(byDayRunes[1]) - int('0')
			if unicode.IsDigit(byDayRunes[2]) {
				orderWeek = orderWeek*10 + int(byDayRunes[2]) - int('0')
				weekdayIndex = 3
			} else {
				weekdayIndex = 2
			}
		}
		orderWeek = -orderWeek
	}

	// Parse Number
	if unicode.IsDigit(byDayRunes[0]) {
		orderWeek = int(byDayRunes[0]) - int('0')
		if unicode.IsDigit(byDayRunes[1]) {
			orderWeek = orderWeek*10 + int(byDayRunes[1]) - int('0')
			weekdayIndex = 2
		} else {
			weekdayIndex = 1
		}
	}

	// Convert string to time.Weekday
	weekday = byDaySrc[weekdayIndex:]

	// Generate and validate value
	byDay.OrderWeek = orderWeek
	byDayWeekday, ok := byDayToWeekdayMap[weekday]
	if ok {
		byDay.Weekday = byDayWeekday
	} else {
		return nil, Errorf("Got invalid BYDAY weekday: %v", weekday)
	}

	valid, err := byDay.validate()
	if valid {
		return byDay, nil
	}

	return nil, err
}

// Division weekday divide ByDayList to unlimitedWeekday and orderLimitedWeekday
func divisionWeekday(byDayList ByDayList) (ByDayList, ByDayList) {
	var unlimitedWeekday, limitedWeekday ByDayList

	for _, v := range byDayList {
		if v.OrderWeek == 0 {
			unlimitedWeekday = append(unlimitedWeekday, v)
		} else {
			limitedWeekday = append(limitedWeekday, v)
		}
	}
	return unlimitedWeekday, limitedWeekday
}

// Contains check if a weekday in a ByDayList and ignore order
func Contains(byDayList ByDayList, w time.Weekday) bool {
	for _, v := range byDayList {
		if v.Weekday == w {
			return true
		}
	}
	return false
}

// ContainsOrder check if a weekday in a ByDayList and ignore order
func ContainsOrder(byDayList ByDayList, order int) (bool, time.Weekday) {
	for _, v := range byDayList {
		if v.OrderWeek == order {
			return true, v.Weekday
		}
	}
	return false, time.Sunday
}

// Satisfy check if a time Satisfy a ByDayList
func Satisfy(byDayList ByDayList, w time.Time, freq Freq) bool {
	var getOrder, getNegativeOrder func(time.Time) int
	if freq == Monthly {
		getOrder = GetOrderInMonth
		getNegativeOrder = GetNegativeOrderInMonth
	} else if freq == Yearly {
		getOrder = GetOrderInYear
		getNegativeOrder = GetNegativeOrderInYear
	}

	weekday := w.Weekday()
	order := getOrder(w)
	negativeOrder := getNegativeOrder(w)

	for _, v := range byDayList {
		if v.Weekday == weekday {
			if v.OrderWeek < 0 && v.OrderWeek == negativeOrder {
				return true
			} else if v.OrderWeek == order {
				return true
			}
		}
	}
	return false
}
