package book

import (
	"fmt"
	"geektrust/constants"
	"geektrust/meeting"
	"time"
)

// NewMeeting creates a new meeting and add it to the meeting list of the given time slot
func NewMeeting(from time.Time, to time.Time, capacity int) {
	flag := true
	if capacity < 2 && capacity > 20 {
		fmt.Println(constants.NoVacancy)
		flag = false
	}
	for _, MeetingHall := range meeting.MeetingHalls {
		if MeetingHall.IsAvailabe(from, to) && MeetingHall.MaxCapacity >= capacity {
			MeetingHall.Book(from, to, capacity)
			fmt.Println(MeetingHall.Name)
			return
		} else {
			flag = false

		}
	}
	if !flag {
		fmt.Println(constants.NoVacancy)
	}

}
