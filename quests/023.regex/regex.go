package regex

import "regexp"

func IsOnlyNumbers(s string) bool {
	return regexp.MustCompile(`^[0-9]+$`).MatchString(s)
}

func IsOnlyAlphabets(s string) bool {
	return regexp.MustCompile(`^[a-zA-Z]+$`).MatchString(s)
}

func IsValidEmail(s string) bool {
	return regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`).MatchString(s)
}

func ContainsGoQuest(s string) bool {
	return regexp.MustCompile(`(?i)go-quest`).MatchString(s)
}

func IsValidUsername(s string) bool {
	return regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_]{2,15}$`).MatchString(s)
}
