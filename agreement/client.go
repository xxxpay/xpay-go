package agreement

import (
	"fmt"
	"log"
	"net/url"

	"github.com/xxxpay/xpay-go"
)

// Client 请求
type Client struct {
	backend xpay.Backend
}

func NewClient(backend xpay.Backend) Client {
	return Client{backend: backend}
}

// New 创建签约
// @param params AgreementParams
// @return Agreement
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
	err := c.backend.Call("POST", "/agreements", nil, paramsString, agreement)
	return agreement, err
}

// Get 查询签约对象
// @param agreementID 签约对象 ID
// @return Agreement
func (c Client) Get(agreementID string) (*xpay.Agreement, error) {
	agreement := &xpay.Agreement{}
	err := c.backend.Call("GET", fmt.Sprintf("/agreements/%s", agreementID), nil, nil, agreement)
	return agreement, err
}

// List 查询签约对象列表
// @param app string
// @param status string
// @param params PagingParams
// @return AgreementList
func (c Client) List(app, status string, params *xpay.PagingParams) (*xpay.AgreementList, error) {
	body := &url.Values{}
	params.Filters.AppendTo(body)
	body.Add("app", app)
	if status != "" && status != "*" {
		body.Add("status", status)
	}
	agreements := &xpay.AgreementList{}
	err := c.backend.Call("GET", "/agreements", body, nil, &agreements)
	return agreements, err
}

// Update 更新签约商户对象
// @param agreementID string
// @param params AgreementUpdateParams
// @return Agreement
func (c Client) Update(agreementID string, params *xpay.AgreementUpdateParams) (*xpay.Agreement, error) {
	paramsString, _ := xpay.JsonEncode(params)
	if xpay.LogLevel > 2 {
		log.Printf("params of update Agreement  to xpay is :\n %v\n ", string(paramsString))
	}

	agreement := &xpay.Agreement{}
	err := c.backend.Call("PUT", fmt.Sprintf("/agreements/%s", agreementID), nil, paramsString, agreement)
	if err != nil {
		if xpay.LogLevel > 0 {
			log.Printf("%v\n", err)
		}
		return nil, err
	}
	return agreement, err
}
