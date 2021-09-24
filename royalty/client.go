package royalty

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

func (c Client) BatchUpdate(params *xpay.RoyaltyBatchUpdateParams) (*xpay.RoyaltyList, error) {
	paramsString, errs := xpay.JsonEncode(params)
	if errs != nil {
		if xpay.LogLevel > 0 {
			log.Printf("RoyaltyBatchUpdateParams Marshall Errors is : %q/n", errs)
		}
		return nil, errs
	}
	if xpay.LogLevel > 2 {
		log.Printf("params of create user is :\n %v\n ", string(paramsString))
	}

	royaltyList := &xpay.RoyaltyList{}
	err := c.backend.Call("PUT", "royalties", nil, paramsString, royaltyList)
	return royaltyList, err
}

func (c Client) Get(royaltyId string) (*xpay.Royalty, error) {
	royalty := &xpay.Royalty{}

	err := c.backend.Call("GET", fmt.Sprintf("/royalties/%s", royaltyId), nil, nil, royalty)
	return royalty, err
}

func (c Client) List(params *xpay.PagingParams) (*xpay.RoyaltyList, error) {
	body := &url.Values{}
	params.Filters.AppendTo(body)

	royaltyList := &xpay.RoyaltyList{}
	err := c.backend.Call("GET", "royalties", body, nil, royaltyList)
	return royaltyList, err
}
