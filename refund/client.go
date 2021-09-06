// Package refund provides the /refunds APIs
package refund

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

func New(ch string, params *xpay.RefundParams) (*xpay.Refund, error) {
	return getC().New(ch, params)
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
	err := c.B.Call("POST", fmt.Sprintf("/charges/%v/refunds", ch), c.Key, nil, paramsString, refund)
	return refund, err
}

func Get(chid string, reid string) (*xpay.Refund, error) {
	return getC().Get(chid, reid)
}

func (c Client) Get(chid string, reid string) (*xpay.Refund, error) {
	var body *url.Values
	body = &url.Values{}
	refund := &xpay.Refund{}
	err := c.B.Call("GET", fmt.Sprintf("/charges/%v/refunds/%v", chid, reid), c.Key, body, nil, refund)
	return refund, err
}

func List(chid string, params *xpay.RefundListParams) *Iter {
	return getC().List(chid, params)
}

func (c Client) List(chid string, params *xpay.RefundListParams) *Iter {
	body := &url.Values{}
	var lp *xpay.ListParams

	params.AppendTo(body)
	lp = &params.ListParams

	return &Iter{xpay.GetIter(lp, body, func(b url.Values) ([]interface{}, xpay.ListMeta, error) {
		list := &xpay.RefundList{}
		err := c.B.Call("GET", fmt.Sprintf("/charges/%v/refunds", chid), c.Key, &b, nil, list)

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

func getC() Client {
	return Client{xpay.GetBackend(xpay.APIBackend), xpay.Key}
}
