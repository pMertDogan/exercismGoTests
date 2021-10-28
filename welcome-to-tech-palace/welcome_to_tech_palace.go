package techpalace

import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(
		customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	var x string

	for i := 0; i < numStarsPerLine; i++ {
		x += "*"
	}
	s := []string{x, welcomeMsg, x}
	return strings.Join(s, "\n")
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {

	return strings.TrimSpace(strings.ReplaceAll(oldMsg, "*", ""))
}
