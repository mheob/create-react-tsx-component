package page

import (
	"fmt"

	"github.com/mheob/create-react-tsx-component/models"
)

func Run() {
	options := models.CmdOptions

	options.SetNames()

	// TODO: Remove the dummy output after the function is created
	fmt.Println("Creating Page ...")
	fmt.Println("Name:", options.Name)
	fmt.Println("FileName:", options.FileName)
	fmt.Println("ReactName:", options.ReactName)
	fmt.Println("Destination:", options.Dest)
	fmt.Println("OnlyDryRun:", options.OnlyDryRun)
}
