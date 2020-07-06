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
	"github.com/xxxpay/xpay-go/xpay/royaltyTemplate"
)

var TmplDemo = new(RoyaltyTmplDemo)

type RoyaltyTmplDemo struct {
	demoAppID         string
	demoroyaltyTmplId string
}

func (c *RoyaltyTmplDemo) Setup(app string) {
	c.demoAppID = app
	c.demoroyaltyTmplId = "451170807182300001"
}

//创建分润模版
func (c *RoyaltyTmplDemo) New() (*xpay.RoyaltyTmpl, error) {
	params := &xpay.RoyaltyTmplParams{
		App:  c.demoAppID,
		Name: "royalty_template_name",
		Rule: xpay.Rule{
			Royalty_mode:    "rate",
			Refund_mode:     "no_refund",
			Allocation_mode: "receipt_reserved",
			Data: []xpay.RuleData{
				xpay.RuleData{
					Level: 1, Value: 30,
				},
				xpay.RuleData{
					Level: 2, Value: 20,
				},
				xpay.RuleData{
					Level: 3, Value: 10,
				},
			},
		},
		Description: "Your description",
	}
	return royaltyTemplate.New(params)
}

// 查询分润模版
func (c *RoyaltyTmplDemo) Get() (*xpay.RoyaltyTmpl, error) {
	return royaltyTemplate.Get(c.demoroyaltyTmplId)
}

// 查询分润模版列表
func (c *RoyaltyTmplDemo) List() (*xpay.RoyaltyTmplList, error) {
	params := &xpay.PagingParams{}
	params.Filters.AddFilter("page", "", "1")
	params.Filters.AddFilter("per_page", "", "10")
	return royaltyTemplate.List(params)
}

// 更新分润模版
func (c *RoyaltyTmplDemo) Update() (*xpay.RoyaltyTmpl, error) {
	params := &xpay.RoyaltyTmplUpdateParams{
		Name: "royalty_template_name",
		Rule: xpay.Rule{
			Royalty_mode:    "fixed",
			Refund_mode:     "full_refund",
			Allocation_mode: "service_reserved",
			Data: []xpay.RuleData{
				xpay.RuleData{
					Level: 1, Value: 33,
				},
				xpay.RuleData{
					Level: 2, Value: 22,
				},
				xpay.RuleData{
					Level: 3, Value: 11,
				},
			},
		},
		Description: "Your description",
	}
	return royaltyTemplate.Update(c.demoroyaltyTmplId, params)
}

// 删除分润模版
func (c *RoyaltyTmplDemo) Delete() (*xpay.DeleteResult, error) {
	return royaltyTemplate.Delete(c.demoroyaltyTmplId)
}

func (c *RoyaltyTmplDemo) Run() {
	tpl, err := c.New()
	common.Response(tpl, err)
	c.demoroyaltyTmplId = tpl.ID
	common.Response(c.Get())
	common.Response(c.List())
	common.Response(c.Update())
}
