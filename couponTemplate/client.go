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
	backend xpay.Backend
}

func NewClient(backend xpay.Backend) Client {
	return Client{backend: backend}
}

// 创建优惠券模板
func (c Client) New(appId string, params *xpay.CouponTmplParams) (*xpay.CouponTmpl, error) {
	paramsString, _ := xpay.JsonEncode(params)
	if xpay.LogLevel > 2 {
		log.Printf("params of create coupon_template request to xpay is :\n %v\n ", string(paramsString))
	}

	couponTemplate := &xpay.CouponTmpl{}

	err := c.backend.Call("POST", fmt.Sprintf("/apps/%s/coupon_templates", appId), nil, paramsString, couponTemplate)
	if err != nil {
		if xpay.LogLevel > 0 {
			log.Printf("%v\n", err)
		}
		return nil, err
	}
	return couponTemplate, err

}

//查询指定的优惠券模板
func (c Client) Get(appId, couponTmplId string) (*xpay.CouponTmpl, error) {
	var body *url.Values
	body = &url.Values{}
	couponTmpl := &xpay.CouponTmpl{}

	err := c.backend.Call("GET", fmt.Sprintf("/apps/%s/coupon_templates/%s", appId, couponTmplId), body, nil, couponTmpl)
	if err != nil {
		if xpay.LogLevel > 0 {
			log.Printf("Get Coupon Template error: %v\n", err)
		}
	}
	return couponTmpl, err
}

//更新优惠券模板
func (c Client) Update(appId, couponTmplId string, params *xpay.CouponTmplUpdateParams) (*xpay.CouponTmpl, error) {
	paramsString, _ := xpay.JsonEncode(params)
	if xpay.LogLevel > 2 {
		log.Printf("params of update coupon template to xpay is :\n %v\n ", string(paramsString))
	}

	couponTmpl := &xpay.CouponTmpl{}

	err := c.backend.Call("PUT", fmt.Sprintf("/apps/%s/coupon_templates/%s", appId, couponTmplId), nil, paramsString, couponTmpl)
	if err != nil {
		if xpay.LogLevel > 0 {
			log.Printf("%v\n", err)
		}
		return nil, err
	}
	return couponTmpl, err
}

//删除优惠券模板
func (c Client) Delete(appId, couponTmplId string) (*xpay.DeleteResult, error) {
	result := &xpay.DeleteResult{}

	err := c.backend.Call("DELETE", fmt.Sprintf("/apps/%s/coupon_templates/%s", appId, couponTmplId), nil, nil, result)
	if err != nil {
		if xpay.LogLevel > 0 {
			log.Printf("Delete Coupon Template error: %v\n", err)
		}
	}
	return result, err
}

//查询优惠券模板列表
func (c Client) List(appId string, params *xpay.PagingParams) (*xpay.CouponTmplList, error) {
	body := &url.Values{}
	params.Filters.AppendTo(body)

	couponTmplList := &xpay.CouponTmplList{}
	err := c.backend.Call("GET", fmt.Sprintf("/apps/%s/coupon_templates", appId), body, nil, couponTmplList)
	return couponTmplList, err
}

//优惠券模板下的优惠券列表
func (c Client) CouponList(appId, couponTmplId string, params *xpay.PagingParams) (*xpay.CouponList, error) {
	body := &url.Values{}
	params.Filters.AppendTo(body)

	couponList := &xpay.CouponList{}
	err := c.backend.Call("GET", fmt.Sprintf("/apps/%s/coupon_templates/%s/coupons", appId, couponTmplId), body, nil, couponList)
	if err != nil {
		if xpay.LogLevel > 0 {
			log.Printf("Get Coupon error: %v\n", err)
		}
	}
	return couponList, err
}
