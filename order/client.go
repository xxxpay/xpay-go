package order

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
	err := c.backend.Call("POST", "/orders", nil, paramsString, order)
	return order, err
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
	err := c.backend.Call("POST", fmt.Sprintf("/orders/%s/pay", id), nil, paramsString, order)
	return order, err
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
	err := c.backend.Call("PUT", "/orders/"+id, nil, paramsString, order)

	return order, err
}

func (c Client) Get(id string) (*xpay.Order, error) {
	order := &xpay.Order{}

	err := c.backend.Call("GET", "/orders/"+id, nil, nil, order)
	return order, err
}

func (c Client) List(params *xpay.PagingParams) (*xpay.OrderList, error) {
	body := &url.Values{}
	params.Filters.AppendTo(body)

	orderList := &xpay.OrderList{}
	err := c.backend.Call("GET", "/orders", body, nil, orderList)
	return orderList, err
}

func (c Client) Payment(orderID, chargeID string) (*xpay.Payment, error) {
	payment := &xpay.Payment{}

	err := c.backend.Call("GET", fmt.Sprintf("/orders/%s/payments/%s", orderID, chargeID), nil, nil, payment)
	return payment, err
}

func (c Client) PaymentList(orderID string, params *xpay.PagingParams) (*xpay.PaymentList, error) {
	body := &url.Values{}
	params.Filters.AppendTo(body)

	chargeList := &xpay.PaymentList{}
	err := c.backend.Call("GET", fmt.Sprintf("/orders/%s/payments", orderID), body, nil, chargeList)
	return chargeList, err
}
