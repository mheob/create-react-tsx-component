package models

import (
	"bytes"
	"fmt"
	"path"
	"sync"
	"text/template"
)

type TmplVars map[string]interface{}

type Generator struct {
	mutex    *sync.Mutex
	tmplPath string
	Vars     TmplVars
}

func GenerateTemplate(cmdType string) *Generator {
	tmplPath := path.Join("generators", cmdType, "templates")
	return &Generator{mutex: &sync.Mutex{}, tmplPath: tmplPath}
}

func (g *Generator) GenerateFile(file string, vars TmplVars) {
	g.mutex.Lock()

	tmplFile, err := template.ParseFiles(path.Join(g.tmplPath, file))
	if err != nil {
		panic(err)
	}

	var tmpl bytes.Buffer

	err = tmplFile.Execute(&tmpl, vars)
	if err != nil {
		panic(err)
	}

	newFile := tmpl.String()

	// TODO: Create files
	fmt.Println(newFile)

	g.mutex.Unlock()
}
