// Package profitTransaction 分账明细示例
/* *
 * XPay Server SDK
 * 说明：
 * 以下代码只是为了方便商户测试而提供的样例代码，商户可以根据自己网站的需要，按照技术文档编写, 并非一定要使用该代码。
 * 该代码仅供学习和研究 XPay SDK 使用，只是提供一个参考。
 */
package profitTransaction

import (
	"github.com/xxxpay/xpay-go/xpay"
	"github.com/xxxpay/xpay-go/xpay/profitTransaction"
)

// Demo 分账明细示例
var Demo = new(demo)

// demo 分账明细示例
type demo struct {
	app string
}

// Setup 设置参数
func (c *demo) Setup(app string) {
	c.app = app
}

// Get 查询 分账明细 对象
func (c *demo) Get() (*xpay.ProfitTransaction, error) {
	return profitTransaction.Get("ptxn_1m3x7aGbDK2cpl")
}

// List 查询 分账明细 对象列表
func (c *demo) List() (xpay.ProfitTransactionList, error) {
	params := &xpay.PagingParams{}
	params.Filters.AddFilter("page", "", "1")
	params.Filters.AddFilter("per_page", "", "100")
	return profitTransaction.List(c.app, "", "", "", params)
}

// Run 运行示例
func (c *demo) Run() {
	c.Get()
	c.List()
}
