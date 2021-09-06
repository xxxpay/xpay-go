package agreement

import (
	"fmt"
	"log"
	"net/url"

	"github.com/xxxpay/xpay-go"
)

// Client 请求
type Client struct {
	B   xpay.Backend
	Key string
}

func getC() Client {
	return Client{xpay.GetBackend(xpay.APIBackend), xpay.Key}
}

// New 创建签约
// @param appId string
// @param params AgreementParams
// @return Agreement
func New(params *xpay.AgreementParams) (*xpay.Agreement, error) {
	return getC().New(params)
}

// New 创建签约
func (c Client) New(params *xpay.AgreementParams) (*xpay.Agreement, error) {
	paramsString, errs := xpay.JsonEncode(params)
	if errs != nil {
		if xpay.LogLevel > 0 {
			log.Printf("AgreementParams Marshall Errors is : %q/n", errs)
		}
		return nil, errs
	}
	if xpay.LogLevel > 2 {
		log.Printf("params of create agreement is :\n %v\n ", string(paramsString))
	}

	agreement := &xpay.Agreement{}
	err := c.B.Call("POST", "/agreements", c.Key, nil, paramsString, agreement)
	return agreement, err
}

// Get 查询签约对象
// @param agreementID 签约对象 ID
// @return Agreement
func Get(agreementID string) (*xpay.Agreement, error) {
	return getC().Get(agreementID)
}

// Get 查询签约对象
func (c Client) Get(agreementID string) (*xpay.Agreement, error) {
	agreement := &xpay.Agreement{}
	err := c.B.Call("GET", fmt.Sprintf("/agreements/%s", agreementID), c.Key, nil, nil, agreement)
	return agreement, err
}

// List 查询签约对象列表
// @param app string
// @param status string
// @param params PagingParams
// @return AgreementList
func List(app, status string, params *xpay.PagingParams) (*xpay.AgreementList, error) {
	return getC().List(app, status, params)
}

// List 查询签约对象列表
func (c Client) List(app, status string, params *xpay.PagingParams) (*xpay.AgreementList, error) {
	body := &url.Values{}
	params.Filters.AppendTo(body)
	body.Add("app", app)
	if status != "" && status != "*" {
		body.Add("status", status)
	}
	agreements := &xpay.AgreementList{}
	err := c.B.Call("GET", "/agreements", c.Key, body, nil, &agreements)
	return agreements, err
}

// Update 更新签约商户对象
// @param agreementID string
// @param params AgreementUpdateParams
// @return Agreement
func Update(agreementID string, params *xpay.AgreementUpdateParams) (*xpay.Agreement, error) {
	return getC().Update(agreementID, params)
}

// Update 更新子商户对象
func (c Client) Update(agreementID string, params *xpay.AgreementUpdateParams) (*xpay.Agreement, error) {
	paramsString, _ := xpay.JsonEncode(params)
	if xpay.LogLevel > 2 {
		log.Printf("params of update Agreement  to xpay is :\n %v\n ", string(paramsString))
	}

	agreement := &xpay.Agreement{}
	err := c.B.Call("PUT", fmt.Sprintf("/agreements/%s", agreementID), c.Key, nil, paramsString, agreement)
	if err != nil {
		if xpay.LogLevel > 0 {
			log.Printf("%v\n", err)
		}
		return nil, err
	}
	return agreement, err
}
