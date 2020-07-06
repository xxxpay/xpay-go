/* *
 * XPay Server SDK
 * 说明：
 * 以下代码只是为了方便商户测试而提供的样例代码，商户可以根据自己网站的需要，按照技术文档编写, 并非一定要使用该代码。
 * 该代码仅供学习和研究 XPay SDK 使用，只是提供一个参考。
 */
package order_refund

import (
	"github.com/xxxpay/xpay-go/demo/common"
	xpay "github.com/xxxpay/xpay-go/xpay"
	"github.com/xxxpay/xpay-go/xpay/orderRefund"
)

var Demo = new(OrderRefundDemo)

type OrderRefundDemo struct {
	demoAppID string
}

func (c *OrderRefundDemo) Setup(app string) {
	c.demoAppID = app
}

func (c *OrderRefundDemo) New() (*xpay.RefundList, error) {
	params := &xpay.OrderRefundParams{
		Description: "Go SDK Test",
	}

	return orderRefund.New("2011609290000001291", params)
}

func (c *OrderRefundDemo) Get() (*xpay.Refund, error) {
	return orderRefund.Get("2011609290000001291", "2111609290000001601")
}

func (c *OrderRefundDemo) List() (*xpay.RefundList, error) {
	params := &xpay.PagingParams{}
	params.Filters.AddFilter("page", "", "1")     //取第一页数据
	params.Filters.AddFilter("per_page", "", "2") //每页两个Order对象
	return orderRefund.List("2011609290000001291", params)
}

func (c *OrderRefundDemo) Run() {
	common.Response(c.New())
	common.Response(c.Get())
	common.Response(c.List())
}
