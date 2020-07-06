/* *
 * XPay Server SDK
 * 说明：
 * 以下代码只是为了方便商户测试而提供的样例代码，商户可以根据自己网站的需要，按照技术文档编写, 并非一定要使用该代码。
 * 该代码仅供学习和研究 XPay SDK 使用，只是提供一个参考。
 */
package user

import (
	"fmt"
	"math/rand"
	"time"

	xpay "github.com/xxxpay/xpay-go/xpay"
	"github.com/xxxpay/xpay-go/xpay/user"
)

var Demo = new(UserDemo)

type UserDemo struct {
	demoAppID string
}

func (c *UserDemo) Setup(app string) {
	c.demoAppID = app
}

// 创建 user 对象
func (c *UserDemo) New() (*xpay.User, error) {
	//这里是随便设置的随机数作为用户唯一标识，仅作示例
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	userId := r.Intn(999999999999999)

	params := &xpay.UserParams{
		ID:      fmt.Sprintf("%d", userId),
		Address: "中国.上海.浦东",
		Email:   "demo@lucfish.com",
	}

	return user.New("app_1Gqj58ynP0mHeX1q", params)
}

// 查询 user 对象
func (c *UserDemo) Get() (*xpay.User, error) {
	return user.Get("app_1Gqj58ynP0mHeX1q", "test_user_002")
}

// 更新 user 对象
func (c *UserDemo) Update() (*xpay.User, error) {
	params := map[string]interface{}{
		"address": "新地址",
		"email":   "123@lucfish.com",
		//"disabled":false 是否禁用。使用该参数时，不能同时使用其他参数。
	}
	return user.Update("app_1Gqj58ynP0mHeX1q", "test_user_002", params)
}

// 查询列表
func (c *UserDemo) List() (*xpay.UserList, error) {
	params := &xpay.PagingParams{}
	params.Filters.AddFilter("page", "", "1")     //取第一页数据
	params.Filters.AddFilter("per_page", "", "2") //每页两个User对象
	//params.Filters.AddFilter("created", "", "1475127952")
	return user.List("app_1Gqj58ynP0mHeX1q", params)
}

func (c *UserDemo) Run() {
	c.New()
	c.Get()
	c.List()
	c.Update()
}
