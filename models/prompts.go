package models

import (
	"fmt"
	"os"
	"strings"

	"github.com/manifoldco/promptui"
)

func TypePrompt(opt *CmdOptionsModel) {
	prompt := promptui.Select{
		Label: "What type of component do you want to create",
		Items: []string{"Component", "Hook", "Page"},
	}

	_, typeChoice, err := prompt.Run()

	checkPromptError(err)

	opt.Type = strings.ToLower(typeChoice)
}

func NamePrompt(opt *CmdOptionsModel) {
	prompt := promptui.Prompt{
		Label:    "What is the name of the component",
		Validate: validateEmpty("Name"),
	}

	name, err := prompt.Run()

	checkPromptError(err)

	opt.Name = name
}

func DestPrompt(opt *CmdOptionsModel) {
	prompt := promptui.Prompt{
		Label:    fmt.Sprintf(`Where do you want to create the "%s"`, opt.Type),
		Default:  fmt.Sprintf("./src/%ss", opt.Type),
		Validate: validateEmpty("Destination"),
	}

	dest, err := prompt.Run()

	checkPromptError(err)

	opt.Dest = dest
}

func UsesKebabCasePrompt(opt *CmdOptionsModel) {
	prompt := promptui.Select{
		Label: "Which \"Case\" do you want to use for the filename names",
		Items: []string{"PascalCase", "kebab-case"},
	}

	_, selectedCase, err := prompt.Run()

	checkPromptError(err)

	if selectedCase == "kebab-case" {
		opt.UsesKebabCase = true
	} else {
		opt.UsesKebabCase = false
	}
}

func WithTestPrompt(opt *CmdOptionsModel) {
	prompt := promptui.Select{
		Label: "Do you want to create a test file",
		Items: []string{"Yes", "No"},
	}

	_, shouldSkipTest, err := prompt.Run()

	checkPromptError(err)

	if shouldSkipTest == "Yes" {
		opt.WithTest = true
	} else {
		opt.WithTest = false
	}
}

func WithStorybookPrompt(opt *CmdOptionsModel) {
	prompt := promptui.Select{
		Label: "Do you want to create a stories file",
		Items: []string{"Yes", "No"},
	}

	_, shouldSkipStorybook, err := prompt.Run()

	checkPromptError(err)

	if shouldSkipStorybook == "Yes" {
		opt.WithStorybook = true
	} else {
		opt.WithStorybook = false
	}
}

func DirectoryExistPrompt() {
	prompt := promptui.Prompt{
		Label:     "Directory already exists, do you want to overwrite it",
		IsConfirm: true,
	}

	if _, err := prompt.Run(); err != nil {
		fmt.Println("Generation cancelled. Nothing was generated.")
		os.Exit(0)
	}
}

func checkPromptError(err error) {
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}
}

func validateEmpty(prompt string) func(input string) error {
	return func(input string) error {
		if input == "" {
			return fmt.Errorf("%s can't be empty", prompt)
		}
		return nil
	}
}
