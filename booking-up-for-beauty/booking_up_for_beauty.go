package booking

import (
	"time"
)

// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {
	//https://gist.github.com/unstppbl/26942512b3ca6a92857c87124445ca0b
	x, _ := time.Parse("1/02/2006 15:04:05", date)
	return x
}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {
	panic("Please implement the HasPassed function")
}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {
	panic("Please implement the IsAfternoonAppointment function")
}

// Description returns a formatted string of the appointment time
func Description(date string) string {
	panic("Please implement the Description function")
}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
	panic("Please implement the AnniversaryDate function")
}


func countBits(num uint32) int32 {
	n:= int64(num)
    var x string
    x= string(n)
	count:= 0
	for _ , v := range x{
		if v == '1' {
		count++
		}
	}
	return int32(count)


	m := make([]int, 1*2)
	m = append(m, v[0])
	m = append(m, v[1])
}