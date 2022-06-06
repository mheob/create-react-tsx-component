package models

type CmdOptionsModel struct {
	Name string

	Dest             string
	OnlyDryRun       bool
	UsesKebabCase    bool
	WithoutStorybook bool
	WithoutTest      bool
}

var CmdOptions = new(CmdOptionsModel)
