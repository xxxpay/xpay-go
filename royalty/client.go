package royalty

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

func BatchUpdate(params *xpay.RoyaltyBatchUpdateParams) (*xpay.RoyaltyList, error) {
	return getC().BatchUpdate(params)
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
	err := c.B.Call("PUT", "royalties", c.Key, nil, paramsString, royaltyList)
	return royaltyList, err
}

func Get(royaltyId string) (*xpay.Royalty, error) {
	return getC().Get(royaltyId)
}

func (c Client) Get(royaltyId string) (*xpay.Royalty, error) {
	royalty := &xpay.Royalty{}

	err := c.B.Call("GET", fmt.Sprintf("/royalties/%s", royaltyId), c.Key, nil, nil, royalty)
	return royalty, err
}

func List(params *xpay.PagingParams) (*xpay.RoyaltyList, error) {
	return getC().List(params)
}

func (c Client) List(params *xpay.PagingParams) (*xpay.RoyaltyList, error) {
	body := &url.Values{}
	params.Filters.AppendTo(body)

	royaltyList := &xpay.RoyaltyList{}
	err := c.B.Call("GET", "royalties", c.Key, body, nil, royaltyList)
	return royaltyList, err
}
