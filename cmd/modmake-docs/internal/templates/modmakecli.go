package templates

func ModmakeCLISection(params Params) *Section {
	return NewSection("modmake-cli", "Modmake CLI", ModmakeCLI(params)).
		AddSubSection("resolution", "CLI Build Resolution", ModmakeCLI_Resolution()).
		AddSubSection("invocation", "CLI Invocation", ModmakeCLI_Invocation())
}
