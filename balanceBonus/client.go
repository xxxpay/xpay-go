package balanceBonus

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

func (c Client) New(appId string, params *xpay.BalanceBonusParams) (*xpay.BalanceBonus, error) {
	paramsString, _ := xpay.JsonEncode(params)
	if xpay.LogLevel > 2 {
		log.Printf("params of balance bonus to xpay is :\n %v\n ", string(paramsString))
	}
	balanceBonus := &xpay.BalanceBonus{}

	err := c.backend.Call("POST", fmt.Sprintf("/apps/%s/balance_bonuses", appId), nil, paramsString, balanceBonus)
	if err != nil {
		if xpay.LogLevel > 0 {
			log.Printf("Balance Bonus error: %v\n", err)
		}
	}
	return balanceBonus, err
}

func (c Client) Get(appId, balanceBonusID string) (*xpay.BalanceBonus, error) {
	balanceBonus := &xpay.BalanceBonus{}

	err := c.backend.Call("GET", fmt.Sprintf("/apps/%s/balance_bonuses/%s", appId, balanceBonusID), nil, nil, balanceBonus)
	return balanceBonus, err
}

func (c Client) List(appID string, params *xpay.PagingParams) (*xpay.BalanceBonusList, error) {
	body := &url.Values{}
	params.Filters.AppendTo(body)
	balanceBonusList := &xpay.BalanceBonusList{}

	err := c.backend.Call("GET", fmt.Sprintf("/apps/%s/balance_bonuses", appID), body, nil, balanceBonusList)
	return balanceBonusList, err
}
