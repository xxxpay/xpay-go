package couponTemplate

import (
	"fmt"
	"log"
	"net/url"

	"github.com/xxxpay/xpay-go"
)

const (
	CASH_COUPON     = 1
	DISCOUNT_COUPON = 2
)

type Client struct {
	B   xpay.Backend
	Key string
}

func getC() Client {
	return Client{xpay.GetBackend(xpay.APIBackend), xpay.Key}
}

// 创建优惠券模板
func New(appId string, params *xpay.CouponTmplParams) (*xpay.CouponTmpl, error) {
	return getC().New(appId, params)
}

func (c Client) New(appId string, params *xpay.CouponTmplParams) (*xpay.CouponTmpl, error) {
	paramsString, _ := xpay.JsonEncode(params)
	if xpay.LogLevel > 2 {
		log.Printf("params of create coupon_template request to xpay is :\n %v\n ", string(paramsString))
	}

	couponTemplate := &xpay.CouponTmpl{}

	err := c.B.Call("POST", fmt.Sprintf("/apps/%s/coupon_templates", appId), c.Key, nil, paramsString, couponTemplate)
	if err != nil {
		if xpay.LogLevel > 0 {
			log.Printf("%v\n", err)
		}
		return nil, err
	}
	return couponTemplate, err

}

//查询指定的优惠券模板
func Get(appId, couponTmplId string) (*xpay.CouponTmpl, error) {
	return getC().Get(appId, couponTmplId)
}

func (c Client) Get(appId, couponTmplId string) (*xpay.CouponTmpl, error) {
	var body *url.Values
	body = &url.Values{}
	couponTmpl := &xpay.CouponTmpl{}

	err := c.B.Call("GET", fmt.Sprintf("/apps/%s/coupon_templates/%s", appId, couponTmplId), c.Key, body, nil, couponTmpl)
	if err != nil {
		if xpay.LogLevel > 0 {
			log.Printf("Get Coupon Template error: %v\n", err)
		}
	}
	return couponTmpl, err
}

//更新优惠券模板
func Update(appId, couponTmplId string, params *xpay.CouponTmplUpdateParams) (*xpay.CouponTmpl, error) {
	return getC().Update(appId, couponTmplId, params)
}

func (c Client) Update(appId, couponTmplId string, params *xpay.CouponTmplUpdateParams) (*xpay.CouponTmpl, error) {
	paramsString, _ := xpay.JsonEncode(params)
	if xpay.LogLevel > 2 {
		log.Printf("params of update coupon template to xpay is :\n %v\n ", string(paramsString))
	}

	couponTmpl := &xpay.CouponTmpl{}

	err := c.B.Call("PUT", fmt.Sprintf("/apps/%s/coupon_templates/%s", appId, couponTmplId), c.Key, nil, paramsString, couponTmpl)
	if err != nil {
		if xpay.LogLevel > 0 {
			log.Printf("%v\n", err)
		}
		return nil, err
	}
	return couponTmpl, err
}

//删除优惠券模板

func Delete(appId, couponTmplId string) (*xpay.DeleteResult, error) {
	return getC().Delete(appId, couponTmplId)
}

func (c Client) Delete(appId, couponTmplId string) (*xpay.DeleteResult, error) {
	result := &xpay.DeleteResult{}

	err := c.B.Call("DELETE", fmt.Sprintf("/apps/%s/coupon_templates/%s", appId, couponTmplId), c.Key, nil, nil, result)
	if err != nil {
		if xpay.LogLevel > 0 {
			log.Printf("Delete Coupon Template error: %v\n", err)
		}
	}
	return result, err
}

//查询优惠券模板列表
func List(appId string, params *xpay.PagingParams) (*xpay.CouponTmplList, error) {
	return getC().List(appId, params)
}
func (c Client) List(appId string, params *xpay.PagingParams) (*xpay.CouponTmplList, error) {
	body := &url.Values{}
	params.Filters.AppendTo(body)

	couponTmplList := &xpay.CouponTmplList{}
	err := c.B.Call("GET", fmt.Sprintf("/apps/%s/coupon_templates", appId), c.Key, body, nil, couponTmplList)
	return couponTmplList, err
}

//优惠券模板下的优惠券列表
func CouponList(appId, couponTmplId string, params *xpay.PagingParams) (*xpay.CouponList, error) {
	return getC().CouponList(appId, couponTmplId, params)
}

func (c Client) CouponList(appId, couponTmplId string, params *xpay.PagingParams) (*xpay.CouponList, error) {
	body := &url.Values{}
	params.Filters.AppendTo(body)

	couponList := &xpay.CouponList{}
	err := c.B.Call("GET", fmt.Sprintf("/apps/%s/coupon_templates/%s/coupons", appId, couponTmplId), c.Key, body, nil, couponList)
	if err != nil {
		if xpay.LogLevel > 0 {
			log.Printf("Get Coupon error: %v\n", err)
		}
	}
	return couponList, err
}
