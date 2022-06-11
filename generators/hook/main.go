package hook

import (
	"fmt"

	"github.com/mheob/create-react-tsx-component/models"
)

func Run(opt *models.CmdOptionsModel) {
	opt.SetNames()

	// TODO: Remove the dummy output after the function is created
	fmt.Println("Creating Hook ...")
	fmt.Println("Name:", opt.Name)
	fmt.Println("FileName:", opt.FileName)
	fmt.Println("ReactName:", opt.ReactName)
	fmt.Println("Destination:", opt.Dest)
	fmt.Println("OnlyDryRun:", opt.OnlyDryRun)
	fmt.Println("WithTest:", opt.WithTest)
}
