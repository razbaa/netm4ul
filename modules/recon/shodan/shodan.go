package shodan

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"github.com/netm4ul/netm4ul/modules"
	"net"
)

//get ip from domain name
func get_ip_from_domain(domain string) (ip_str string) {
	//retrieve ip
	ip, err := net.LookupIP(domain)
	//handle error
	modules.Check_err(err)
	//get string ip
	ip_str = ip[0].String()
	//return
	return
}

//Run fonction
func Run(domain string, api_key string) (data string){
	//retrieve ip from domain name
	ip := get_ip_from_domain(domain)
	//shodan api https://api.shodan.io/shodan/host/{ip}?key={YOUR_API_KEY}
	reply, err := http.Get("https://api.shodan.io/shodan/host/" + ip + "?key=" + api_key)
	//handle error
	modules.Check_err(err)
	//retrieve data
	temp, _ := ioutil.ReadAll(reply.Body)
	//convert data to string
	data = string(temp)
	//return
	return
}
