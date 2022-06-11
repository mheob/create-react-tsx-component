package models

import (
	"os"
	"path"
	"text/template"
)

type TmplVars map[string]any

type Generator struct {
	tmplPath string
	dest     string
	vars     TmplVars
	options  *CmdOptionsModel
}

func NewGenerator(cmdType string, opt *CmdOptionsModel, vars TmplVars) *Generator {
	wd, err := os.Getwd()
	check(err)
	destPath := path.Join(wd, opt.Dest, opt.FileName)

	tmplPath := path.Join("generators", cmdType, "templates")

	return &Generator{tmplPath: tmplPath, dest: destPath, vars: vars, options: opt}
}

func (g *Generator) GenerateFiles(files []string) {
	for _, file := range files {
		g.generateFile(file)
	}
	g.generateFile("index")
}

func (g *Generator) generateFile(file string) {
	var extension string
	switch file {
	case "stories":
		extension = ".stories.tsx"
	case "test":
		extension = ".test.tsx"
	default:
		extension = ".tsx"
	}

	err := os.MkdirAll(g.dest, 0750)
	check(err)

	tmplFile, err := template.ParseFiles(path.Join(g.tmplPath, file+".tmpl"))
	check(err)

	fileName := g.options.FileName + extension
	if file == "index" {
		fileName = file + extension
	}

	f, err := os.Create(path.Join(g.dest, fileName))
	check(err)
	//goland:noinspection GoUnhandledErrorResult
	defer f.Close()

	err = tmplFile.Execute(f, g.vars)
	check(err)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
