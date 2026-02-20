package utils

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/google/uuid"
)

func ValidationRequired(fieldName, value string) error {
	if value == "" {
		return fmt.Errorf("%s is required", fieldName)
	}

	return nil
}

func ValidationLength(fieldName, value string, min, max int) error {
	if len(value) < min {
		return fmt.Errorf("%s must be at least %d characters long", fieldName, min)
	}
	if len(value) > max {
		return fmt.Errorf("%s must be at most %d characters long", fieldName, max)
	}
	return nil
}

func ValidationRegex(fieldName, value, regex, errorMessage string) error {
	matched, err := regexp.MatchString(regex, value)
	if err != nil {
		return fmt.Errorf("Internal server error")
	}
	if !matched {
		return fmt.Errorf("%s %s", fieldName, errorMessage)
	}
	return nil
}

func ValidationPositiveInteger(fieldName, valueString string) (int, error) {
	value, err := strconv.ParseInt(valueString, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("%s must be a positive integer", fieldName)
	}
	if value <= 0 {
		return 0, fmt.Errorf("%s must be a positive integer", fieldName)
	}
	return int(value), nil
}

func ValidationUUID(fieldName, value string) (uuid.UUID, error) {
	u, err := uuid.Parse(value)
	if err != nil || u == uuid.Nil {
		return uuid.Nil, fmt.Errorf("%s must be a valid UUID", fieldName)
	}
	return u, nil
}

func ValidatuonInMap(fieldName, value string, validValues map[string]bool) error {
	if !validValues[value] {
		return fmt.Errorf("%s must be one of the following: %v", fieldName, getMapKeys(validValues))
	}
	return nil
}

func getMapKeys(m map[string]bool) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
