package page

import (
	"fmt"

	"github.com/mheob/create-react-tsx-component/models"
	"github.com/mheob/create-react-tsx-component/utils"
)

func Run() {
	fmt.Println("Creating Page ...")
	fmt.Println("Name:", utils.ConvertFileName(models.CmdOptions.Name, models.CmdOptions.UsesKebabCase))
	fmt.Println("Destination:", models.CmdOptions.Dest)
	fmt.Println("OnlyDryRun:", models.CmdOptions.OnlyDryRun)
}
