package settleAccount

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

func (c Client) New(appId, userId string, params *xpay.SettleAccountParams) (*xpay.SettleAccount, error) {
	paramsString, errs := xpay.JsonEncode(params)
	if errs != nil {
		if xpay.LogLevel > 0 {
			log.Printf("SettleAccountParams Marshall Errors is : %q/n", errs)
		}
		return nil, errs
	}
	if xpay.LogLevel > 2 {
		log.Printf("params of create SettleAccount is :\n %v\n ", string(paramsString))
	}

	settle_account := &xpay.SettleAccount{}
	err := c.backend.Call("POST", fmt.Sprintf("/apps/%s/users/%s/settle_accounts", appId, userId), nil, paramsString, settle_account)
	return settle_account, err
}

func (c Client) Get(appId, userId, settleAccountId string) (*xpay.SettleAccount, error) {
	settleAccount := &xpay.SettleAccount{}

	err := c.backend.Call("GET", fmt.Sprintf("/apps/%s/users/%s/settle_accounts/%s", appId, userId, settleAccountId), nil, nil, settleAccount)
	return settleAccount, err
}
func (c Client) List(appId, userId string, params *xpay.PagingParams) (*xpay.SettleAccountList, error) {
	body := &url.Values{}
	params.Filters.AppendTo(body)

	settleAccountList := &xpay.SettleAccountList{}
	err := c.backend.Call("GET", fmt.Sprintf("/apps/%s/users/%s/settle_accounts", appId, userId), body, nil, settleAccountList)
	return settleAccountList, err
}

func (c Client) Delete(appId, userId, settleAccountId string) (*xpay.DeleteResult, error) {
	result := &xpay.DeleteResult{}

	err := c.backend.Call("DELETE", fmt.Sprintf("/apps/%s/users/%s/settle_accounts/%s", appId, userId, settleAccountId), nil, nil, result)
	if err != nil {
		if xpay.LogLevel > 0 {
			log.Printf("%v\n", err)
		}
		return nil, err
	}
	return result, err
}
