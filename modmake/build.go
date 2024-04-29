package main

import (
	. "github.com/saylorsolutions/modmake"
)

func main() {
	b := NewBuild()
	b.Test().Does(Go().TestAll().Silent())
	b.Tools().DependsOnRunner("install-templ", "", Go().Install("github.com/a-h/templ/cmd/templ@latest"))
	b.Generate().DependsOnRunner("run-templ", "", Exec("templ", "generate"))

	docs := NewAppBuild("modmake-docs", "./cmd/modmake-docs", "0.1.0")
	docs.Build(func(gb *GoBuild) {
		gb.StripDebugSymbols()
	})
	docs.HostVariant().NoPackage()
	b.ImportApp(docs)
	b.Step("modmake-docs:install").DependsOn(b.Test())

	b.AddNewStep("run", "serves the docs locally", Go().Run("./cmd/modmake-docs", "serve",
		F("--base-path=${MM_BASE_PATH:/modmake}"),
		F("--latest-go=${MM_LATEST_GO:1.22}"),
		F("--latest-supported=${MM_LATEST_GO_SUPPORTED:1.20}"),
		F("--modmake-version=${MODMAKE_VERSION:v0.4.0}"),
	)).DependsOn(b.Generate())

	b.Execute()
}
