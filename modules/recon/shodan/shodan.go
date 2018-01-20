package shodan

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func Shodan_search(ip string) {
	fmt.Println("Shodan!")
	//shodan api https://api.shodan.io/shodan/host/{ip}?key={YOUR_API_KEY}
	api_key := "my_api_key"
	reply, err := http.Get("https://api.shodan.io/shodan/host/" + ip + "?key=" + api_key)

	if err != nil{
		fmt.Println(err)
	} else {
		data, _ := ioutil.ReadAll(reply.Body)
		fmt.Println(string(data))
	}
}