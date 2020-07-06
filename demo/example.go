package main

import (
	"io/ioutil"
	"os"

	payment "github.com/xxxpay/xpay-go/demo/payment"
	xpay "github.com/xxxpay/xpay-go/xpay"
)

func readFile(path string) string {
	fi, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	if err != nil {
		panic(err)
	}
	return string(fd)
}

type Demo struct {
	examples []Example
	App      string
}

type Example interface {
	Setup(app string)
	Run()
}

func (c *Demo) Setup() {
	c.App = "f5fe7e65ad7f46af89921565f78a9879"
	xpay.Key = "3d6285c77d1646f096c21210f9a6db22"
	xpay.LogLevel = 2
	xpay.AccountPrivateKey = readFile("your_rsa_private_key.pem")

	for i := len(c.examples) - 1; i >= 0; i-- {
		c.examples[i].Setup(c.App)
	}
}

func (c *Demo) Run() {
	for i := len(c.examples) - 1; i >= 0; i-- {
		c.examples[i].Run()
	}
}

func (c *Demo) Use(e Example) {
	c.examples = append(c.examples, e)
}

func main() {
	demo := new(Demo)
	// demo.Use(sub_app.Demo)
	// demo.Use(balance.BonusDemo)
	// demo.Use(balance.TransferDemo)
	// demo.Use(balance.TransactionDemo)
	// demo.Use(batch_refund.Demo)
	// demo.Use(batch_transfer.Demo)
	// demo.Use(batch_withdraw.Demo)
	// demo.Use(card.Demo)
	// demo.Use(channel.Demo)
	demo.Use(payment.Demo)
	// demo.Use(coupon.Demo)
	// demo.Use(coupon.TmplDemo)
	// demo.Use(customer.Demo)
	// demo.Use(customs.Demo)
	// demo.Use(event.Demo)
	// demo.Use(identification.Demo)
	// demo.Use(order.Demo)
	// demo.Use(order_refund.Demo)
	// demo.Use(recharge.Demo)
	// demo.Use(red_envelope.Demo)
	// demo.Use(refund.Demo)
	// demo.Use(royalty.Demo)
	// demo.Use(royalty.SettlementDemo)
	// demo.Use(royalty.TmplDemo)
	// demo.Use(royalty.TransactionDemo)
	// demo.Use(settle_account.Demo)
	// demo.Use(token.Demo)
	// demo.Use(transfer.Demo)
	// demo.Use(user.Demo)
	// demo.Use(verify.Demo)
	// demo.Use(withdrawal.Demo)
	// demo.Use(agreement.Demo)
	// demo.Use(wxlite.Demo)
	// demo.Use(splitProfit.Demo)
	// demo.Use(profitTransaction.Demo)
	// demo.Use(subBank.Demo)
	// demo.Use(splitReceiver.Demo)

	// demo.Use(webhookDemo) //http server只能单独测试，别Use其他demo
	// demo.Use(payServerDemo) //http server只能单独测试，别Use其他demo
	demo.Setup()
	demo.Run()
}
