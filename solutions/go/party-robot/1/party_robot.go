package partyrobot

import "fmt"

// Welcome greets a person by name.
func Welcome(name string) string {
	return fmt.Sprintf("Welcome to my party, %s!", name)
}

// HappyBirthday wishes happy birthday and mentions the age.
func HappyBirthday(name string, age int) string {
	return fmt.Sprintf("Happy birthday %s! You are now %d years old!", name, age)
}

// AssignTable assigns a table to each guest with formatted table number and distance.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	// Format table number to 3 digits: e.g., 7 -> 007
	tableFormatted := fmt.Sprintf("%03d", table)

	// Format distance to 1 decimal place
	distanceFormatted := fmt.Sprintf("%.1f", distance)

	return fmt.Sprintf(
		"Welcome to my party, %s!\nYou have been assigned to table %s. Your table is %s, exactly %s meters from here.\nYou will be sitting next to %s.",
		name, tableFormatted, direction, distanceFormatted, neighbor,
	)
}
