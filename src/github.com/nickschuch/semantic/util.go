package main

import (
	"errors"
	"strconv"
	"strings"
)

// Helper function to convert a tag string into an int.
func TagToInt(t Tag) (int, error) {
	var tag int

	// Strip the dots out of the string.
	raw := strings.Replace(t.String(), ".", "", -1)

	// Conver to int.
	tag, err := strconv.Atoi(raw)
	if err != nil {
		return tag, errors.New("Cannot convert to integer.")
	}

	return tag, nil
}
