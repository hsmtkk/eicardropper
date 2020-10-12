package main

import (
	"io/ioutil"
	"log"
)

func main() {
	s := `X5O!P%@AP[4\PZX54(P^)7CC)7}$EICAR-STANDARD-ANTIVIRUS-TEST-FILE!$H+H*`
	if err := ioutil.WriteFile("eicar.dat", []byte(s), 0644); err != nil {
		log.Fatal(err)
	}
}
