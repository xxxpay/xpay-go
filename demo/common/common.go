package common

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/xxxpay/xpay-go/demo/common/payment"
	"github.com/xxxpay/xpay-go/demo/common/transfer"
	"github.com/xxxpay/xpay-go/demo/common/withdrawal"
)

func Response(data interface{}, err error) {
	if err != nil {
		log.Fatalln("response error:", err)
		return
	}
	PrintResponse(data)
}

func PrintResponse(data interface{}) {
	content, _ := json.MarshalIndent(data, "", "  ")
	fmt.Println(string(content))
}

type extra struct {
	PaymentExtra  map[string]map[string]interface{}
	TransferExtra map[string]map[string]interface{}
	WithdrawExtra map[string]map[string]interface{}
}

var Extra = extra{
	PaymentExtra: PaymentExtra,
}

var PaymentExtra = map[string]map[string]interface{}{
	"alipay":           payment.Alipay,
	"alipay_pc_direct": payment.Alipay_pc_direct,
	"alipay_scan":      payment.Alipay_scan,
	"alipay_wap":       payment.Alipay_wap,
	"applepay_upacp":   payment.Applepay_upacp,
	"balance":          payment.Balance,
	"bfb_wap":          payment.Bfb_wap,
	"cb_alipay":        payment.Cb_alipay,
	"cb_wx":            payment.Cb_wx,
	"cb_wx_pub":        payment.Cb_wx_pub,
	"cb_wx_pub_qr":     payment.Cb_wx_pub_qr,
	"cb_wx_pub_scan":   payment.Cb_wx_pub_scan,
	"cmb_wallet":       payment.Cmb_wallet,
	"fqlpay_wap":       payment.Fqlpay_wap,
	"isv_qr":           payment.Isv_qr,
	"isv_scan":         payment.Isv_scan,
	"isv_wap":          payment.Isv_wap,
	"jdpay_wap":        payment.Jdpay_wap,
	"mmdpay_wap":       payment.Mmdpay_wap,
	"qgbc_wap":         payment.Qgbc_wap,
	"qpay":             payment.Qpay,
	"upacp":            payment.Upacp,
	"upacp_pc":         payment.Upacp_pc,
	"upacp_wap":        payment.Upacp_wap,
	"wx":               payment.Wx,
	"wx_lite":          payment.Wx_lite,
	"wx_pub":           payment.Wx_pub,
	"wx_pub_qr":        payment.Wx_pub_qr,
	"wx_pub_scan":      payment.Wx_pub_scan,
	"wx_wap":           payment.Wx_wap,
	"yeepay_wap":       payment.Yeepay_wap,
}

var TransferExtra = map[string]map[string]interface{}{
	"alipay":   transfer.Alipay,
	"allinpay": transfer.Allinpay,
	"balance":  transfer.Balance,
	"jdpay":    transfer.Jdpay,
	"unionpay": transfer.Unionpay,
	"wx_pub":   transfer.Wx_pub,
}

var WithdrawExtra = map[string]map[string]interface{}{
	"alipay":   withdrawal.Alipay,
	"allinpay": withdrawal.Allinpay,
	"jdpay":    withdrawal.Jdpay,
	"unionpay": withdrawal.Unionpay,
	"wx_pub":   withdrawal.Wx_pub,
}
