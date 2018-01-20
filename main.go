package main

//import des modules
import (
	"fmt"
	"github.com/netm4ul/netm4ul/modules/recon"
)


//main function
func main() {
	//165.225.76.71 AFTI_ip
	//NETM4UL motherfucker
	fmt.Println("- NetM4ul -")
	//shodan test
	fmt.Println(recon.Shodan("bruneau.fr"))
	//whois test
	//fmt.Println(recon.Whois("bruneau.fr"))
}