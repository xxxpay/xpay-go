package withdrawal

import (
	"fmt"
	"log"
	"net/url"

	"github.com/xxxpay/xpay-go"
)

type Client struct {
	backend xpay.Backend
}

func NewClient(backend xpay.Backend) Client {
	return Client{backend: backend}
}

func (c Client) New(appId string, params *xpay.WithdrawalParams) (*xpay.Withdrawal, error) {
	paramsString, _ := xpay.JsonEncode(params)
	withdrawal := &xpay.Withdrawal{}

	err := c.backend.Call("POST", fmt.Sprintf("/apps/%s/withdrawals", appId), nil, paramsString, withdrawal)
	if err != nil {
		if xpay.LogLevel > 0 {
			log.Printf("Balance Withdrawal error error: %v\n", err)
		}
	}
	return withdrawal, err
}

func (c Client) Get(appId, withdrawalId string) (*xpay.Withdrawal, error) {
	withdrawal := &xpay.Withdrawal{}

	err := c.backend.Call("GET", fmt.Sprintf("/apps/%s/withdrawals/%s", appId, withdrawalId), nil, nil, withdrawal)
	if err != nil {
		if xpay.LogLevel > 0 {
			log.Printf("Get BalanceWithdrawal error: %v\n", err)
		}
	}
	return withdrawal, err
}

//用户的余额提现列表
func (c Client) List(appId string, params *xpay.PagingParams) (*xpay.WithdrawalList, error) {
	var body *url.Values
	body = &url.Values{}
	params.Filters.AppendTo(body)
	withdrawalList := &xpay.WithdrawalList{}

	err := c.backend.Call("GET", fmt.Sprintf("/apps/%s/withdrawals", appId), body, nil, withdrawalList)
	if err != nil {
		if xpay.LogLevel > 0 {
			log.Printf("Get Withdrawal List error: %v\n", err)
		}
	}
	return withdrawalList, err
}

func (c Client) Cancel(appId, withdrawalId string) (*xpay.Withdrawal, error) {
	cancelParams := struct {
		Status string `json:"status"`
	}{
		Status: "canceled",
	}
	paramsString, _ := xpay.JsonEncode(cancelParams)

	withdrawal := &xpay.Withdrawal{}

	err := c.backend.Call("PUT", fmt.Sprintf("/apps/%s/withdrawals/%s", appId, withdrawalId), nil, paramsString, withdrawal)
	if err != nil {
		if xpay.LogLevel > 0 {
			log.Printf("Balance Withdrawal error error: %v\n", err)
		}
	}
	return withdrawal, err
}

func (c Client) Confirm(appId, withdrawalId string) (*xpay.Withdrawal, error) {
	confirmParams := struct {
		Status string `json:"status"`
	}{
		Status: "pending",
	}
	paramsString, _ := xpay.JsonEncode(confirmParams)

	withdrawal := &xpay.Withdrawal{}

	err := c.backend.Call("PUT", fmt.Sprintf("/apps/%s/withdrawals/%s", appId, withdrawalId), nil, paramsString, withdrawal)
	if err != nil {
		if xpay.LogLevel > 0 {
			log.Printf("Balance Withdrawal error: %v\n", err)
		}
	}
	return withdrawal, err
}
