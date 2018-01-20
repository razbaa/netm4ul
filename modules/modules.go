package modules

import (
	"log"
)

//error handler
func Check_err(e error) {
	if e != nil {
		//if error, log fatal
		log.Fatal("Error : ", e)
	}
}