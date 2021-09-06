package balanceBonus

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

func New(appId string, params *xpay.BalanceBonusParams) (*xpay.BalanceBonus, error) {
	return getC().New(appId, params)
}

func (c Client) New(appId string, params *xpay.BalanceBonusParams) (*xpay.BalanceBonus, error) {
	paramsString, _ := xpay.JsonEncode(params)
	if xpay.LogLevel > 2 {
		log.Printf("params of balance bonus to xpay is :\n %v\n ", string(paramsString))
	}
	balanceBonus := &xpay.BalanceBonus{}

	err := c.B.Call("POST", fmt.Sprintf("/apps/%s/balance_bonuses", appId), c.Key, nil, paramsString, balanceBonus)
	if err != nil {
		if xpay.LogLevel > 0 {
			log.Printf("Balance Bonus error: %v\n", err)
		}
	}
	return balanceBonus, err
}

func Get(appId, balanceBonusId string) (*xpay.BalanceBonus, error) {
	return getC().Get(appId, balanceBonusId)
}

func (c Client) Get(appId, balanceBonusID string) (*xpay.BalanceBonus, error) {
	balanceBonus := &xpay.BalanceBonus{}

	err := c.B.Call("GET", fmt.Sprintf("/apps/%s/balance_bonuses/%s", appId, balanceBonusID), c.Key, nil, nil, balanceBonus)
	return balanceBonus, err
}

func List(appID string, params *xpay.PagingParams) (*xpay.BalanceBonusList, error) {
	return getC().List(appID, params)
}

func (c Client) List(appID string, params *xpay.PagingParams) (*xpay.BalanceBonusList, error) {
	body := &url.Values{}
	params.Filters.AppendTo(body)
	balanceBonusList := &xpay.BalanceBonusList{}

	err := c.B.Call("GET", fmt.Sprintf("/apps/%s/balance_bonuses", appID), c.Key, body, nil, balanceBonusList)
	return balanceBonusList, err
}
