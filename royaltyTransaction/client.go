package royaltyTransaction

import (
	"fmt"
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

func Get(royaltyTransacitonId string) (*xpay.RoyaltyTransaction, error) {
	return getC().Get(royaltyTransacitonId)
}

func (c Client) Get(royaltyTransacitonId string) (*xpay.RoyaltyTransaction, error) {
	royaltyTransaction := &xpay.RoyaltyTransaction{}

	err := c.B.Call("GET", fmt.Sprintf("/royalty_transactions/%s", royaltyTransacitonId), c.Key, nil, nil, royaltyTransaction)
	return royaltyTransaction, err
}

func List(params *xpay.PagingParams) (*xpay.RoyaltyTransactionList, error) {
	return getC().List(params)
}

func (c Client) List(params *xpay.PagingParams) (*xpay.RoyaltyTransactionList, error) {
	body := &url.Values{}
	params.Filters.AppendTo(body)

	royaltyTransactionList := &xpay.RoyaltyTransactionList{}
	err := c.B.Call("GET", "/royalty_transactions", c.Key, body, nil, royaltyTransactionList)
	return royaltyTransactionList, err
}
