package redEnvelope

import (
	"github.com/xxxpay/xpay-go"
	"log"
	"net/url"
	"strconv"
)

type Client struct {
	backend xpay.Backend
}

func NewClient(backend xpay.Backend) Client {
	return Client{backend: backend}
}

func (c Client) New(params *xpay.RedEnvelopeParams) (*xpay.RedEnvelope, error) {
	paramsString, errs := xpay.JsonEncode(params)
	if errs != nil {
		if xpay.LogLevel > 0 {
			log.Printf("PaymentParams Marshall Errors is : %q/n", errs)
		}
		return nil, errs
	}
	if xpay.LogLevel > 2 {
		log.Printf("params of redEnvelope request to xpay is :\n %v\n ", string(paramsString))
	}
	redEnvelope := &xpay.RedEnvelope{}
	err := c.backend.Call("POST", "/red_envelopes", nil, paramsString, redEnvelope)
	return redEnvelope, err
}

func (c Client) Get(id string) (*xpay.RedEnvelope, error) {
	var body *url.Values
	body = &url.Values{}
	redEnvelope := &xpay.RedEnvelope{}
	err := c.backend.Call("GET", "/red_envelopes/"+id, body, nil, redEnvelope)
	return redEnvelope, err
}

func (c Client) List(params *xpay.RedEnvelopeListParams) *Iter {
	type redEnvelopeList struct {
		xpay.ListMeta
		Values []*xpay.RedEnvelope `json:"data"`
	}

	var body *url.Values
	var lp *xpay.ListParams

	if params != nil {
		body = &url.Values{}

		if params.Created > 0 {
			body.Add("created", strconv.FormatInt(params.Created, 10))
		}
		params.AppendTo(body)
		lp = &params.ListParams
	}

	return &Iter{xpay.GetIter(lp, body, func(b url.Values) ([]interface{}, xpay.ListMeta, error) {
		list := &redEnvelopeList{}
		err := c.backend.Call("GET", "/red_envelopes", &b, nil, list)

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

func (i *Iter) RedEnvelope() *xpay.RedEnvelope {
	return i.Current().(*xpay.RedEnvelope)
}
