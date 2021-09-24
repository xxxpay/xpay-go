package recharegRefund

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

func (c Client) New(appID, id string, params *xpay.RechargeRefundParams) (*xpay.Refund, error) {
	paramsString, errs := xpay.JsonEncode(params)
	if errs != nil {
		if xpay.LogLevel > 0 {
			log.Printf("OrderRefundParams Marshall Errors is : %q/n", errs)
		}
		return nil, errs
	}
	if xpay.LogLevel > 2 {
		log.Printf("params of orderRefund  is :\n %v\n ", string(paramsString))
	}

	rechargeRefund := &xpay.Refund{}
	err := c.backend.Call("POST", fmt.Sprintf("/apps/%s/recharges/%s/refunds", appID, id), nil, paramsString, rechargeRefund)
	return rechargeRefund, err
}

func (c Client) Get(appID, rechargeID, refundID string) (*xpay.Refund, error) {
	refund := &xpay.Refund{}

	err := c.backend.Call("GET", fmt.Sprintf("/apps/%s/recharges/%s/refunds/%s", appID, rechargeID, refundID), nil, nil, refund)
	return refund, err
}

func (c Client) List(appID, rechargeID string, params *xpay.PagingParams) (*xpay.RefundList, error) {
	body := &url.Values{}
	params.Filters.AppendTo(body)
	rechargeRefundList := &xpay.RefundList{}

	err := c.backend.Call("GET", fmt.Sprintf("/apps/%s/recharges/%s/refunds", appID, rechargeID), body, nil, rechargeRefundList)
	return rechargeRefundList, err
}
