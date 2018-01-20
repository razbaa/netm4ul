package recon

import (
	"fmt"
	"github.com/netm4ul/netm4ul/modules/recon/whois"
	"github.com/netm4ul/netm4ul/modules/recon/shodan"
)

//Shodan call
func Shodan(domain string) (data string) {
	//banner
	fmt.Println("----------Shodan module----------")
	//need shodan API (from netm4ul.conf)
	api_key := "my_api_key"
	//call shodan
	data = shodan.Run(domain, api_key)
	//return
	return
}

//Whois call
func Whois(domain string) (data string){
	//banner
	fmt.Println("----------Whois module----------")
	//call whois
	data = whois.Run(domain)
	//return
	return
}

