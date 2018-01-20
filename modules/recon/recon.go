package recon

import (
	"fmt"
	"github.com/netm4ul/netm4ul/modules/recon/whois"
	"github.com/netm4ul/netm4ul/modules/recon/shodan"
)

func Shodan(domain string) (data string) {
	fmt.Println("----------Shodan module----------")
	api_key := "my_api_key"
	data = shodan.Run(domain, api_key)
	return
}

func Whois(domain string) (data string){
	fmt.Println("----------Whois module----------")
	data = whois.Run(domain)
	return
}

