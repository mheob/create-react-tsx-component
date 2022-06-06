package page

import (
	"fmt"

	"github.com/mheob/create-react-tsx-component/models"
	"github.com/mheob/create-react-tsx-component/utils"
)

var name string
var dryRun bool

func Init(options *models.CmdOptions) {
	name = options.Name
	dryRun = options.DryRun

	fmt.Println("Creating Page ...")
	fmt.Println("name:", utils.ConvertFileName(options.Name, options.KebabCase))
	fmt.Println("dry-run:", dryRun)
}
