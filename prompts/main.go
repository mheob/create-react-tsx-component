package prompts

import (
	"fmt"
	"os"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/mheob/create-react-tsx-component/models"
)

func TypePrompt(opt *models.CmdOptionsModel) {
	prompt := promptui.Select{
		Label: "What type of component do you want to create",
		Items: []string{"Component", "Hook", "Page"},
	}

	_, typeChoice, err := prompt.Run()

	if err != nil {
		printPromptError(err)
		os.Exit(1)
	}

	opt.Type = strings.ToLower(typeChoice)
}

func NamePrompt(opt *models.CmdOptionsModel) {
	prompt := promptui.Prompt{
		Label: "What is the name of the component",
	}

	name, err := prompt.Run()

	if err != nil {
		printPromptError(err)
		os.Exit(1)
	}

	opt.Name = name
}

func DestPrompt(opt *models.CmdOptionsModel) {
	prompt := promptui.Prompt{
		Label:   fmt.Sprintf(`Where do you want to create the "%s"`, opt.Type),
		Default: fmt.Sprintf("./%ss", opt.Type),
	}

	dest, err := prompt.Run()

	if err != nil {
		printPromptError(err)
		os.Exit(1)
	}

	opt.Dest = dest
}

func UsesKebabCasePrompt(opt *models.CmdOptionsModel) {
	prompt := promptui.Select{
		Label: "Do you want to use \"PascalCase\" for the filename names",
		Items: []string{"Yes", "No"},
	}

	_, usesKebabCase, err := prompt.Run()

	if err != nil {
		printPromptError(err)
		os.Exit(1)
	}

	if usesKebabCase == "Yes" {
		opt.UsesKebabCase = false
	} else {
		opt.UsesKebabCase = true
	}
}

func WithTestPrompt(opt *models.CmdOptionsModel) {
	prompt := promptui.Select{
		Label: "Do you want to create a test file",
		Items: []string{"Yes", "No"},
	}

	_, shouldSkipTest, err := prompt.Run()

	if err != nil {
		printPromptError(err)
		os.Exit(1)
	}

	if shouldSkipTest == "Yes" {
		opt.WithTest = true
	} else {
		opt.WithTest = false
	}
}

func WithStorybookPrompt(opt *models.CmdOptionsModel) {
	prompt := promptui.Select{
		Label: "Do you want to create a stories file",
		Items: []string{"Yes", "No"},
	}

	_, shouldSkipStorybook, err := prompt.Run()

	if err != nil {
		printPromptError(err)
		os.Exit(1)
	}

	if shouldSkipStorybook == "Yes" {
		opt.WithStorybook = true
	} else {
		opt.WithStorybook = false
	}
}

func printPromptError(err error) {
	fmt.Printf("Prompt failed %v\n", err)
}
