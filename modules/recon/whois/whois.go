package whois

import (
	"fmt"
	"net"
	"time"
	"strings"
	"io/ioutil"
	"log"
)

func checkErr(e error) {
	if e != nil {
		log.Fatal("Error : ", e)
	}
}

//Check if passed domain is really a domain name
func check_domain(domain string) (zone string, err error) {
	//seperate domain with "." separator
	splited_domain := strings.Split(domain, ".")
	//domain have to be like 'x.domain.zone'
	if len(splited_domain) < 2 {
		//create error var
		err = fmt.Errorf("'domain' argument is not a domain: %s", domain)
		//no zone
		zone := ""
		//return
		return zone, err
	}
	//retrieve domain zone
	zone = splited_domain[len(splited_domain) - 1]
	//return
	return zone, err
}

//Get whois server to contact
func get_whois_server(zone string) (whois_server string, err error){
	//retrieve whois server from whois_servers.go
	whois_server, ok := servers[zone]
	//do we have whois server
	if !ok {
		//generate error message
		err := fmt.Errorf("No server for zone %s.", zone)
		//return error
		return whois_server, err
	}
	//return whois server
	return whois_server, err
}

//run function of whois.go
func Run(domain string) (result string) {
	//get zone
	zone, err := check_domain(domain)
	//handle error
	checkErr(err)
	//retrieve whois server
	whois_server, err := get_whois_server(zone)
	//connect to whois server on port 43 (timeout needed to get reply)
	connection, err := net.DialTimeout("tcp", net.JoinHostPort(whois_server, "43"), time.Second * 5)
	//handle error
	checkErr(err)
	//handle socket closure
	defer connection.Close()
	//send data to whois server
	connection.Write([]byte(domain + "\r\n"))
	//get data from whois server
	buffer, err := ioutil.ReadAll(connection)
	//handle error
	checkErr(err)
	//return whois data
	result = string(buffer[:])

	return
}