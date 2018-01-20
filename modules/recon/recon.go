package recon

import (
	"fmt"
	//"net/http"
	//"io/ioutil"
	"log"
	"github.com/netm4ul/netm4ul/modules/recon/whois"
)

func checkErr(e error) {
	if e != nil {
		log.Fatal("Error : ", e)
	}
}

func Whois(domain string) {
	fmt.Println("----------Whois module module----------")

}

