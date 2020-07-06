/* *
 * XPay Server SDK
 * 说明：
 * 以下代码只是为了方便商户测试而提供的样例代码，商户可以根据自己网站的需要，按照技术文档编写, 并非一定要使用该代码。
 * 该代码仅供学习和研究 XPay SDK 使用，只是提供一个参考。
 */
package royalty

import (
	"github.com/xxxpay/xpay-go/demo/common"
	xpay "github.com/xxxpay/xpay-go/xpay"
	"github.com/xxxpay/xpay-go/xpay/royalty"
)

var Demo = new(RoyaltyDemo)

type RoyaltyDemo struct {
	demoAppID string
	royaltyId string
}

func (c *RoyaltyDemo) Setup(app string) {
	c.demoAppID = app
	c.royaltyId = "411170318160900002"
}

//查询分润对象
func (c *RoyaltyDemo) Get() (*xpay.Royalty, error) {
	return royalty.Get(c.royaltyId)
}

//查询分润对象列表
func (c *RoyaltyDemo) List() (*xpay.RoyaltyList, error) {
	params := &xpay.PagingParams{}
	params.Filters.AddFilter("per_page", "", "3")
	return royalty.List(params)
}

// 批量更新分润对象
func (c *RoyaltyDemo) BatchUpdate() (*xpay.RoyaltyList, error) {
	params := &xpay.RoyaltyBatchUpdateParams{
		Ids:    []string{"411170321144900002", "411170321174700002", "421170321175200002"},
		Method: "manual",
	}
	return royalty.BatchUpdate(params)
}

func (c *RoyaltyDemo) Run() {
	common.Response(c.Get())
	common.Response(c.List())
	common.Response(c.BatchUpdate())
}
