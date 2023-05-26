package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	a := "07:05:45PM"

	// Extract time in the format "HH:MM:SS" from string 'a'
	re := regexp.MustCompile(`\d{2}:\d{2}:\d{2}`)
	timeA := re.FindString(a)

	// Convert 'timeA' from 12-hour format to 24-hour format
	parts := strings.Split(timeA, ":")
	hoursA, _ := strconv.Atoi(parts[0])
	minutesA, _ := strconv.Atoi(parts[1])
	secondsA, _ := strconv.Atoi(parts[2])

	if strings.Contains(a, "PM") && hoursA != 12 {
		hoursA += 12
	} else if strings.Contains(a, "AM") && hoursA == 12 {
		hoursA = 0
	}

	// Construct the final output string
	outputA := fmt.Sprintf("%02d:%02d:%02d", hoursA, minutesA, secondsA)

	fmt.Println("Output a:", outputA) // Output a: 19:05:45
}
