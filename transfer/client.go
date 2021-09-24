package transfer

import (
	"log"
	"net/url"
	"strconv"

	"github.com/xxxpay/xpay-go"
)

type Client struct {
	backend xpay.Backend
}

func NewClient(backend xpay.Backend) Client {
	return Client{backend: backend}
}

func (c Client) New(params *xpay.TransferParams) (*xpay.Transfer, error) {
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
	transfer := &xpay.Transfer{}
	err := c.backend.Call("POST", "/transfers", nil, paramsString, transfer)
	return transfer, err
}

func (c Client) Update(id string) (*xpay.Transfer, error) {
	cancelParams := struct {
		Status string `json:"status"`
	}{
		Status: "canceled",
	}

	paramsString, _ := xpay.JsonEncode(cancelParams)
	transfer := &xpay.Transfer{}
	err := c.backend.Call("PUT", "/transfers/"+id, nil, paramsString, transfer)
	return transfer, err
}

// Get returns the details of a redenvelope.
func (c Client) Get(id string) (*xpay.Transfer, error) {
	var body *url.Values
	body = &url.Values{}
	transfer := &xpay.Transfer{}
	err := c.backend.Call("GET", "/transfers/"+id, body, nil, transfer)
	return transfer, err
}

// List returns a list of transfer.
func (c Client) List(params *xpay.TransferListParams) *Iter {
	type transferList struct {
		xpay.ListMeta
		Values []*xpay.Transfer `json:"data"`
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
		list := &transferList{}
		err := c.backend.Call("GET", "/transfers", &b, nil, list)

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

func (i *Iter) Transfer() *xpay.Transfer {
	return i.Current().(*xpay.Transfer)
}
