package payment

var Cb_wx_pub = map[string]interface{}{
	// 可选，指定支付方式，指定不能使用信用卡支付可设置为 no_credit 。
	"limit_pay": "no_credit",
	// 必填，用户在商户 appid 下的唯一标识。
	"open_id": "openidxxxxxxxxxxxx",
	// 必填，商品列表，字段解释：goods_name:商品名称，goods_num:数量。
	"goods_list": []map[string]interface{}{
		map[string]interface{}{
			"goods_name": "iPhone",
			"goods_num":  "1",
		},
		map[string]interface{}{
			"goods_name": "iPad",
			"goods_num":  "2",
		},
	},
}
