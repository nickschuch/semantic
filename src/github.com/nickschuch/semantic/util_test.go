package main

import (
  "testing"

	"github.com/stretchr/testify/assert"
)

func TestTagToInt(t *testing.T) {
  tag := Tag{
    Major: 3,
    Minor: 2,
    Patch: 1,
  }
  number, _ := TagToInt(tag)
  assert.Equal(t, 321, number, "Convert to integer.")
}
