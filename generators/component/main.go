package component

import (
	"github.com/mheob/create-react-tsx-component/models"
)

var vars = make(models.TmplVars)

func Run(opt *models.CmdOptionsModel) {
	opt.SetNames()

	vars["Name"] = opt.ReactName
	vars["FileName"] = opt.FileName
	vars["WithTest"] = opt.WithTest

	g := models.NewGenerator("component", opt, vars)

	files := make([]string, 1, 3)
	files[0] = "component"

	if opt.WithStorybook {
		files = append(files, "stories")
	}

	if opt.WithTest {
		files = append(files, "test")
	}

	g.GenerateFiles(files)
}
