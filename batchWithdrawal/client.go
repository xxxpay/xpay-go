package batchWithdrawal

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

func (c Client) Confirm(appId string, params *xpay.BatchWithdrawalParams) (*xpay.BatchWithdrawal, error) {
	params.Status = "pending"
	paramsString, _ := xpay.JsonEncode(params)
	batchWithdrawal := &xpay.BatchWithdrawal{}
	err := c.backend.Call("POST", fmt.Sprintf("/apps/%s/batch_withdrawals", appId), nil, paramsString, batchWithdrawal)
	if err != nil {
		if xpay.LogLevel > 0 {
			log.Printf("Balance BatchWithdrawal error: %v\n", err)
		}
	}
	return batchWithdrawal, err
}

func (c Client) Cancel(appId string, params *xpay.BatchWithdrawalParams) (*xpay.BatchWithdrawal, error) {
	params.Status = "canceled"
	paramsString, _ := xpay.JsonEncode(params)
	batchWithdrawal := &xpay.BatchWithdrawal{}
	err := c.backend.Call("POST", fmt.Sprintf("/apps/%s/batch_withdrawals", appId), nil, paramsString, batchWithdrawal)
	if err != nil {
		if xpay.LogLevel > 0 {
			log.Printf("Balance BatchWithdrawal error: %v\n", err)
		}
	}
	return batchWithdrawal, err
}

func (c Client) Get(appId, batchWithdrawalId string) (*xpay.BatchWithdrawal, error) {
	batchWithdrawal := &xpay.BatchWithdrawal{}
	err := c.backend.Call("GET", fmt.Sprintf("/apps/%s/batch_withdrawals/%s", appId, batchWithdrawalId), nil, nil, batchWithdrawal)
	if err != nil {
		if xpay.LogLevel > 0 {
			log.Printf("Get Balance BatchWithdrawal error: %v\n", err)
		}
	}
	return batchWithdrawal, err
}

func (c Client) List(appId string, params *xpay.PagingParams) (*xpay.BatchWithdrawalList, error) {
	body := &url.Values{}
	params.Filters.AppendTo(body)

	batchWithdrawalList := &xpay.BatchWithdrawalList{}
	err := c.backend.Call("GET", fmt.Sprintf("/apps/%s/batch_withdrawals", appId), body, nil, batchWithdrawalList)
	if err != nil {
		if xpay.LogLevel > 0 {
			log.Printf("Get Balance BatchWithdrawal error: %v\n", err)
		}
	}
	return batchWithdrawalList, err
}
