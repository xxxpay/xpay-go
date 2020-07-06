/* *
 * XPay Server SDK
 * 说明：
 * 以下代码只是为了方便商户测试而提供的样例代码，商户可以根据自己网站的需要，按照技术文档编写, 并非一定要使用该代码。
 * 该代码仅供学习和研究 XPay SDK 使用，只是提供一个参考。
 */
package coupon

import (
	xpay "github.com/xxxpay/xpay-go/xpay"
	"github.com/xxxpay/xpay-go/xpay/coupon"
)

var Demo = new(CouponDemo)

type CouponDemo struct {
	demoAppID        string
	demoCouponTmplID string
	demoUser         string
}

func (c *CouponDemo) Setup(app string) {
	c.demoAppID = app
	c.demoCouponTmplID = "300216111619300600019101"
	c.demoUser = "uid582d1756b1650"
}

//创建 Coupon 对象
func (c *CouponDemo) New() (*xpay.Coupon, error) {
	params := &xpay.CouponParams{
		Coupon_tmpl_id: c.demoCouponTmplID,
	}
	return coupon.New(c.demoAppID, "uid582d1756b1650", params)
}

//批量创建优惠券
func (c *CouponDemo) Batch() (*xpay.CouponList, error) {
	params := &xpay.BatchCouponParams{
		Users: []string{"xtest1@lucfish.com", "xtest2@lucfish.com", "xtest@lucfish.com"},
	}
	return coupon.BatchNew(c.demoAppID, c.demoCouponTmplID, params)
}

//查询 Coupon 对象
func (c *CouponDemo) Get() (*xpay.Coupon, error) {
	return coupon.Get(c.demoAppID, c.demoUser, c.demoCouponTmplID)
}

//查询 Coupon 对象列表
func (c *CouponDemo) List() (*xpay.CouponList, error) {
	params := &xpay.PagingParams{}
	params.Filters.AddFilter("page", "", "1")     //页码，取值范围：1~1000000000；默认值为"1"
	params.Filters.AddFilter("per_page", "", "2") //每页数量，取值范围：1～100；默认值为"20"
	// params.Filters.AddFilter("redeemed", "", "false") //查询用户未核销优惠券

	return coupon.UserList(c.demoAppID, c.demoUser, params)
}

//更新 Coupon 对象
func (c *CouponDemo) Update() (*xpay.Coupon, error) {
	params := &xpay.CouponUpdateParams{
		Metadata: map[string]interface{}{
			"key": "value",
		},
	}
	return coupon.Update(c.demoAppID, c.demoUser, c.demoCouponTmplID, params)
}

//删除 Coupon 对象
func (c *CouponDemo) Delete() (*xpay.DeleteResult, error) {
	return coupon.Delete(c.demoAppID, c.demoUser, c.demoCouponTmplID)
}

func (c *CouponDemo) Run() {
	c.New()
	c.Batch()
	c.Get()
	c.List()
	c.Update()
	c.Delete()
}
