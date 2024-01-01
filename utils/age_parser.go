package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func ParseToDays(age string) (int, error) {
	parts := strings.Fields(age)
	if len(parts) != 2 {
		return 0, fmt.Errorf("invalid age format")
	}

	number, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, fmt.Errorf("invalid age number")
	}

	switch unit := parts[1]; unit {
	case "year", "years":
		return number * 365, nil
	case "month", "months":
		return number * 30, nil
	case "week", "weeks":
		return number * 7, nil
	case "day", "days":
		return number, nil
	default:
		return 0, fmt.Errorf("invalid time unit in age")
	}
}
