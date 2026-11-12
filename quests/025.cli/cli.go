package cli

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func RunCLI(args []string) error {
	if len(args) < 2 {
		return errors.New("missing subcommand")
	}

	switch args[1] {
	case "count":
		if len(args) != 3 || args[2] == "" {
			return errors.New("usage: file-cli count <filename>")
		}

		return countLines(args[2])

	case "search":
		searchCmd := flag.NewFlagSet("search", flag.ContinueOnError)
		searchCmd.SetOutput(io.Discard)

		caseInsensitive := searchCmd.Bool(
			"case-insensitive",
			false,
			"ignore case when matching",
		)

		if err := searchCmd.Parse(args[2:]); err != nil {
			return err
		}

		remaining := searchCmd.Args()
		if len(remaining) != 2 || remaining[0] == "" || remaining[1] == "" {
			return errors.New("usage: file-cli search [--case-insensitive] <pattern> <filename>")
		}

		return searchFile(remaining[0], remaining[1], *caseInsensitive)

	default:
		return fmt.Errorf("unknown subcommand: %s", args[1])
	}
}

func countLines(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	count := 0

	for scanner.Scan() {
		count++
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	fmt.Printf("lines: %d\n", count)
	return nil
}

func searchFile(pattern, filename string, caseInsensitive bool) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	if caseInsensitive {
		pattern = strings.ToLower(pattern)
	}

	for scanner.Scan() {
		line := scanner.Text()
		searchLine := line

		if caseInsensitive {
			searchLine = strings.ToLower(line)
		}

		if strings.Contains(searchLine, pattern) {
			fmt.Println(line)
		}
	}

	return scanner.Err()
}
