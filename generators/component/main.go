package component

import (
	"fmt"

	"github.com/mheob/create-react-tsx-component/models"
)

var vars = make(models.TmplVars)

func Run() {
	models.CmdOptions.SetNames()

	vars["Name"] = models.CmdOptions.ReactName
	vars["FileName"] = models.CmdOptions.FileName
	vars["ShouldSkipStorybook"] = models.CmdOptions.ShouldSkipStorybook
	vars["ShouldSkipTest"] = models.CmdOptions.ShouldSkipTest

	g := models.GenerateTemplate("Component")

	// TODO: Generate asynchronous
	g.GenerateFile("component.tmpl", vars)

	if !models.CmdOptions.ShouldSkipStorybook {
		fmt.Printf("---------------------------------\n\n")
		g.GenerateFile("stories.tmpl", vars)
	}

	if !models.CmdOptions.ShouldSkipTest {
		fmt.Printf("---------------------------------\n\n")
		g.GenerateFile("test.tmpl", vars)
	}
}
