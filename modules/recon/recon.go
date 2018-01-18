package recon

import (
	"fmt"
	"net/http"
	"net"
	"io/ioutil"
	"time"
	//"os"
	//"strings
	"log"
)

func checkErr(e error) {
	if e != nil {
		log.Fatal("Error : ", e)
	}
}

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

func Whois_search(domain string)(string) {

	fmt.Println("Whois search incoming!!")

	ip, err := net.LookupIP(domain)
	checkErr(err)

	connection, err := net.DialTimeout("tcp", net.JoinHostPort("whois.apnic.net", "43"), time.Second * 5)
	checkErr(err)

	defer connection.Close()

	connection.Write([]byte(ip[0].String() + "\r\n"))

	buffer, err := ioutil.ReadAll(connection)
	checkErr(err)

	ret := string(buffer)

	return ret
}