package component

import (
	"fmt"

	"github.com/mheob/create-react-tsx-component/models"
	"github.com/mheob/create-react-tsx-component/utils"
)

var name string
var dryRun bool
var storybook bool
var test bool

func Init(options *models.CmdOptions) {
	dryRun = options.DryRun
	storybook = options.Storybook
	test = options.Test

	fmt.Println("Creating Component ...")
	fmt.Println("name:", utils.ConvertFileName(options.Name, options.KebabCase))
	fmt.Println("dry-run:", dryRun)
	fmt.Println("storybook:", storybook)
	fmt.Println("test:", test)
}
