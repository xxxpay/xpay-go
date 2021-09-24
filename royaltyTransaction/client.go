package royaltyTransaction

import (
	"fmt"
	"net/url"

	"github.com/xxxpay/xpay-go"
)

type Client struct {
	backend xpay.Backend
}

func NewClient(backend xpay.Backend) Client {
	return Client{backend: backend}
}

func (c Client) Get(royaltyTransacitonId string) (*xpay.RoyaltyTransaction, error) {
	royaltyTransaction := &xpay.RoyaltyTransaction{}

	err := c.backend.Call("GET", fmt.Sprintf("/royalty_transactions/%s", royaltyTransacitonId), nil, nil, royaltyTransaction)
	return royaltyTransaction, err
}

func (c Client) List(params *xpay.PagingParams) (*xpay.RoyaltyTransactionList, error) {
	body := &url.Values{}
	params.Filters.AppendTo(body)

	royaltyTransactionList := &xpay.RoyaltyTransactionList{}
	err := c.backend.Call("GET", "/royalty_transactions", body, nil, royaltyTransactionList)
	return royaltyTransactionList, err
}
