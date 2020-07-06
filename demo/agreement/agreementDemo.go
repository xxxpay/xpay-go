package agreement

/*
 * XPay Server SDK
 * 说明：
 * 以下代码只是为了方便商户测试而提供的样例代码，商户可以根据自己网站的需要，按照技术文档编写, 并非一定要使用该代码。
 * 该代码仅供学习和研究 XPay SDK 使用，只是提供一个参考。
 */

import (
	"github.com/xxxpay/xpay-go/demo/common"
	xpay "github.com/xxxpay/xpay-go/xpay"
	agreement "github.com/xxxpay/xpay-go/xpay/agreement"
)

// Demo 示例对象
var Demo = new(AgreementDemo)

// AgreementDemo 签约示例
type AgreementDemo struct {
	demoAppID       string
	demoAgreementID string
}

// Setup 初始化环境
func (c *AgreementDemo) Setup(app string) {
	c.demoAppID = app
}

// New 创建签约对象 agreement
func (c *AgreementDemo) New() (*xpay.Agreement, error) {
	params := &xpay.AgreementParams{
		App:        c.demoAppID,
		ContractNo: "签约协议号",
		Channel:    "qpay",
		Metadata: map[string]interface{}{
			"key": "value",
		},
	}
	return agreement.New(params)
}

// Get 查询签约对象 agreement
func (c *AgreementDemo) Get() (*xpay.Agreement, error) {
	return agreement.Get(c.demoAgreementID)
}

// List 查询签约对象列表
func (c *AgreementDemo) List() (*xpay.AgreementList, error) {
	params := &xpay.PagingParams{}
	params.Filters.AddFilter("per_page", "", "3")
	return agreement.List(c.demoAppID, "*", params)
}

// Update 更新签约对象
func (c *AgreementDemo) Update() (*xpay.Agreement, error) {
	params := &xpay.AgreementUpdateParams{
		Status: "canceled",
	}
	return agreement.Update(c.demoAgreementID, params)
}

// Run 运行
func (c *AgreementDemo) Run() {
	agreement, err := c.New()
	common.Response(agreement, err)
	c.demoAgreementID = agreement.ID
	agreement, err = c.Get()
	common.Response(agreement, err)
	agreementList, err := c.List()
	common.Response(agreementList, err)
	agreement, err = c.Update()
	common.Response(agreement, err)
}
