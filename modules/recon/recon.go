package recon

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func Scan_output() {
	fmt.Println("Scan module!")
}

func Shodan_search(ip string) {
	fmt.Println("Shodan!")
	//shodan api https://api.shodan.io/shodan/host/{ip}?key={YOUR_API_KEY}
	api_key := "eLrejd2wwMlxr3Ktxa0RR8u1qlnyauuk"
	reply, err := http.Get("https://api.shodan.io/shodan/host/" + ip + "?key=" + api_key)

	if err != nil{
		fmt.Println(err)
	} else {
		data, _ := ioutil.ReadAll(reply.Body)
		fmt.Println(string(data))
	}
}