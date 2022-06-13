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
	wd, _ := os.Getwd()
	destPath := path.Join(wd, opt.Dest, opt.FileName)
	tmplPath := path.Join("templates", cmdType)

	return &Generator{tmplPath: tmplPath, dest: destPath, vars: vars, options: opt}
}

func (g *Generator) GenerateFiles(files []string) {
	createDir(g.dest)

	for _, file := range files {
		g.generateFile(file)
	}
	g.generateFile("index")
}

func (g *Generator) generateFile(file string) {
	err := os.MkdirAll(g.dest, 0750)
	check(err)

	tmplFile, err := template.ParseFiles(path.Join(g.tmplPath, file+".tmpl"))
	check(err)

	fileName := g.options.FileName
	if file == "index" {
		fileName = file
	}

	f, err := os.Create(path.Join(g.dest, fileName+getExtension(file)))
	check(err)
	//goland:noinspection GoUnhandledErrorResult
	defer f.Close()

	err = tmplFile.Execute(f, g.vars)
	check(err)
}

func createDir(dirName string) {
	if _, err := os.Stat(dirName); os.IsNotExist(err) {
		if err := os.MkdirAll(dirName, 0755); err != nil {
			panic(err)
		}
	} else {
		DirectoryExistPrompt()
	}
}

func getExtension(file string) string {
	switch file {
	case "stories":
		return ".stories.tsx"
	case "test":
		return ".test.tsx"
	default:
		return ".tsx"
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
