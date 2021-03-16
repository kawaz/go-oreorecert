package main

import (
	"fmt"

	"github.com/kawaz/go-oreorecert"
)

func main() {
	p := oreorecert.GetKeyPairOreoreNet()
	_, err := p.Certificate()
	if err != nil {
		panic(err)
	}
	fmt.Println(p.CertFile)
	fmt.Println(p.KeyFile)
}
