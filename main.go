package main

//import des modules
import (
	"fmt"
	"github.com/netm4ul/netm4ul/modules/recon"
)

//main function
func main() {
	//print from main package
	fmt.Println("- NetM4ul -")
	//call recon module function
	recon.Scan_output()
}
