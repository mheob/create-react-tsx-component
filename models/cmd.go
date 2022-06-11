package models

import "github.com/mheob/create-react-tsx-component/utils"

type CmdOptionsModel struct {
	Type string

	Name string

	Dest          string
	OnlyDryRun    bool
	UsesKebabCase bool
	WithStorybook bool
	WithTest      bool

	FileName  string
	ReactName string
}

func NewCmdOptions() *CmdOptionsModel {
	return &CmdOptionsModel{}
}

func (m *CmdOptionsModel) SetNames() {
	m.FileName = utils.ConvertFileName(m.Name, m.UsesKebabCase)
	m.ReactName = utils.ToPascalCase(m.Name)
}
