package main

import (
	"fmt"

	"github.com/h3xry/voucher-go"
)

func main() {
	ret, err := voucher.Redeem("0999999999", "https://gift.truemoney.com/campaign/?v=1982376419827369830569i0c9uhKJHkl203")
	if err != nil {
		panic(err)
	}
	fmt.Println(ret)
}
