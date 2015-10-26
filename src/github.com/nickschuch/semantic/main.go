package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	cliPatch = kingpin.Flag("patch", "Patch version when you make backwards-compatible bug fixes.").Bool()
	cliMinor = kingpin.Flag("minor", "Minor version when you add functionality in a backwards-compatible manner.").Bool()
	cliMajor = kingpin.Flag("major", "Major version when you make incompatible API changes.").Bool()
)

func main() {
	kingpin.Parse()

	out, err := exec.Command("git", "tag").Output()
	if err != nil {
		log.Fatal(err)
	}

	tags := strings.Split(string(out), "\n")
	tag, err := filter(tags)
	if err != nil {
		Exit("Cannot get the most recent tag.")
	}

	// Get the new tag based on the current tag.
	if *cliMinor {
		tag.BumpMinor()
	} else if *cliMajor {
		tag.BumpMajor()
	} else {
		tag.BumpPatch()
	}

	// Return the new tag so this can be used for scripting.
	fmt.Println(tag.String())
}
