package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	a := "12:01:00PM"
	b := "12:01:00AM"

	// Extract time in the format "HH:MM:SS" from string 'a'
	re := regexp.MustCompile(`\d{2}:\d{2}:\d{2}`)
	timeA := re.FindString(a)

	// Convert 'timeA' from 12-hour format to 24-hour format
	hoursA, minutesA, secondsA := splitTime(timeA)
	if strings.Contains(a, "PM") && hoursA != "12" {
		hoursA = strconv.Itoa(strToInt(hoursA) + 12)
	} else if strings.Contains(a, "AM") && hoursA == "12" {
		hoursA = "00"
	}

	// Extract time in the format "HH:MM:SS" from string 'b'
	timeB := re.FindString(b)

	// Convert 'timeB' from 12-hour format to 24-hour format
	hoursB, minutesB, secondsB := splitTime(timeB)
	if strings.Contains(b, "PM") && hoursB != "12" {
		hoursB = strconv.Itoa(strToInt(hoursB) + 12)
	} else if strings.Contains(b, "AM") && hoursB == "12" {
		hoursB = "00"
	}

	// Construct the final output strings
	outputA := fmt.Sprintf("%s:%s:%s", hoursA, minutesA, secondsA)
	outputB := fmt.Sprintf("%s:%s:%s", hoursB, minutesB, secondsB)

	fmt.Println("Output a:", outputA) // Output a: 12:01:00
	fmt.Println("Output b:", outputB) // Output b: 00:01:00
}

func splitTime(time string) (string, string, string) {
	parts := strings.Split(time, ":")
	return parts[0], parts[1], parts[2]
}

func strToInt(str string) int {
	num, _ := strconv.Atoi(str)
	return num
}
