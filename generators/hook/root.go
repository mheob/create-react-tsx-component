package hook

import (
	"fmt"

	"github.com/mheob/create-react-tsx-component/models"
	"github.com/mheob/create-react-tsx-component/utils"
)

var name string
var dryRun bool
var test bool

func Init(options *models.CmdOptions) {
	name = options.Name
	dryRun = options.DryRun
	test = options.Test

	fmt.Println("Creating Hook ...")
	fmt.Println("name:", utils.ConvertFileName(options.Name, options.KebabCase))
	fmt.Println("dry-run:", dryRun)
	fmt.Println("test:", test)
}
