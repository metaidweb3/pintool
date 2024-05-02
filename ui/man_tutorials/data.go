package man_tutorials

import "fyne.io/fyne/v2/cmd/fyne_demo/tutorials"

var (
	// Tutorials defines the metadata for each tutorial
	Tutorials = map[string]tutorials.Tutorial{
		"welcome": {"Welcome", "", welcomeScreen, true},
		"metaID-Inscription": {"MetaID-Inscription",
			"Make MetaID pin.",
			metaidInscriptionScreen,
			true,
		},
	}

	// TutorialIndex  defines how our tutorials should be laid out in the index tree
	TutorialIndex = map[string][]string{
		"": {"welcome", "metaID-Inscription"},
	}
)
