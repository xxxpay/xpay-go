/* *
 * XPay Server SDK
 * 说明：
 * 以下代码只是为了方便商户测试而提供的样例代码，商户可以根据自己网站的需要，按照技术文档编写, 并非一定要使用该代码。
 * 该代码仅供学习和研究 XPay SDK 使用，只是提供一个参考。
 */
package sub_app

import (
	"time"

	"github.com/xxxpay/xpay-go/demo/common"
	xpay "github.com/xxxpay/xpay-go/xpay"
	"github.com/xxxpay/xpay-go/xpay/app"
)

var Demo = new(AppDemo)

type AppDemo struct {
	demoAppID    string
	demoSubappID string
}

func (c *AppDemo) Setup(app string) {
	c.demoAppID = app
}

// 创建子商户应用 sub_app
func (c *AppDemo) New() (*xpay.SubApp, error) {
	params := &xpay.SubAppParams{
		DisplayName: "sub_app_wuxinyi_test" + time.Now().Format("060102150405"),
		User:        "wuxinyi_001" + time.Now().Format("060102150405"),
		Metadata: map[string]interface{}{
			"key": "value",
		},
	}
	return app.New(c.demoAppID, params)
}

//查询子商户应用 sub_app
func (c *AppDemo) Get() (*xpay.SubApp, error) {
	return app.Get(c.demoAppID, c.demoSubappID)
}

// 查询子商户应用列表
func (c *AppDemo) List() (*xpay.SubAppList, error) {
	params := &xpay.PagingParams{}
	params.Filters.AddFilter("per_page", "", "3")
	return app.List(c.demoAppID, params)
}

//更新子商户应用
func (c *AppDemo) Update() (*xpay.SubApp, error) {
	params := &xpay.SubAppUpdateParams{
		ParentApp: "app_1Gqj58ynP0mHeX1q",
		Metadata: map[string]interface{}{
			"hello": "world",
		},
	}
	return app.Update(c.demoAppID, c.demoSubappID, *params)
}

//删除子商户应用
func (c *AppDemo) Delete() (*xpay.DeleteResult, error) {
	return app.Delete(c.demoAppID, c.demoSubappID)
}

func (c *AppDemo) Run() {
	subApp, err := c.New()
	common.Response(subApp, err)
	c.demoSubappID = subApp.ID
	subApp, err = c.Get()
	common.Response(subApp, err)
	subAppList, err := c.List()
	common.Response(subAppList, err)
	subApp, err = c.Update()
	common.Response(subApp, err)
	delResult, err := c.Delete()
	common.Response(delResult, err)
}
