package prompts

import (
	"fmt"
	"os"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/mheob/create-react-tsx-component/models"
)

func TypePrompt() {
	prompt := promptui.Select{
		Label: "What type of component do you want to create",
		Items: []string{"Component", "Hook", "Page"},
	}

	_, typeChoice, err := prompt.Run()

	if err != nil {
		printPromptError(err)
		os.Exit(1)
	}

	models.CmdOptions.Type = strings.ToLower(typeChoice)
}

func NamePrompt() {
	prompt := promptui.Prompt{
		Label: "What is the name of the component",
	}

	name, err := prompt.Run()

	if err != nil {
		printPromptError(err)
		os.Exit(1)
	}

	models.CmdOptions.Name = name
}

func DestPrompt() {
	prompt := promptui.Prompt{
		Label:   fmt.Sprintf(`Where do you want to create the "%s"`, models.CmdOptions.Type),
		Default: fmt.Sprintf("./%ss", models.CmdOptions.Type),
	}

	dest, err := prompt.Run()

	if err != nil {
		printPromptError(err)
		os.Exit(1)
	}

	models.CmdOptions.Dest = dest
}

func UsesKebabCasePrompt() {
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
		models.CmdOptions.UsesKebabCase = false
	} else {
		models.CmdOptions.UsesKebabCase = true
	}
}

func ShouldSkipTestPrompt() {
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
		models.CmdOptions.ShouldSkipTest = false
	} else {
		models.CmdOptions.ShouldSkipTest = true
	}
}

func ShouldSkipStorybookPrompt() {
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
		models.CmdOptions.ShouldSkipStorybook = false
	} else {
		models.CmdOptions.ShouldSkipStorybook = true
	}
}

func printPromptError(err error) {
	fmt.Printf("Prompt failed %v\n", err)
}
