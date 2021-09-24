package balanceTransfer

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

func (c Client) New(appId string, params *xpay.BalanceTransferParams) (*xpay.BalanceTransfer, error) {
	paramsString, _ := xpay.JsonEncode(params)
	if xpay.LogLevel > 2 {
		log.Printf("params of balance transfer to xpay is :\n %v\n ", string(paramsString))
	}
	balanceTransfer := &xpay.BalanceTransfer{}

	err := c.backend.Call("POST", fmt.Sprintf("/apps/%s/balance_transfers", appId), nil, paramsString, balanceTransfer)
	if err != nil {
		if xpay.LogLevel > 0 {
			log.Printf("Balance Transfer error: %v\n", err)
		}
	}
	return balanceTransfer, err
}

func (c Client) Get(appId, balanceTransferID string) (*xpay.BalanceTransfer, error) {
	balanceTransfer := &xpay.BalanceTransfer{}

	err := c.backend.Call("GET", fmt.Sprintf("/apps/%s/balance_transfers/%s", appId, balanceTransferID), nil, nil, balanceTransfer)
	return balanceTransfer, err
}

func (c Client) List(appID string, params *xpay.PagingParams) (*xpay.BalanceTransferList, error) {
	body := &url.Values{}
	params.Filters.AppendTo(body)
	balanceTransferList := &xpay.BalanceTransferList{}

	err := c.backend.Call("GET", fmt.Sprintf("/apps/%s/balance_transfers", appID), body, nil, balanceTransferList)
	return balanceTransferList, err
}
