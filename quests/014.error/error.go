package error

import (
	"errors"
	"fmt"
	"strings"
)

// Part 1: Sentinel errors
var (
	ErrEmptyFilename = errors.New("filename cannot be empty")
	ErrFileTooLarge  = errors.New("file size exceeds limit")
)

// Part 2: Custom error type
type ValidationError struct {
	Filename string
	Reason   string
}

func (ve *ValidationError) Error() string {
	return fmt.Sprintf(
		"validation failed for '%s': %s",
		ve.Filename,
		ve.Reason,
	)
}

// Part 3: Validation functions

func ValidateFilename(filename string) error {
	if filename == "" {
		return ErrEmptyFilename
	}

	return nil
}

func ValidateFileSize(size int64, maxSize int64) error {
	if size < 0 {
		return errors.New("file size cannot be negative")
	}

	if size > maxSize {
		return ErrFileTooLarge
	}

	return nil
}

func ValidateFileExtension(filename string, allowedExts []string) error {
	for _, ext := range allowedExts {
		if strings.HasSuffix(filename, ext) {
			return nil
		}
	}

	return &ValidationError{
		Filename: filename,
		Reason:   "unsupported file extension",
	}
}

func ValidateFile(filename string, size int64, maxSize int64) error {
	if err := ValidateFilename(filename); err != nil {
		return fmt.Errorf("file validation failed: %w", err)
	}

	if err := ValidateFileSize(size, maxSize); err != nil {
		return fmt.Errorf("file validation failed: %w", err)
	}

	return nil
}

// Part 4: Error checking

func CanRetry(err error) bool {
	var validationErr *ValidationError

	return errors.As(err, &validationErr)
}
