package royaltySettlement

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
	err := c.backend.Call("POST", "/royalty_settlements", nil, paramsString, royaltySettlement)
	return royaltySettlement, err
}

func (c Client) Get(royaltySettlementId string) (*xpay.RoyaltySettlement, error) {
	royaltySettlement := &xpay.RoyaltySettlement{}

	err := c.backend.Call("GET", fmt.Sprintf("/royalty_settlements/%s", royaltySettlementId), nil, nil, royaltySettlement)
	return royaltySettlement, err
}

func (c Client) Update(royaltySettlementId string, params xpay.RoyaltySettlementUpdateParams) (*xpay.RoyaltySettlement, error) {
	paramsString, _ := xpay.JsonEncode(params)
	if xpay.LogLevel > 2 {
		log.Printf("params of update RoyaltySettlement  to xpay is :\n %v\n ", string(paramsString))
	}

	royaltySettlement := &xpay.RoyaltySettlement{}

	err := c.backend.Call("PUT", fmt.Sprintf("/royalty_settlements/%s", royaltySettlementId), nil, paramsString, royaltySettlement)
	if err != nil {
		if xpay.LogLevel > 0 {
			log.Printf("%v\n", err)
		}
		return nil, err
	}
	return royaltySettlement, err
}

func (c Client) List(params *xpay.PagingParams) (*xpay.RoyaltySettlementList, error) {
	body := &url.Values{}
	params.Filters.AppendTo(body)

	royaltySettlementList := &xpay.RoyaltySettlementList{}
	err := c.backend.Call("GET", "/royalty_settlements", body, nil, royaltySettlementList)
	return royaltySettlementList, err
}
