package partyrobot

import (
	"fmt"
	"strconv"
)

// Welcome greets a person by name.
func Welcome(name string) string {
	x := "Welcome to my party, " + name + "!"
	// println(x)
	return x
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	y := "Happy birthday " + name + "! You are now " + strconv.Itoa(age) + " years old!"
	println(y)
	return y
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	x :=	fmt.Sprintf("Welcome to my party, %s!\nYou have been assigned to table %03d. Your table is %s, exactly %.1f meters from here.\nYou will be sitting next to %s.", name,table, direction, distance,neighbor)
	return x
}
