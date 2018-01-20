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
	//165.225.76.71 AFTI_ip
	fmt.Println(recon.whois.Run())
}