package settleAccount

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

func getC() Client {
	return Client{xpay.GetBackend(xpay.APIBackend), xpay.Key}
}

func New(appId, userId string, params *xpay.SettleAccountParams) (*xpay.SettleAccount, error) {
	return getC().New(appId, userId, params)
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
	err := c.B.Call("POST", fmt.Sprintf("/apps/%s/users/%s/settle_accounts", appId, userId), c.Key, nil, paramsString, settle_account)
	return settle_account, err
}

func Get(appId, userId, settleAccountId string) (*xpay.SettleAccount, error) {
	return getC().Get(appId, userId, settleAccountId)
}

func (c Client) Get(appId, userId, settleAccountId string) (*xpay.SettleAccount, error) {
	settleAccount := &xpay.SettleAccount{}

	err := c.B.Call("GET", fmt.Sprintf("/apps/%s/users/%s/settle_accounts/%s", appId, userId, settleAccountId), c.Key, nil, nil, settleAccount)
	return settleAccount, err
}

func List(appId, userId string, params *xpay.PagingParams) (*xpay.SettleAccountList, error) {
	return getC().List(appId, userId, params)
}
func (c Client) List(appId, userId string, params *xpay.PagingParams) (*xpay.SettleAccountList, error) {
	body := &url.Values{}
	params.Filters.AppendTo(body)

	settleAccountList := &xpay.SettleAccountList{}
	err := c.B.Call("GET", fmt.Sprintf("/apps/%s/users/%s/settle_accounts", appId, userId), c.Key, body, nil, settleAccountList)
	return settleAccountList, err
}

func Delete(appId, userId, settleAccountId string) (*xpay.DeleteResult, error) {
	return getC().Delete(appId, userId, settleAccountId)
}

func (c Client) Delete(appId, userId, settleAccountId string) (*xpay.DeleteResult, error) {
	result := &xpay.DeleteResult{}

	err := c.B.Call("DELETE", fmt.Sprintf("/apps/%s/users/%s/settle_accounts/%s", appId, userId, settleAccountId), c.Key, nil, nil, result)
	if err != nil {
		if xpay.LogLevel > 0 {
			log.Printf("%v\n", err)
		}
		return nil, err
	}
	return result, err
}
