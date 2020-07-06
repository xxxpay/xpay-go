package card

import (
	"github.com/xxxpay/xpay-go/demo/common"
	"github.com/xxxpay/xpay-go/xpay"
	"github.com/xxxpay/xpay-go/xpay/cardInfo"
)

var Demo = new(CardInfoDemo)

type CardInfoDemo struct {
	demoAppID string
}

func (c *CardInfoDemo) Setup(app string) {
	c.demoAppID = app
}

func (c *CardInfoDemo) New() (*xpay.CardInfo, error) {
	param := &xpay.CardInfoParams{App: c.demoAppID, BankAccount: "6228480402564890018"}
	return cardInfo.New(param)
}

func (c *CardInfoDemo) Run() {
	card, err := c.New()
	common.Response(card, err)
}
