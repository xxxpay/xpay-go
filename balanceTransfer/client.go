package balanceTransfer

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

func New(appId string, params *xpay.BalanceTransferParams) (*xpay.BalanceTransfer, error) {
	return getC().New(appId, params)
}

func (c Client) New(appId string, params *xpay.BalanceTransferParams) (*xpay.BalanceTransfer, error) {
	paramsString, _ := xpay.JsonEncode(params)
	if xpay.LogLevel > 2 {
		log.Printf("params of balance transfer to xpay is :\n %v\n ", string(paramsString))
	}
	balanceTransfer := &xpay.BalanceTransfer{}

	err := c.B.Call("POST", fmt.Sprintf("/apps/%s/balance_transfers", appId), c.Key, nil, paramsString, balanceTransfer)
	if err != nil {
		if xpay.LogLevel > 0 {
			log.Printf("Balance Transfer error: %v\n", err)
		}
	}
	return balanceTransfer, err
}

func Get(appId, balanceTransferId string) (*xpay.BalanceTransfer, error) {
	return getC().Get(appId, balanceTransferId)
}

func (c Client) Get(appId, balanceTransferID string) (*xpay.BalanceTransfer, error) {
	balanceTransfer := &xpay.BalanceTransfer{}

	err := c.B.Call("GET", fmt.Sprintf("/apps/%s/balance_transfers/%s", appId, balanceTransferID), c.Key, nil, nil, balanceTransfer)
	return balanceTransfer, err
}

func List(orderId string, params *xpay.PagingParams) (*xpay.BalanceTransferList, error) {
	return getC().List(orderId, params)
}

func (c Client) List(appID string, params *xpay.PagingParams) (*xpay.BalanceTransferList, error) {
	body := &url.Values{}
	params.Filters.AppendTo(body)
	balanceTransferList := &xpay.BalanceTransferList{}

	err := c.B.Call("GET", fmt.Sprintf("/apps/%s/balance_transfers", appID), c.Key, body, nil, balanceTransferList)
	return balanceTransferList, err
}
