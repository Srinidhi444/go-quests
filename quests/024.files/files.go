package files

import (
	"errors"
	"os"
	"strings"
)

func ProcessFile() error {
	path := "task.txt"

	// Read the entire file.
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	// Split the content into words.
	words := strings.Fields(string(data))

	// Validate word count.
	if len(words) < 10 {
		return errors.New("task.txt must contain at least 10 words")
	}

	// Capitalize the first five words.
	for i := 0; i < 5; i++ {
		words[i] = strings.ToUpper(words[i])
	}

	// Lowercase the last five words.
	for i := len(words) - 5; i < len(words); i++ {
		words[i] = strings.ToLower(words[i])
	}

	// Join words with a single space and append newline.
	result := strings.Join(words, " ") + "\n"

	// Overwrite the original file.
	return os.WriteFile(path, []byte(result), 0644)
}
