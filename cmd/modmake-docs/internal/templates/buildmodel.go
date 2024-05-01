package templates

import _ "embed"

//go:embed buildmodel/composability.md
var composability []byte

func BuildModelSection(params Params) *Section {
	return NewSection("build-model", "Build Model",
		BuildModel(),
	).
		AddSubSection("steps", "Steps", BuildModel_Steps()).
		AddSubSection("tasks", "Tasks", BuildModel_Tasks()).
		AddSubSection("invocation", "Invocation", BuildModel_Invocation(params)).
		AddSubSection("composability", "Build Composability", Markdown(composability)).
		AddSubSection("api-patterns", "API Patterns", BuildModel_APIPatterns())
}
