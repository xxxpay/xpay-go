package recharge

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

func (c Client) New(appId string, params *xpay.RechargeParams) (*xpay.Recharge, error) {
	paramsString, errs := xpay.JsonEncode(params)
	if errs != nil {
		if xpay.LogLevel > 0 {
			log.Printf("RechargeParams Marshall Errors is : %q/n", errs)
		}
		return nil, errs
	}
	if xpay.LogLevel > 2 {
		log.Printf("params of recharge request to xpay is :\n %v\n ", string(paramsString))
	}

	recharge := &xpay.Recharge{}
	err := c.backend.Call("POST", fmt.Sprintf("/apps/%s/recharges", appId), nil, paramsString, recharge)
	return recharge, err
}

func (c Client) Get(appID, rechargeID string) (*xpay.Recharge, error) {
	recharge := &xpay.Recharge{}

	err := c.backend.Call("GET", fmt.Sprintf("/apps/%s/recharges/%s", appID, rechargeID), nil, nil, recharge)
	return recharge, err
}

func (c Client) List(appID string, params *xpay.PagingParams) (*xpay.RechargeList, error) {
	body := &url.Values{}
	params.Filters.AppendTo(body)
	rechargeList := &xpay.RechargeList{}

	err := c.backend.Call("GET", fmt.Sprintf("/apps/%s/recharges", appID), body, nil, rechargeList)
	return rechargeList, err
}
