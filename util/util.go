package util

import (
	"log"
	"strconv"
	"time"
)

// ParseTime converts string to time
func ParseTime(timestr string) time.Time {
	t, err := time.Parse("15:04", timestr)
	if err != nil {
		log.Fatalf("failed to parse time")
	}
	return t
}

// ValidateTime should be started and ended only in 15 minute intervals, i.e. XX:00, XX:15, XX:30, XX:45 and from 00:00 to 23:45
func ValidateTime(from time.Time, to time.Time) bool {
	if from.Before(to) &&
		from.Day() == to.Day() &&
		from.Minute()%15 == 0 &&
		to.Minute()%15 == 0 &&
		int(to.Sub(from).Minutes())%15 == 0 &&
		from.After(ParseTime("00:00")) || from.Equal(ParseTime("00:00")) &&
		to.Before(ParseTime("23:45")) || from.Equal(ParseTime("23:45")) {
		return true
	}
	return false
}

// ParseInt converts string to int
func ParseInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		log.Fatalf("failed to convert string to int")
	}
	return i
}
