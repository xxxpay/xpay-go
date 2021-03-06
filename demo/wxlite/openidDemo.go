package wxlite

/*
 * XPay Server SDK
 * 说明：
 * 以下代码只是为了方便商户测试而提供的样例代码，商户可以根据自己网站的需要，按照技术文档编写, 并非一定要使用该代码。
 * 该代码仅供学习和研究 XPay SDK 使用，只是提供一个参考。
 */

import (
	"github.com/xxxpay/xpay-go/demo/common"

	xpay "github.com/xxxpay/xpay-go/xpay"
	wxlite "github.com/xxxpay/xpay-go/xpay/wxlite"
)

// Demo 示例对象
var Demo = new(wxliteDemo)

// wxliteDemo 签约示例
type wxliteDemo struct{}

func (c *wxliteDemo) Setup(app string) {
}

// New 创建签约对象 agreement
func (c *wxliteDemo) GetOpenid() (*xpay.Openid, error) {
	params := &xpay.OpenidParams{
		AppID:     "app id",
		AppSecret: "app secret",
		Code:      "code",
	}
	return wxlite.GetOpenid(params)
}

// Run 运行
func (c *wxliteDemo) Run() {
	openid, err := c.GetOpenid()
	common.Response(openid, err)
}
