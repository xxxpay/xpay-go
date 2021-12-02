// Package refund provides the /refunds APIs
package refund

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

func (c Client) New(ch string, params *xpay.RefundParams) (*xpay.Refund, error) {

	paramsString, errs := xpay.JsonEncode(params)

	if errs != nil {
		if xpay.LogLevel > 0 {
			log.Printf("RefundParams Marshall Errors is : %q\n", errs)
		}
		return nil, errs
	}
	if xpay.LogLevel > 2 {
		log.Printf("params of refund request to xpay is :\n %v\n ", string(paramsString))
	}
	refund := &xpay.Refund{}
	err := c.backend.Call("POST", fmt.Sprintf("/payments/%v/refunds", ch), nil, paramsString, refund)
	return refund, err
}

func (c Client) Get(chid string, reid string) (*xpay.Refund, error) {
	var body *url.Values
	body = &url.Values{}
	refund := &xpay.Refund{}
	err := c.backend.Call("GET", fmt.Sprintf("/payments/%v/refunds/%v", chid, reid), body, nil, refund)
	return refund, err
}

func (c Client) GetByTransactionNo(transactionNo string) (*xpay.Refund, error) {
	var body *url.Values
	body = &url.Values{}
	body.Add("transaction_no", transactionNo)
	refund := &xpay.Refund{}
	err := c.backend.Call("GET", "/refunds", body, nil, refund)
	return refund, err
}

func (c Client) List(appId string, params *xpay.RefundListParams) *Iter {
	if params == nil {
		params = &xpay.RefundListParams{}
	}
	params.Filters.AddFilter("app[id]", "", appId)

	body := &url.Values{}
	var lp *xpay.ListParams

	params.AppendTo(body)
	lp = &params.ListParams

	return &Iter{xpay.GetIter(lp, body, func(b url.Values) ([]interface{}, xpay.ListMeta, error) {
		list := &xpay.RefundList{}
		err := c.backend.Call("GET", "/refunds", &b, nil, list)

		ret := make([]interface{}, len(list.Values))
		for i, v := range list.Values {
			ret[i] = v
		}

		return ret, list.ListMeta, err
	})}
}

func (c Client) ListByPayment(paymentId string, params *xpay.RefundListParams) *Iter {
	if params == nil {
		params = &xpay.RefundListParams{}
	}

	body := &url.Values{}
	var lp *xpay.ListParams

	params.AppendTo(body)
	lp = &params.ListParams

	return &Iter{xpay.GetIter(lp, body, func(b url.Values) ([]interface{}, xpay.ListMeta, error) {
		list := &xpay.RefundList{}
		err := c.backend.Call("GET", fmt.Sprintf("/payments/%v/refunds", paymentId), &b, nil, list)

		ret := make([]interface{}, len(list.Values))
		for i, v := range list.Values {
			ret[i] = v
		}

		return ret, list.ListMeta, err
	})}
}

type Iter struct {
	*xpay.Iter
}

func (i *Iter) Refund() *xpay.Refund {
	return i.Current().(*xpay.Refund)
}
