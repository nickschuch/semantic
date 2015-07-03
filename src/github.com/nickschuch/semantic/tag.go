package main

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

type Tag struct {
	Major int
	Minor int
	Patch int
}

// Helper function to take a string and convert it to a tag.
func CreateTag(raw string) (Tag, error) {
	var tag Tag

	// Check if the raw value could become a semanitc tag.
	matched, err := regexp.MatchString("[0-9].[0-9].[0-9]", raw)
	if !matched || err != nil {
		return tag, errors.New("This is not a semantic tag.")
	}

	slice := strings.Split(raw, ".")

	major, err := strconv.Atoi(slice[0])
	if err != nil {
		return tag, errors.New("Cannot get the major value.")
	}

	minor, err := strconv.Atoi(slice[1])
	if err != nil {
		return tag, errors.New("Cannot get the minor value.")
	}

	patch, err := strconv.Atoi(slice[2])
	if err != nil {
		return tag, errors.New("Cannot get the patch value.")
	}

	tag = Tag{
		Major: major,
		Minor: minor,
		Patch: patch,
	}

	return tag, nil
}

// Helper function to bump the patch version.
//   eg. 1.1.0 to 1.1.1
func (t *Tag) BumpPatch() {
	t.Patch++
}

// Helper function to bump the minor version.
//   eg. 1.1.0 to 1.2.0
func (t *Tag) BumpMinor() {
	t.Minor++
	t.Patch = 0
}

// Helper function to bump the major version.
//   eg. 1.1.0 to 2.0.0
func (t *Tag) BumpMajor() {
	t.Major++
	t.Minor = 0
	t.Patch = 0
}

// Helper function to convert the tag into a string.
func (t *Tag) String() string {
	return strconv.Itoa(t.Major) + "." + strconv.Itoa(t.Minor) + "." + strconv.Itoa(t.Patch)
}
