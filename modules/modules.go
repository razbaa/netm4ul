package modules

import (
	"log"
)

func Check_err(e error) {
	if e != nil {
		log.Fatal("Error : ", e)
	}
}