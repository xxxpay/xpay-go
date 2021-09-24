package coupon

import (
	"fmt"
	"log"
	"net/url"

	"github.com/xxxpay/xpay-go"
)

type Client struct {
	backend xpay.Backend
}

func NewClient(backend xpay.Backend) Client {
	return Client{backend: backend}
}

//创建优惠券
func (c Client) New(appId, userId string, params *xpay.CouponParams) (*xpay.Coupon, error) {
	paramsString, _ := xpay.JsonEncode(params)
	if xpay.LogLevel > 2 {
		log.Printf("params of create coupon request to xpay is :\n %v\n ", string(paramsString))
	}

	coupon := &xpay.Coupon{}

	err := c.backend.Call("POST", fmt.Sprintf("/apps/%s/users/%s/coupons", appId, userId), nil, paramsString, coupon)
	if err != nil {
		if xpay.LogLevel > 0 {
			log.Printf("%v\n", err)
		}
		return nil, err
	}
	return coupon, err
}

//批量创建优惠券
func (c Client) BatchNew(appId, couponTmplId string, params *xpay.BatchCouponParams) (*xpay.CouponList, error) {
	paramsString, _ := xpay.JsonEncode(params)
	if xpay.LogLevel > 2 {
		log.Printf("params of create coupons request to xpay is :\n %v\n ", string(paramsString))
	}

	couponList := &xpay.CouponList{}

	err := c.backend.Call("POST", fmt.Sprintf("/apps/%s/coupon_templates/%s/coupons", appId, couponTmplId), nil, paramsString, couponList)
	if err != nil {
		if xpay.LogLevel > 0 {
			log.Printf("%v\n", err)
		}
		return nil, err
	}
	return couponList, err
}

//更新优惠券
func (c Client) Update(appId, userId, couponId string, params *xpay.CouponUpdateParams) (*xpay.Coupon, error) {
	paramsString, _ := xpay.JsonEncode(params)
	if xpay.LogLevel > 2 {
		log.Printf("params of update coupon  to xpay is :\n %v\n ", string(paramsString))
	}

	coupon := &xpay.Coupon{}

	err := c.backend.Call("PUT", fmt.Sprintf("/apps/%s/users/%s/coupons/%s", appId, userId, couponId), nil, paramsString, coupon)
	if err != nil {
		if xpay.LogLevel > 0 {
			log.Printf("%v\n", err)
		}
		return nil, err
	}
	return coupon, err
}

//删除优惠券
func (c Client) Delete(appId, userId, couponId string) (*xpay.DeleteResult, error) {
	result := &xpay.DeleteResult{}

	err := c.backend.Call("DELETE", fmt.Sprintf("/apps/%s/users/%s/coupons/%s", appId, userId, couponId), nil, nil, result)
	if err != nil {
		if xpay.LogLevel > 0 {
			log.Printf("Delete Coupon Template error: %v\n", err)
		}
	}
	return result, err
}

//查询指定的优惠券模板
func (c Client) Get(appId, userId, couponId string) (*xpay.Coupon, error) {
	var body *url.Values
	body = &url.Values{}
	coupon := &xpay.Coupon{}

	err := c.backend.Call("GET", fmt.Sprintf("/apps/%s/users/%s/coupons/%s", appId, userId, couponId), body, nil, coupon)
	if err != nil {
		if xpay.LogLevel > 0 {
			log.Printf("Get Coupon error: %v\n", err)
		}
	}
	return coupon, err
}

//用户的优惠券列表
func (c Client) UserList(appId, userId string, params *xpay.PagingParams) (*xpay.CouponList, error) {
	body := &url.Values{}
	params.Filters.AppendTo(body)

	couponList := &xpay.CouponList{}
	err := c.backend.Call("GET", fmt.Sprintf("/apps/%s/users/%s/coupons", appId, userId), body, nil, couponList)
	if err != nil {
		if xpay.LogLevel > 0 {
			log.Printf("Get Coupon error: %v\n", err)
		}
	}
	return couponList, err
}
