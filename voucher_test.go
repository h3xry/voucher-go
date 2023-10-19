package voucher

import (
	"fmt"
	"testing"
)

func TestRedeem(t *testing.T) {
	phoneNumber := "09999999999"
	voucherLink := "https://gift.truemoney.com/campaign/?v=1982376419827369830569i0c9uhKJHkl203"
	result, err := Redeem(phoneNumber, voucherLink)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(result)
}
