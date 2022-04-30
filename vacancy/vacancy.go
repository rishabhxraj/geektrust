package vacancy

import (
	"geektrust/meeting"
	"time"
)

// CheckVacancy return the array meeting hall which are vacant in the given time slot
func CheckVacancy(from time.Time, to time.Time) []meeting.MeetingHall {
	var meetingHalls []meeting.MeetingHall
	for _, MeetingHall := range meeting.MeetingHalls {
		if MeetingHall.IsAvailabe(from, to) {
			meetingHalls = append(meetingHalls, *MeetingHall)
		}
	}
	return meetingHalls
}
