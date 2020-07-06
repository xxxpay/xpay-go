/* *
 * XPay Server SDK
 * 说明：
 * 以下代码只是为了方便商户测试而提供的样例代码，商户可以根据自己网站的需要，按照技术文档编写, 并非一定要使用该代码。
 * 该代码仅供学习和研究 XPay SDK 使用，只是提供一个参考。
 */
package balance

import (
	"github.com/xxxpay/xpay-go/demo/common"
	"github.com/xxxpay/xpay-go/xpay"
	"github.com/xxxpay/xpay-go/xpay/balanceTransaction"
)

var TransactionDemo = new(BalanceTransactionDemo)

type BalanceTransactionDemo struct {
	demoAppID            string
	balanceTransactionId string
}

func (c *BalanceTransactionDemo) Setup(app string) {
	c.demoAppID = app
}

// 用户账户交易明细
func (c *BalanceTransactionDemo) Get() (*xpay.BalanceTransaction, error) {
	return balanceTransaction.Get(c.demoAppID, c.balanceTransactionId)
}

// 查询用户交易列表
func (c *BalanceTransactionDemo) List() (*xpay.BalanceTransactionList, error) {
	params := &xpay.PagingParams{}
	params.Filters.AddFilter("per_page", "", "2")
	return balanceTransaction.List(c.demoAppID, params)
}

func (c *BalanceTransactionDemo) Run() {
	list, err := c.List()
	common.Response(list, err)
	if len(list.Values) >= 1 {
		c.balanceTransactionId = list.Values[0].Id
		bt, err := c.Get()
		common.Response(bt, err)
	}
}
