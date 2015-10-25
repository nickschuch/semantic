package main

import (
	"fmt"
	"os"

	"github.com/gogits/git"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	cliBranch = kingpin.Flag("branch", "The branch to tag against.").Default("master").String()
	cliPatch  = kingpin.Flag("patch", "Patch version when you make backwards-compatible bug fixes.").Bool()
	cliMinor  = kingpin.Flag("minor", "Minor version when you add functionality in a backwards-compatible manner.").Bool()
	cliMajor  = kingpin.Flag("major", "Major version when you make incompatible API changes.").Bool()
)

func main() {
	kingpin.Version("0.0.1")
	kingpin.Parse()

	// Load the current repository.
	repo, err := git.OpenRepository(".git")
	if err != nil {
		fmt.Println("Cannot open the repository.")
		os.Exit(1)
	}

	// Now we get a list of all the tags.
	tags, err := repo.GetTags()
	if err != nil {
		fmt.Println("Cannot get tags.")
		os.Exit(1)
	}

	// Get the most recent tag.
	tag, err := filter(tags)
	if err != nil {
		fmt.Println("Cannot get the most recent tag.")
		os.Exit(1)
	}

	// Get the new tag based on the current tag.
	if *cliMinor {
		tag.BumpMinor()
	} else if *cliMajor {
		tag.BumpMajor()
	} else {
		tag.BumpPatch()
	}

	// Get the latest tag for the branch and assign the new tag.
	commit, err := repo.GetCommitOfBranch(*cliBranch)
	if err != nil {
		fmt.Println("Cannot get the latest commit from the " + *cliBranch + " branch.")
		os.Exit(1)
	}
	err = repo.CreateTag(tag.String(), commit.Id.String())
	if err != nil {
		fmt.Println("Cannot create tag.")
		os.Exit(1)
	}

	// Return the new tag so this can be used for scripting.
	fmt.Println(tag.String())
}

// Helper function to:
//  * Only semantic versions
//  * The most recent tag
func filter(tags []string) (Tag, error) {
	var tag Tag

	for _, t := range tags {
		// Convert the string into a Tag object so we can compare.
		tmp, err := CreateTag(t)
		if err != nil {
			// If we cannot convert a tag into a semanitc tag
			// we skip it eg. 123-topic is not a semanitc tag.
			continue
		}

		switch {
		case tmp.Major >= tag.Major:
			tag.Major = tmp.Major
		case tmp.Minor >= tag.Minor:
			tag.Minor = tmp.Minor
		case tmp.Patch >= tag.Patch:
			tag.Patch = tmp.Patch
		}
	}

	return tag, nil
}
