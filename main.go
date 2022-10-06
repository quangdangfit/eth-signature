package main

import (
	"log"

	"github.com/quangdangfit/eth-signature/signature"
)

func main() {
	types := []string{"address", "address", "uint256", "bytes32", "uint256"}
	values := []interface{}{
		"0x4a047bccee30a27247f6429bd8ade94b932bbf69",
		"0x0631aa71dC522EfF8d13b00cc5dA8221567A01cA",
		"100000000000000000000",
		"0x4341af530580a883c6741c72cb8c03bd30949ca6d11d8e5ef4d152c2fd0cae44",
		"1675029590",
	}
	pk := "0xafdfdba96e737cd62cbd2e509035bd293759b2da070226656a289e85c70482bc"

	sign, err := signature.Sign(pk, types, values)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(sign)
}
