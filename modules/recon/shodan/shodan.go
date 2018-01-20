package shodan

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"github.com/netm4ul/netm4ul/modules"
	"net"
)


func get_ip_from_domain(domain string) (ip_str string) {
	ip, err := net.LookupIP(domain)
	modules.Check_err(err)
	ip_str = ip[0].String()
	fmt.Println(ip_str)
	return
}

func Run(domain string, api_key string) (data string){
	fmt.Println("Shodan!")
	ip := get_ip_from_domain(domain)
	//shodan api https://api.shodan.io/shodan/host/{ip}?key={YOUR_API_KEY}
	reply, err := http.Get("https://api.shodan.io/shodan/host/" + ip + "?key=" + api_key)
	//handle error
	modules.Check_err(err)
	//retrieve data
	temp, _ := ioutil.ReadAll(reply.Body)
	//convert data to string
	data = string(temp)
	return
}
