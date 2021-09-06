package transfer

import (
	"log"
	"net/url"
	"strconv"

	"github.com/xxxpay/xpay-go"
)

type Client struct {
	B   xpay.Backend
	Key string
}

func New(params *xpay.TransferParams) (*xpay.Transfer, error) {
	return getC().New(params)
}

func (c Client) New(params *xpay.TransferParams) (*xpay.Transfer, error) {
	paramsString, errs := xpay.JsonEncode(params)
	if errs != nil {
		if xpay.LogLevel > 0 {
			log.Printf("ChargeParams Marshall Errors is : %q/n", errs)
		}
		return nil, errs
	}
	if xpay.LogLevel > 2 {
		log.Printf("params of redEnvelope request to xpay is :\n %v\n ", string(paramsString))
	}
	transfer := &xpay.Transfer{}
	err := c.B.Call("POST", "/transfers", c.Key, nil, paramsString, transfer)
	return transfer, err
}

func Update(id string) (*xpay.Transfer, error) {
	return getC().Update(id)
}

func (c Client) Update(id string) (*xpay.Transfer, error) {
	cancelParams := struct {
		Status string `json:"status"`
	}{
		Status: "canceled",
	}

	paramsString, _ := xpay.JsonEncode(cancelParams)
	transfer := &xpay.Transfer{}
	err := c.B.Call("PUT", "/transfers/"+id, c.Key, nil, paramsString, transfer)
	return transfer, err
}

// Get returns the details of a redenvelope.
func Get(id string) (*xpay.Transfer, error) {
	return getC().Get(id)
}

func (c Client) Get(id string) (*xpay.Transfer, error) {
	var body *url.Values
	body = &url.Values{}
	transfer := &xpay.Transfer{}
	err := c.B.Call("GET", "/transfers/"+id, c.Key, body, nil, transfer)
	return transfer, err
}

// List returns a list of transfer.
func List(params *xpay.TransferListParams) *Iter {
	return getC().List(params)
}

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
		err := c.B.Call("GET", "/transfers", c.Key, &b, nil, list)

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

func getC() Client {
	return Client{xpay.GetBackend(xpay.APIBackend), xpay.Key}
}
