package shodan

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"github.com/netm4ul/netm4ul/modules"
)

func Run(ip string, api_key string) (data string){
	fmt.Println("Shodan!")
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