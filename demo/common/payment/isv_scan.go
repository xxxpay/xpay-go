package payment

var Isv_scan = map[string]interface{}{
	// 必须，客户端软件中展示的条码值，扫码设备扫描获取。
	"scan_code": "286801346868493272",
	// 必须，终端号，要求不同终端此号码不一样，会显示在对账单中，如A01、SH008等。
	"terminal_id": "SH008",
	// 可选，商品列表，上送格式参照下面示例。
	// 字段解释：goods_id:商户定义商品编号（一般商品条码）unified_goods_id:统一商品编号(可选)，goods_name:商品名称，goods_num:数量，
	// price:单价(单位分)，goods_category:商品类目(可选)，body:商品描述信息(可选)，show_url:商品的展示网址(可选)
	// "goods_list": []map[string]interface{}{
	// 	map[string]interface{}{
	// 		"goods_id":         "iphone6s16G",
	// 		"unified_goods_id": "1001",
	// 		"goods_name":       "iPhone6s 16G",
	// 		"goods_num":        "1",
	// 		"price":            "528800",
	// 		"goods_category":   "123456",
	// 		"body":             "苹果手机16G",
	// 		"show_url":         "https://www.example.com",
	// 	},
	// 	map[string]interface{}{
	// 		"goods_id":         "iphone6s32G",
	// 		"unified_goods_id": "1002",
	// 		"goods_name":       "iPhone6s 32G",
	// 		"goods_num":        "1",
	// 		"price":            "608800",
	// 		"goods_category":   "123789",
	// 		"body":             "苹果手机32G",
	// 		"show_url":         "https://www.example.com",
	// 	},
	// },
}
