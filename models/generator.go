package models

import (
	"fmt"
	"os"
	"path"
	"text/template"
)

type TmplVars map[string]any

type File struct {
	Name      string
	Extension string
}

type Generator struct {
	tmplType string
	tmplPath string
	dest     string
	vars     TmplVars
	options  *CmdOptionsModel
}

func (g *Generator) GenerateFiles(files []File) {
	createDir(g.dest)

	for _, file := range files {
		g.generateFile(file.Name, file.Extension)
	}

	if g.tmplType != "page" {
		g.generateFile("index", ".ts")
	}
}

func (g *Generator) generateFile(file, extension string) {
	err := os.MkdirAll(g.dest, 0750)
	check(err)

	tmplFile, err := template.ParseFiles(path.Join(g.tmplPath, file+".tmpl"))
	check(err)

	fileName := g.options.FileName
	if file == "index" {
		fileName = file
	}

	f, err := os.Create(path.Join(g.dest, fileName+extension))
	check(err)
	//goland:noinspection GoUnhandledErrorResult
	defer f.Close()

	err = tmplFile.Execute(f, g.vars)
	check(err)
}

func NewGenerator(cmdType string, opt *CmdOptionsModel, vars TmplVars) *Generator {
	wd, _ := os.Getwd()
	destPath := path.Join(wd, opt.Dest, opt.FileName)
	tmplPath := path.Join("templates", cmdType)

	return &Generator{
		tmplType: cmdType,
		tmplPath: tmplPath,
		dest:     destPath,
		vars:     vars,
		options:  opt,
	}
}

func createDir(dirName string) {
	if _, err := os.Stat(dirName); !os.IsNotExist(err) {
		if shouldOverwrite := ShouldOverwritePrompt(); !shouldOverwrite {
			fmt.Println("Generation cancelled. Nothing was generated.")
			os.Exit(0)
		}

		err = os.RemoveAll(dirName)
		check(err)
	}

	if err := os.MkdirAll(dirName, 0755); err != nil {
		panic(err)
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
