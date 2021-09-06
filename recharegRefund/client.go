package recharegRefund

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

func New(appID, id string, params *xpay.RechargeRefundParams) (*xpay.Refund, error) {
	return getC().New(appID, id, params)
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
	err := c.B.Call("POST", fmt.Sprintf("/apps/%s/recharges/%s/refunds", appID, id), c.Key, nil, paramsString, rechargeRefund)
	return rechargeRefund, err
}

func Get(appID, rechargeID, refundID string) (*xpay.Refund, error) {
	return getC().Get(appID, rechargeID, refundID)
}

func (c Client) Get(appID, rechargeID, refundID string) (*xpay.Refund, error) {
	refund := &xpay.Refund{}

	err := c.B.Call("GET", fmt.Sprintf("/apps/%s/recharges/%s/refunds/%s", appID, rechargeID, refundID), c.Key, nil, nil, refund)
	return refund, err
}

func List(appID, rechargeID string, params *xpay.PagingParams) (*xpay.RefundList, error) {
	return getC().List(appID, rechargeID, params)
}

func (c Client) List(appID, rechargeID string, params *xpay.PagingParams) (*xpay.RefundList, error) {
	body := &url.Values{}
	params.Filters.AppendTo(body)
	rechargeRefundList := &xpay.RefundList{}

	err := c.B.Call("GET", fmt.Sprintf("/apps/%s/recharges/%s/refunds", appID, rechargeID), c.Key, body, nil, rechargeRefundList)
	return rechargeRefundList, err
}

func getC() Client {
	return Client{xpay.GetBackend(xpay.APIBackend), xpay.Key}
}
