package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	tags := []string{
		"0.0.1",
		"0.0.2",
		"0.0.3",
		"0.1.0",
		"0.2.0",
		"0.2.100",
		"1.5.8",
		"2.5.8",
		"2.6.8",
		"2.7.0",
		"2.8.0",
		"2.9.0",
		"3.0.0",
	}
	tag, _ := filter(tags)
	assert.Equal(t, "3.0.0", tag.String(), "Can get the latest tag.")
}
