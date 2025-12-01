package techpalace

import (
	"strings"
)

// WelcomeMessage returns a welcome message for the customer.
// Example: WelcomeMessage("Judy") => "Welcome to the Tech Palace, JUDY"
func WelcomeMessage(customer string) string {
	// Convert customer name to uppercase
	upperName := strings.ToUpper(customer)
	return "Welcome to the Tech Palace, " + upperName
}

// AddBorder adds a border of stars around the message.
// Example: AddBorder("Welcome!", 10) =>
// **********
// Welcome!
// **********
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	// Create a line of stars
	stars := strings.Repeat("*", numStarsPerLine)
	// Combine: stars line + message + stars line with newline
	return stars + "\n" + welcomeMsg + "\n" + stars
}

// CleanupMessage removes all stars (*) and trims leading/trailing whitespace
// Example:
// message := `
// **************************
// *    BUY NOW, SAVE 10%   *
// **************************
// `
// CleanupMessage(message) => "BUY NOW, SAVE 10%"
func CleanupMessage(oldMsg string) string {
	// Remove all stars
	noStars := strings.ReplaceAll(oldMsg, "*", "")
	// Trim leading and trailing whitespace
	cleaned := strings.TrimSpace(noStars)
	return cleaned
}
