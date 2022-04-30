package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"geektrust/book"
	"geektrust/constants"
	"geektrust/meeting"
	"geektrust/util"
	"geektrust/vacancy"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Missing parameter, provide file name!")
		return
	}
	meeting.Init()
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("failed to open file")

	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, constants.VACANCY) {
			from := util.ParseTime(strings.Fields(line)[1])
			to := util.ParseTime(strings.Fields(line)[2])
			if util.ValidateTime(from, to) {
				meetinghalls := vacancy.CheckVacancy(from, to)
				if len(meetinghalls) == 0 {
					fmt.Println(constants.NoVacancy)
				} else {
					var availablehalls []string
					for _, meetinghall := range meetinghalls {
						availablehalls = append(availablehalls, meetinghall.Name)
					}
					fmt.Println(strings.Join(availablehalls, " "))
				}
			} else {
				fmt.Println(constants.IncorrectInput)
			}
		} else if strings.HasPrefix(line, constants.BOOK) {
			from := util.ParseTime(strings.Fields(line)[1])
			to := util.ParseTime(strings.Fields(line)[2])
			capacity := util.ParseInt(strings.Fields(line)[3])
			if util.ValidateTime(from, to) {
				book.NewMeeting(from, to, capacity)
			} else {
				fmt.Println(constants.IncorrectInput)
			}
		} else {
			fmt.Printf("Invalid Input: %s\n", line)
		}
	}

}
