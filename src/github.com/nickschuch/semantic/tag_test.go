package main

import (
  "testing"

	"github.com/stretchr/testify/assert"
)

func TestTag(t *testing.T) {
  var tag Tag

  assert.Contains(t, "0.0.0", tag.String(), "Initial tag.")

  tag.BumpPatch()
  assert.Contains(t, "0.0.1", tag.String(), "Bump patch.")

  tag.BumpMinor()
  assert.Contains(t, "0.1.0", tag.String(), "Bump minor.")

  tag.BumpMajor()
  assert.Contains(t, "1.0.0", tag.String(), "Bump major.")

  tag.BumpPatch()
  assert.Contains(t, "1.0.1", tag.String(), "Bump patch.")

  tag.BumpPatch()
  assert.Contains(t, "1.0.2", tag.String(), "Bump patch.")

  tag.BumpMinor()
  assert.Contains(t, "1.1.0", tag.String(), "Bump minor.")

  tag.BumpPatch()
  assert.Contains(t, "1.1.1", tag.String(), "Bump patch.")

  tag.BumpMajor()
  assert.Contains(t, "2.0.0", tag.String(), "Bump major.")
}

func TestCreateTag(t *testing.T) {
  tag, _ := CreateTag("1.2.3")
  assert.Contains(t, "1.2.3", tag.String(), "Can create tags from strings.")
}
