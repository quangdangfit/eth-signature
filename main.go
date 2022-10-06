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
		1665029590,
	}
	pk := "0x6f66fe85a7b88402e30ed9e480568897177a59d185cbef375797053bcd9a12db"

	sign, err := signature.Sign(pk, types, values)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(sign)
}
