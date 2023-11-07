package main

import (
	cepgo "github.com/igormpb/cep-go"
)

func main() {

	for i := 0; i < 20; i++ {
		x, e := cepgo.CEP("01001000")
		if e != nil {
			println(e.Error())
		}
		println(x.City)
	}

}
