package templates

func GettingStartedSection(params Params) *Section {
	return NewSection("getting-started", "Getting started",
		GettingStarted(params),
	)
}
