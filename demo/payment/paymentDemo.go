/* *
 * XPay Server SDK
 * 说明：
 * 以下代码只是为了方便商户测试而提供的样例代码，商户可以根据自己网站的需要，按照技术文档编写, 并非一定要使用该代码。
 * 该代码仅供学习和研究 XPay SDK 使用，只是提供一个参考。
 */
package payment

import (
	//"encoding/base64"

	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/xxxpay/xpay-go/demo/common"
	xpay "github.com/xxxpay/xpay-go/xpay"
	"github.com/xxxpay/xpay-go/xpay/payment"
	//"io/ioutil"
)

var Demo = new(PaymentDemo)

type PaymentDemo struct {
	demoAppID   string
	demoChannel string
	demoPayment string
}

func (c *PaymentDemo) Setup(app string) {
	c.demoAppID = app
	c.demoChannel = "alipay"
	c.demoPayment = "f6676f79487842f394db045e93395359"
}

func (c *PaymentDemo) New() (*xpay.Payment, error) {
	//针对metadata字段，可以在每一个 payment 对象中加入订单的一些详情，如颜色、型号等属性
	metadata := make(map[string]interface{})
	metadata["color"] = "red"
	//metadata["type"] = "shoes"
	//metadata["size"] = "40"

	//这里是随便设置的随机数作为订单号，仅作示例，该方法可能产生相同订单号，商户请自行生成，不要纠结该方法。
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	orderno := r.Intn(999999999999999)

	params := &xpay.PaymentParams{
		Order_no:  strconv.Itoa(orderno),
		App:       xpay.App{Id: c.demoAppID},
		Amount:    1000,
		Channel:   c.demoChannel,
		Currency:  "cny",
		Client_ip: "127.0.0.1",
		Subject:   "Your Subject",
		Body:      "Your Body",
		Extra:     common.Extra.PaymentExtra[c.demoChannel],
		Metadata:  metadata,
	}

	//返回的第一个参数是 payment 对象，你需要将其转换成 json 给客户端，或者客户端接收后转换。
	return payment.New(params)
}

// 查询 payment 对象
func (c *PaymentDemo) Get() (*xpay.Payment, error) {
	return payment.Get(c.demoPayment)
}

// 撤销 payment，此接口仅接受线下 isv_scan、isv_wap、isv_qr 渠道的订单调用
func (c *PaymentDemo) Reverse() (*xpay.Payment, error) {
	return payment.Reverse(c.demoPayment)
}

// 查询 payment 对象列表
func (c *PaymentDemo) List() *payment.Iter {
	params := &xpay.PaymentListParams{}
	params.Filters.AddFilter("limit", "", "3")
	//设置是不是只需要之前设置的 limit 这一个查询参数
	params.Single = true
	return payment.List(c.demoAppID, params)
}

func (c *PaymentDemo) Run() {
	payment, err := c.New()
	if err != nil {
		log.Printf("Create payment error: %v\n", err)
		return
	}
	log.Printf("Create payment: %v\n", payment)

	c.Get()
	c.List()
}
