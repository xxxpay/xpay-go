package orderRefund

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

func New(id string, params *xpay.OrderRefundParams) (*xpay.RefundList, error) {
	return getC().New(id, params)
}

func (c Client) New(id string, params *xpay.OrderRefundParams) (*xpay.RefundList, error) {
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

	orderRefund := &xpay.RefundList{}
	err := c.B.Call("POST", fmt.Sprintf("/orders/%s/order_refunds", id), c.Key, nil, paramsString, orderRefund)
	return orderRefund, err
}

func Get(orderId, refundId string) (*xpay.Refund, error) {
	return getC().Get(orderId, refundId)
}

func (c Client) Get(orderId, refundId string) (*xpay.Refund, error) {
	refund := &xpay.Refund{}

	err := c.B.Call("GET", fmt.Sprintf("/orders/%s/order_refunds/%s", orderId, refundId), c.Key, nil, nil, refund)
	return refund, err
}

func List(orderId string, params *xpay.PagingParams) (*xpay.RefundList, error) {
	return getC().List(orderId, params)
}

func (c Client) List(orderId string, params *xpay.PagingParams) (*xpay.RefundList, error) {
	body := &url.Values{}
	params.Filters.AppendTo(body)
	orderRefundList := &xpay.RefundList{}

	err := c.B.Call("GET", fmt.Sprintf("/orders/%s/order_refunds", orderId), c.Key, body, nil, orderRefundList)
	return orderRefundList, err
}

func getC() Client {
	return Client{xpay.GetBackend(xpay.APIBackend), xpay.Key}
}
