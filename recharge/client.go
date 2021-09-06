package recharge

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

func New(appId string, params *xpay.RechargeParams) (*xpay.Recharge, error) {
	return getC().New(appId, params)
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
	err := c.B.Call("POST", fmt.Sprintf("/apps/%s/recharges", appId), c.Key, nil, paramsString, recharge)
	return recharge, err
}

func Get(appID, rechargeID string) (*xpay.Recharge, error) {
	return getC().Get(appID, rechargeID)
}

func (c Client) Get(appID, rechargeID string) (*xpay.Recharge, error) {
	recharge := &xpay.Recharge{}

	err := c.B.Call("GET", fmt.Sprintf("/apps/%s/recharges/%s", appID, rechargeID), c.Key, nil, nil, recharge)
	return recharge, err
}

func List(appID string, params *xpay.PagingParams) (*xpay.RechargeList, error) {
	return getC().List(appID, params)
}

func (c Client) List(appID string, params *xpay.PagingParams) (*xpay.RechargeList, error) {
	body := &url.Values{}
	params.Filters.AppendTo(body)
	rechargeList := &xpay.RechargeList{}

	err := c.B.Call("GET", fmt.Sprintf("/apps/%s/recharges", appID), c.Key, body, nil, rechargeList)
	return rechargeList, err
}

func getC() Client {
	return Client{xpay.GetBackend(xpay.APIBackend), xpay.Key}
}
