package models

type CmdOptionsModel struct {
	Type string

	Name string

	Dest                string
	OnlyDryRun          bool
	UsesKebabCase       bool
	ShouldSkipStorybook bool
	ShouldSkipTest      bool
}

var CmdOptions = new(CmdOptionsModel)
