package payment

import (
	"log"
	"net/url"
	"strconv"
	"time"

	"github.com/xxxpay/xpay-go"
)

type Client struct {
	backend xpay.Backend
}

func NewClient(backend xpay.Backend) Client {
	return Client{backend: backend}
}

func (c Client) New(params *xpay.PaymentParams) (*xpay.Payment, error) {
	start := time.Now()
	paramsString, errs := xpay.JsonEncode(params)
	if errs != nil {
		if xpay.LogLevel > 0 {
			log.Printf("PaymentParams Marshall Errors is : %q\n", errs)
		}
	}
	if xpay.LogLevel > 2 {
		log.Printf("params of payment request to xpay is :\n %v\n ", string(paramsString))
	}

	payment := &xpay.Payment{}
	errch := c.backend.Call("POST", "/payments", nil, paramsString, payment)
	if errch != nil {
		if xpay.LogLevel > 0 {
			log.Printf("%v\n", errch)
		}
		return nil, errch
	}
	if xpay.LogLevel > 2 {
		log.Println("Payment completed in ", time.Since(start))
	}
	return payment, errch

}

func (c Client) Reverse(id string) (*xpay.Payment, error) {
	var body *url.Values
	body = &url.Values{}
	payment := &xpay.Payment{}
	err := c.backend.Call("POST", "/payments/"+id+"/reverse", body, nil, payment)
	if err != nil {
		if xpay.LogLevel > 0 {
			log.Printf("Reverse Payment error: %v\n", err)
		}
	}
	return payment, err
}

func (c Client) Get(id string) (*xpay.Payment, error) {
	var body *url.Values
	body = &url.Values{}
	payment := &xpay.Payment{}
	err := c.backend.Call("GET", "/payments/"+id, body, nil, payment)
	if err != nil {
		if xpay.LogLevel > 0 {
			log.Printf("Get Payment error: %v\n", err)
		}
	}
	return payment, err
}

func (c Client) List(appId string, params *xpay.PaymentListParams) *Iter {
	type chargeList struct {
		xpay.ListMeta
		Values []*xpay.Payment `json:"data"`
	}

	var body *url.Values
	var lp *xpay.ListParams

	if params == nil {
		params = &xpay.PaymentListParams{}
	}
	params.Filters.AddFilter("app[id]", "", appId)
	body = &url.Values{}
	if params.Created > 0 {
		body.Add("created", strconv.FormatInt(params.Created, 10))
	}
	params.AppendTo(body)
	lp = &params.ListParams

	return &Iter{xpay.GetIter(lp, body, func(b url.Values) ([]interface{}, xpay.ListMeta, error) {
		list := &chargeList{}
		err := c.backend.Call("GET", "/payments", &b, nil, list)

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

func (i *Iter) Payment() *xpay.Payment {
	return i.Current().(*xpay.Payment)
}
