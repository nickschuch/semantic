package main

// Helper function to:
//  * Only semantic versions
//  * The most recent tag
func filter(refs []string) (Tag, error) {
	var tag Tag

	for _, r := range refs {
		tmp, err := CreateTag(r)
		if err != nil {
			continue
		}

		if tmp.Major > tag.Major {
			tag = tmp
			continue
		}

		if tmp.Minor > tag.Minor {
			tag = tmp
			continue
		}

		if tmp.Patch > tag.Patch {
			tag = tmp
			continue
		}
	}

	return tag, nil
}
