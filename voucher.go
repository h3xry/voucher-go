package voucher

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type RedeemResponse struct {
	Status struct {
		Message string `json:"message"`
		Code    string `json:"code"`
	} `json:"status"`
	Data struct {
		Voucher struct {
			VoucherID          string `json:"voucher_id"`
			AmountBaht         string `json:"amount_baht"`
			RedeemedAmountBaht string `json:"redeemed_amount_baht"`
			Member             int    `json:"member"`
			Status             string `json:"status"`
			Link               string `json:"link"`
			Detail             string `json:"detail"`
			ExpireDate         int64  `json:"expire_date"`
			Redeemed           int    `json:"redeemed"`
			Available          int    `json:"available"`
		} `json:"voucher"`
		OwnerProfile struct {
			FullName string `json:"full_name"`
		} `json:"owner_profile"`
		RedeemerProfile any `json:"redeemer_profile"`
		MyTicket        any `json:"my_ticket"`
		Tickets         []struct {
			MobIle     string `json:"mob ile"`
			UpdateDate int64  `json:"update_date"`
			AmountBaht string `json:"amount_baht"`
			FullName   string `json:"full_name"`
		} `json:"tickets"`
	} `json:"data,omitempty"`
}

func Redeem(phoneNumber, voucherLink string) (*RedeemResponse, error) {
	// https://gift.truemoney.com/campaign/?v=415151b43e2f564c21864605ddb7da9c29Y
	voucherHash := strings.Replace(voucherLink, "https://gift.truemoney.com/campaign/?v=", "", -1)
	redeem_url := fmt.Sprintf("https://gift.truemoney.com/campaign/vouchers/%s/redeem", voucherHash)
	payload, _ := json.Marshal(map[string]string{"mobile": phoneNumber, "voucher_hash": voucherHash})
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				MinVersion: tls.VersionTLS13,
			},
		},
	}
	req, err := http.NewRequest("POST", redeem_url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	req.Header.Add("authority", "gift.truemoney.com")
	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Safari/537.36")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	ret := RedeemResponse{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}
