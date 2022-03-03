package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {
	//   Mon Jan 2 15:04:05 MST 2006  (MST is GMT-0700)
	//   01 /02  03:04:05PM '06 -0700
	// 7/13/2020 20:32:00 
	x, _ := time.Parse("1/02/2006 15:04:05",date)
	// if e != nil{
	// 	fmt.Println(e.Error())
	// }
    
			return x
}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {
	x, _ := time.Parse("1/02/2006 15:04:05",date)
	fmt.Println(x)
   	return time.Now().After(x)
}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {
	//"Thursday, May 13, 2010 20:32:00"
	x, _ := time.Parse("Monday, January 2, 2006 15:4:5", date)
	fmt.Println(x)
	hour,_,_ := x.Clock()
	
	return hour > 12
}

// Description returns a formatted string of the appointment time
func Description(date string) string {
	//"9/19/1994 12:15:00"
	x, _ := time.Parse("1/2/2006 15:4:5", date)
	//"You have an appointment on Monday, June 6, 2005, at 10:30."
	return x.Format("You have an appointment on Monday, January 6, 2006, at 15:4.")
	//return x.Format("2006")
}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
	//yyyy-09-15
	return time.Date(time.Now().Year(), time.Month(9),15,0,0,0,0,time.UTC)
	
}
