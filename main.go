package main

import (
	"fmt"

	"github.com/jfrog/jfrog-cli-plugin-template/src/fileparser"

	"github.com/jfrog/jfrog-cli-core/v2/plugins"
	"github.com/jfrog/jfrog-cli-core/v2/plugins/components"
	"github.com/jfrog/jfrog-cli-plugin-template/commands"
)

func main() {
	plugins.PluginMain(getApp())
	// our test functions
	fileparser.TestParser()
	fileparser.ReadLinebyLine()
	fmt.Println(fileparser.IsExist("ERROR", "/Users/tataraov/apps/code-frog-cli/input/artifactory.log"))
}

func getApp() components.App {
	app := components.App{}
	app.Name = "hello-frog"
	app.Description = "Easily greet anyone."
	app.Version = "v0.1.0"
	app.Commands = getCommands()
	return app
}

func getCommands() []components.Command {
	return []components.Command{
		commands.GetHelloCommand()}
}
