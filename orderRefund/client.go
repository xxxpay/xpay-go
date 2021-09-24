package orderRefund

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
	err := c.backend.Call("POST", fmt.Sprintf("/orders/%s/order_refunds", id), nil, paramsString, orderRefund)
	return orderRefund, err
}

func (c Client) Get(orderId, refundId string) (*xpay.Refund, error) {
	refund := &xpay.Refund{}
	err := c.backend.Call("GET", fmt.Sprintf("/orders/%s/order_refunds/%s", orderId, refundId), nil, nil, refund)
	return refund, err
}

func (c Client) List(orderId string, params *xpay.PagingParams) (*xpay.RefundList, error) {
	body := &url.Values{}
	params.Filters.AppendTo(body)
	orderRefundList := &xpay.RefundList{}
	err := c.backend.Call("GET", fmt.Sprintf("/orders/%s/order_refunds", orderId), body, nil, orderRefundList)
	return orderRefundList, err
}
