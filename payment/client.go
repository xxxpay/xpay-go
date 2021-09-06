package payment

import (
	"log"
	"net/url"
	"strconv"
	"time"

	"github.com/xxxpay/xpay-go"
)

type Client struct {
	B   xpay.Backend
	Key string
}

func getC() Client {
	return Client{xpay.GetBackend(xpay.APIBackend), xpay.Key}
}

// 发送 payment 请求
func New(params *xpay.PaymentParams) (*xpay.Payment, error) {
	return getC().New(params)
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
	errch := c.B.Call("POST", "/payments", c.Key, nil, paramsString, payment)
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

// 撤销charge，此接口仅接受线下 isv_scan、isv_wap、isv_qr 渠道的订单调用
func Reverse(id string) (*xpay.Payment, error) {
	return getC().Reverse(id)
}

func (c Client) Reverse(id string) (*xpay.Payment, error) {
	var body *url.Values
	body = &url.Values{}
	payment := &xpay.Payment{}
	err := c.B.Call("POST", "/payments/"+id+"/reverse", c.Key, body, nil, payment)
	if err != nil {
		if xpay.LogLevel > 0 {
			log.Printf("Reverse Payment error: %v\n", err)
		}
	}
	return payment, err
}

//查询指定 payment 对象
func Get(id string) (*xpay.Payment, error) {
	return getC().Get(id)
}

func (c Client) Get(id string) (*xpay.Payment, error) {
	var body *url.Values
	body = &url.Values{}
	payment := &xpay.Payment{}
	err := c.B.Call("GET", "/payments/"+id, c.Key, body, nil, payment)
	if err != nil {
		if xpay.LogLevel > 0 {
			log.Printf("Get Payment error: %v\n", err)
		}
	}
	return payment, err
}

// 查询 payment 列表
func List(appId string, params *xpay.PaymentListParams) *Iter {
	return getC().List(appId, params)
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
		err := c.B.Call("GET", "/payments", c.Key, &b, nil, list)

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
