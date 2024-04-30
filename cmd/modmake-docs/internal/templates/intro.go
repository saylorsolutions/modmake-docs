package templates

func IntroSection(params Params) *Section {
	return NewSection("introduction", "Introduction",
		Intro(),
	).AddSubSection("project-commitments", "Project Commitments",
		Intro_ProjectCommitments(params),
	).AddSubSection("getting-started", "Getting Started",
		Intro_GettingStarted(params),
	)
}
