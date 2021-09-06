package batchWithdrawal

import (
	"fmt"
	"log"
	"net/url"

	"github.com/xxxpay/xpay-go"
)

type Client struct {
	B   xpay.Backend
	Key string
}

func getC() Client {
	return Client{xpay.GetBackend(xpay.APIBackend), xpay.Key}
}

func Confirm(appId string, params *xpay.BatchWithdrawalParams) (*xpay.BatchWithdrawal, error) {
	return getC().Confirm(appId, params)
}
func (c Client) Confirm(appId string, params *xpay.BatchWithdrawalParams) (*xpay.BatchWithdrawal, error) {
	params.Status = "pending"
	paramsString, _ := xpay.JsonEncode(params)
	batchWithdrawal := &xpay.BatchWithdrawal{}
	err := c.B.Call("POST", fmt.Sprintf("/apps/%s/batch_withdrawals", appId), c.Key, nil, paramsString, batchWithdrawal)
	if err != nil {
		if xpay.LogLevel > 0 {
			log.Printf("Balance BatchWithdrawal error: %v\n", err)
		}
	}
	return batchWithdrawal, err
}

func Cancel(appId string, params *xpay.BatchWithdrawalParams) (*xpay.BatchWithdrawal, error) {
	return getC().Cancel(appId, params)
}
func (c Client) Cancel(appId string, params *xpay.BatchWithdrawalParams) (*xpay.BatchWithdrawal, error) {
	params.Status = "canceled"
	paramsString, _ := xpay.JsonEncode(params)
	batchWithdrawal := &xpay.BatchWithdrawal{}
	err := c.B.Call("POST", fmt.Sprintf("/apps/%s/batch_withdrawals", appId), c.Key, nil, paramsString, batchWithdrawal)
	if err != nil {
		if xpay.LogLevel > 0 {
			log.Printf("Balance BatchWithdrawal error: %v\n", err)
		}
	}
	return batchWithdrawal, err
}

func Get(appId, batchWithdrawalId string) (*xpay.BatchWithdrawal, error) {
	return getC().Get(appId, batchWithdrawalId)
}
func (c Client) Get(appId, batchWithdrawalId string) (*xpay.BatchWithdrawal, error) {
	batchWithdrawal := &xpay.BatchWithdrawal{}
	err := c.B.Call("GET", fmt.Sprintf("/apps/%s/batch_withdrawals/%s", appId, batchWithdrawalId), c.Key, nil, nil, batchWithdrawal)
	if err != nil {
		if xpay.LogLevel > 0 {
			log.Printf("Get Balance BatchWithdrawal error: %v\n", err)
		}
	}
	return batchWithdrawal, err
}

func List(appId string, params *xpay.PagingParams) (*xpay.BatchWithdrawalList, error) {
	return getC().List(appId, params)
}
func (c Client) List(appId string, params *xpay.PagingParams) (*xpay.BatchWithdrawalList, error) {
	body := &url.Values{}
	params.Filters.AppendTo(body)

	batchWithdrawalList := &xpay.BatchWithdrawalList{}
	err := c.B.Call("GET", fmt.Sprintf("/apps/%s/batch_withdrawals", appId), c.Key, body, nil, batchWithdrawalList)
	if err != nil {
		if xpay.LogLevel > 0 {
			log.Printf("Get Balance BatchWithdrawal error: %v\n", err)
		}
	}
	return batchWithdrawalList, err
}
