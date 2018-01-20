package shodan

import "fmt"

type Shodan struct {
	region_code string
	tags string
	ip string
	area_code string
	hostnames string
	contry_code string
}

func Parse (data string) (shodan_struct Shodan, err error) {
	fmt.Println("Yipi")
	return
}