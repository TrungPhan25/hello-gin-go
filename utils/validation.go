package utils

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/google/uuid"
)

func VaildationRequired(filedName, value string) error {
	if value == "" {
		return fmt.Errorf("%s is required", filedName)
	}

	return nil
}

func ValidationStringLength(filedName, value string, min, max int) error {
	if len(value) < min || len(value) > max {
		return fmt.Errorf("%s must be between %d and %d characters long", filedName, min, max)
	}

	return nil
}

func ValidationRegex(filedName, value string, re *regexp.Regexp, errorMessage string) error {
	if !re.MatchString(value) {
		return fmt.Errorf("%s: %s", filedName, errorMessage)
	}
	return nil
}

func ValidationPositiveInt(filedName string, value string) (int, error) {
	v, err := strconv.Atoi(value)
	if err != nil || v <= 0 {
		return 0, fmt.Errorf("%s must be a positive integer", filedName)
	}

	return v, nil
}

func ValidationUUID(filedName, value string) (uuid.UUID, error) {
	uid, err := uuid.Parse(value)
	if err != nil {
		return uuid.Nil, fmt.Errorf("%s must be a valid UUID", filedName)
	}

	return uid, nil
}

func ValidationInList(filedName, value string, allow map[string]bool) error {
	if !allow[value] {
		return fmt.Errorf("%s must be one of: %v", filedName, keys(allow))
	}
	return nil
}

func keys(m map[string]bool) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
