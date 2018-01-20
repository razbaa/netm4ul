package recon

import (
	"fmt"
	"os"
	"github.com/netm4ul/netm4ul/modules/recon/whois"
	"github.com/netm4ul/netm4ul/modules/recon/shodan"
	"strings"
)

//Shodan call
func Shodan(domain string) (data string, err error) {
	//banner
	fmt.Println("----------Shodan module----------")
	//need shodan API (from netm4ul.conf)
	api_key := os.Getenv("SHODAN_API")
	fmt.Println("my api: " + api_key)
	//call shodan
	data = shodan.Run(domain, api_key)

	if strings.HasPrefix(data, "401") {
		err = fmt.Errorf("Unauthorized: Shodan API key is incorrect.")
		return
	}
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

