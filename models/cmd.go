package models

type CmdOptions struct {
	Name      string
	Dest      string
	DryRun    bool
	KebabCase bool
	Storybook bool
	Test      bool
}
