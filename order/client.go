package order

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

func New(params *xpay.OrderCreateParams) (*xpay.Order, error) {
	return getC().New(params)
}

func (c Client) New(params *xpay.OrderCreateParams) (*xpay.Order, error) {
	paramsString, errs := xpay.JsonEncode(params)
	if errs != nil {
		if xpay.LogLevel > 0 {
			log.Printf("OrderCreateParams Marshall Errors is : %q/n", errs)
		}
		return nil, errs
	}
	if xpay.LogLevel > 2 {
		log.Printf("params of create order is :\n %v\n ", string(paramsString))
	}

	order := &xpay.Order{}
	err := c.B.Call("POST", "/orders", c.Key, nil, paramsString, order)
	return order, err
}

func Pay(id string, params *xpay.OrderPayParams) (*xpay.Order, error) {
	return getC().Pay(id, params)
}

func (c Client) Pay(id string, params *xpay.OrderPayParams) (*xpay.Order, error) {
	paramsString, errs := xpay.JsonEncode(params)
	if errs != nil {
		if xpay.LogLevel > 0 {
			log.Printf("OrderPayParams Marshall Errors is : %q/n", errs)
		}
		return nil, errs
	}
	if xpay.LogLevel > 2 {
		log.Printf("params of order pay is :\n %v\n ", string(paramsString))
	}

	order := &xpay.Order{}
	err := c.B.Call("POST", fmt.Sprintf("/orders/%s/pay", id), c.Key, nil, paramsString, order)
	return order, err
}

func Cancel(user, id string) (*xpay.Order, error) {
	return getC().Cancel(user, id)
}
func (c Client) Cancel(user, id string) (*xpay.Order, error) {
	params := struct {
		Status string `json:"status"`
		User   string `json:"user"`
	}{
		Status: "canceled",
		User:   user,
	}
	paramsString, _ := xpay.JsonEncode(params)
	if xpay.LogLevel > 2 {
		log.Printf("params of cancel order  is :\n %v\n ", string(paramsString))
	}

	order := &xpay.Order{}
	err := c.B.Call("PUT", "/orders/"+id, c.Key, nil, paramsString, order)

	return order, err
}

func Get(id string) (*xpay.Order, error) {
	return getC().Get(id)
}

func (c Client) Get(id string) (*xpay.Order, error) {
	order := &xpay.Order{}

	err := c.B.Call("GET", "/orders/"+id, c.Key, nil, nil, order)
	return order, err
}

func List(params *xpay.PagingParams) (*xpay.OrderList, error) {
	return getC().List(params)
}
func (c Client) List(params *xpay.PagingParams) (*xpay.OrderList, error) {
	body := &url.Values{}
	params.Filters.AppendTo(body)

	orderList := &xpay.OrderList{}
	err := c.B.Call("GET", "/orders", c.Key, body, nil, orderList)
	return orderList, err
}

func Payment(orderID, chargeID string) (*xpay.Payment, error) {
	return getC().Payment(orderID, chargeID)
}

func (c Client) Payment(orderID, chargeID string) (*xpay.Payment, error) {
	payment := &xpay.Payment{}

	err := c.B.Call("GET", fmt.Sprintf("/orders/%s/payments/%s", orderID, chargeID), c.Key, nil, nil, payment)
	return payment, err
}

func PaymentList(orderID string, params *xpay.PagingParams) (*xpay.PaymentList, error) {
	return getC().PaymentList(orderID, params)
}

func (c Client) PaymentList(orderID string, params *xpay.PagingParams) (*xpay.PaymentList, error) {
	body := &url.Values{}
	params.Filters.AppendTo(body)

	chargeList := &xpay.PaymentList{}
	err := c.B.Call("GET", fmt.Sprintf("/orders/%s/payments", orderID), c.Key, body, nil, chargeList)
	return chargeList, err
}

func getC() Client {
	return Client{xpay.GetBackend(xpay.APIBackend), xpay.Key}
}
