package main

import (
	. "github.com/saylorsolutions/modmake"
	"github.com/saylorsolutions/modmake/pkg/git"
)

const (
	modmakeVersion = "0.4.0"
)

func main() {
	b := NewBuild()
	b.Tools().DependsOnRunner("get-modmake", "", Go().Get("github.com/saylorsolutions/modmake@v"+modmakeVersion))
	b.Tools().DependsOnRunner("install-templ", "", Go().Install("github.com/a-h/templ/cmd/templ@latest"))
	b.Generate().DependsOnRunner("run-templ", "", Exec("templ", "generate"))
	b.Test().Does(Go().TestAll().Silent())

	docs := NewAppBuild("modmake-docs", "./cmd/modmake-docs", "0.1.0")
	docs.Build(func(gb *GoBuild) {
		gb.StripDebugSymbols()
	})
	docs.HostVariant().NoPackage()
	b.ImportApp(docs)
	b.Step("modmake-docs:install").
		DependsOn(b.Test()).
		BeforeRun(Print("Executing from commit hash: %s", git.CommitHash()))

	b.AddNewStep("run", "serves the docs locally", Go().Run("./cmd/modmake-docs", "serve",
		F("--base-path=${MM_BASE_PATH:/modmake}"),
		F("--latest-go=${MM_LATEST_GO:1.22}"),
		F("--latest-supported=${MM_LATEST_GO_SUPPORTED:1.20}"),
		F("--modmake-version=v"+modmakeVersion),
	)).DependsOn(b.Generate())

	b.Execute()
}
