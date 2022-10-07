package main

import (
	"log"

	"github.com/quangdangfit/eth-signature/signature"
)

func main() {
	types := []string{"address", "address", "uint256", "bytes32", "uint256"}
	values := []interface{}{
		"0x15f50afaf1ebc1679e7767c548c65a6c9c333ad2",
		"0x0631aa71dC522EfF8d13b00cc5dA8221567A01cA",
		"260963194988253700",
		"0x4726621f6668ebf199e4618250ac76410117f3e42aa0ad6c6e98d833fd402b04",
		"1665075116",
	}
	pk := "pk"

	sign, err := signature.Sign(pk, types, values)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(sign)
}
