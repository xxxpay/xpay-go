package royaltySettlement

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

func New(params *xpay.RoyaltySettlementCreateParams) (*xpay.RoyaltySettlement, error) {
	return getC().New(params)
}

func (c Client) New(params *xpay.RoyaltySettlementCreateParams) (*xpay.RoyaltySettlement, error) {
	paramsString, errs := xpay.JsonEncode(params)
	if errs != nil {
		if xpay.LogLevel > 0 {
			log.Printf("UserParams Marshall Errors is : %q/n", errs)
		}
		return nil, errs
	}
	if xpay.LogLevel > 2 {
		log.Printf("params of create user is :\n %v\n ", string(paramsString))
	}

	royaltySettlement := &xpay.RoyaltySettlement{}
	err := c.B.Call("POST", "/royalty_settlements", c.Key, nil, paramsString, royaltySettlement)
	return royaltySettlement, err
}

func Get(royaltySettlementId string) (*xpay.RoyaltySettlement, error) {
	return getC().Get(royaltySettlementId)
}

func (c Client) Get(royaltySettlementId string) (*xpay.RoyaltySettlement, error) {
	royaltySettlement := &xpay.RoyaltySettlement{}

	err := c.B.Call("GET", fmt.Sprintf("/royalty_settlements/%s", royaltySettlementId), c.Key, nil, nil, royaltySettlement)
	return royaltySettlement, err
}

func Update(royaltySettlementId string, params xpay.RoyaltySettlementUpdateParams) (*xpay.RoyaltySettlement, error) {
	return getC().Update(royaltySettlementId, params)
}

func (c Client) Update(royaltySettlementId string, params xpay.RoyaltySettlementUpdateParams) (*xpay.RoyaltySettlement, error) {
	paramsString, _ := xpay.JsonEncode(params)
	if xpay.LogLevel > 2 {
		log.Printf("params of update RoyaltySettlement  to xpay is :\n %v\n ", string(paramsString))
	}

	royaltySettlement := &xpay.RoyaltySettlement{}

	err := c.B.Call("PUT", fmt.Sprintf("/royalty_settlements/%s", royaltySettlementId), c.Key, nil, paramsString, royaltySettlement)
	if err != nil {
		if xpay.LogLevel > 0 {
			log.Printf("%v\n", err)
		}
		return nil, err
	}
	return royaltySettlement, err
}

func List(params *xpay.PagingParams) (*xpay.RoyaltySettlementList, error) {
	return getC().List(params)
}

func (c Client) List(params *xpay.PagingParams) (*xpay.RoyaltySettlementList, error) {
	body := &url.Values{}
	params.Filters.AppendTo(body)

	royaltySettlementList := &xpay.RoyaltySettlementList{}
	err := c.B.Call("GET", "/royalty_settlements", c.Key, body, nil, royaltySettlementList)
	return royaltySettlementList, err
}
