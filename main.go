package main

//import des modules
import (
	"fmt"
	"github.com/netm4ul/netm4ul/modules/recon"
	"github.com/netm4ul/netm4ul/modules"
)


//main function
func main() {
	//165.225.76.71 AFTI_ip
	//NETM4UL motherfucker
	fmt.Println("- NetM4ul -")
	//shodan test
	data, err := recon.Shodan("bruneau.fr")
	modules.Check_err(err)
	fmt.Println(data)
	//whois test
	//fmt.Println(recon.Whois("bruneau.fr"))
}