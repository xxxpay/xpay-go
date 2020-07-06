/* *
 * XPay Server SDK
 * 说明：
 * 以下代码只是为了方便商户测试而提供的样例代码，商户可以根据自己网站的需要，按照技术文档编写, 并非一定要使用该代码。
 * 该代码仅供学习和研究 XPay SDK 使用，只是提供一个参考。
 */
package batch_refund

import (
	"time"

	"github.com/xxxpay/xpay-go/demo/common"
	"github.com/xxxpay/xpay-go/xpay"
	"github.com/xxxpay/xpay-go/xpay/batchRefund"
)

var Demo = new(BatchRefundDemo)

type BatchRefundDemo struct {
	demoAppID         string
	payments          []string
	demoBatchRefundID string
}

func (c *BatchRefundDemo) Setup(app string) {
	c.demoAppID = app
	c.payments = []string{"ch_L8qn10mLmr1GS8e5OODmHaL4", "ch_fdOmHaLmLmr1GOD4qn1dS8e5"} // 需要先支付两笔 payment，才能做批量退款
}

//创建批量退款
func (c *BatchRefundDemo) New() (*xpay.BatchRefund, error) {
	params := &xpay.BatchRefundParams{
		App:         c.demoAppID,
		Batch_no:    "batchrefund" + time.Now().Format("060102150405"),
		Description: "Your Description",
	}

	for _, payment := range c.payments {
		params.Payments = append(params.Payments, map[string]interface{}{
			"payment":     payment,
			"description": "Batch refund description.",
		})
	}
	return batchRefund.New(params)
}

//查询批量退款
func (c *BatchRefundDemo) Get() (*xpay.BatchRefund, error) {
	return batchRefund.Get(c.demoBatchRefundID)
}

//查询 Batch Refund 对象列表
func (c *BatchRefundDemo) List() (*xpay.BatchRefundlList, error) {
	params := &xpay.PagingParams{}
	params.Filters.AddFilter("page", "", "1")
	params.Filters.AddFilter("per_page", "", "2")
	params.Filters.AddFilter("app", "", c.demoAppID)
	return batchRefund.List(params)
}

func (c *BatchRefundDemo) Run() {
	batch_refund, err := c.New()
	common.Response(batch_refund, err)
	c.demoBatchRefundID = batch_refund.Id
	batch_refund, err = c.Get()
	common.Response(batch_refund, err)
	batch_refund_list, err := c.List()
	common.Response(batch_refund_list, err)
}
