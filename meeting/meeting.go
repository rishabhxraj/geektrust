package meeting

import (
	"geektrust/util"
	"math"
	"time"
)

// Meeting represents a meeting
type Meeting struct {
	StartTime time.Time
	EndTime   time.Time
	Capacity  int
}

// MeetingHall is the struct for meeting hall
type MeetingHall struct {
	Name        string
	MaxCapacity int
	Meetings    []*Meeting
}

const MaxCap = math.MaxInt16

//meetings is the array of meeting hall
var MeetingHalls []*MeetingHall

//bufferTimes is used to store the buffer time
var bufferTimes []*Meeting

//Init initializes the meeting hall
func Init() {
	bufferTime1 := Meeting{StartTime: util.ParseTime("09:00"), EndTime: util.ParseTime("9:15"), Capacity: MaxCap}
	bufferTime2 := Meeting{StartTime: util.ParseTime("13:15"), EndTime: util.ParseTime("13:45"), Capacity: MaxCap}
	bufferTime3 := Meeting{StartTime: util.ParseTime("18:45"), EndTime: util.ParseTime("19:00"), Capacity: MaxCap}
	bufferTimes = append(bufferTimes, &bufferTime1, &bufferTime2, &bufferTime3)
	CCave := MeetingHall{Name: "C-Cave", MaxCapacity: 3, Meetings: bufferTimes}
	DTower := MeetingHall{Name: "D-Tower", MaxCapacity: 7, Meetings: bufferTimes}
	GMansion := MeetingHall{Name: "G-Mansion", MaxCapacity: 20, Meetings: bufferTimes}
	MeetingHalls = append(MeetingHalls, &CCave, &DTower, &GMansion)
}

// Book book the meeting hall with the given time slot
func (m *MeetingHall) Book(from time.Time, to time.Time, capacity int) {
	m.Meetings = append(m.Meetings, &Meeting{StartTime: from, EndTime: to, Capacity: capacity})
}

// IsAvailabe checks if the given time slot is available
func (m *MeetingHall) IsAvailabe(from time.Time, to time.Time) bool {
	for _, meeting := range m.Meetings {
		if meeting.StartTime.Before(to) && meeting.EndTime.After(from) {
			return false
		}
	}
	return true
}
