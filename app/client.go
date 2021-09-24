package app

import (
	"fmt"
	"github.com/xxxpay/xpay-go"
	"log"
	"net/url"
)

type Client struct {
	backend xpay.Backend
}

func NewClient(backend xpay.Backend) Client {
	return Client{backend: backend}
}

/*
* 创建子商户对象
* @param appId string
* @param params SubAppParams
* @return SubApp
 */
func (c Client) New(appId string, params *xpay.SubAppParams) (*xpay.SubApp, error) {
	paramsString, errs := xpay.JsonEncode(params)
	if errs != nil {
		if xpay.LogLevel > 0 {
			log.Printf("SubAppParams Marshall Errors is : %q/n", errs)
		}
		return nil, errs
	}
	if xpay.LogLevel > 2 {
		log.Printf("params of create sub_app is :\n %v\n ", string(paramsString))
	}

	subApp := &xpay.SubApp{}
	err := c.backend.Call("POST", fmt.Sprintf("/apps/%s/sub_apps", appId), nil, paramsString, subApp)
	return subApp, err
}

/*
* 查询子商户对象
* @param appId string
* @param subAppId string
* @return SubApp
 */
func (c Client) Get(appId, subAppId string) (*xpay.SubApp, error) {
	subApp := &xpay.SubApp{}

	err := c.backend.Call("GET", fmt.Sprintf("/apps/%s/sub_apps/%s", appId, subAppId), nil, nil, subApp)
	return subApp, err
}

/*
* 查询子商户对象列表
* @param appId string
* @param params PagingParams
* @return SubAppList
 */
func (c Client) List(appId string, params *xpay.PagingParams) (*xpay.SubAppList, error) {
	body := &url.Values{}
	params.Filters.AppendTo(body)

	subList := &xpay.SubAppList{}
	err := c.backend.Call("GET", fmt.Sprintf("/apps/%s/sub_apps", appId), body, nil, subList)
	return subList, err
}

/*
* 更新子商户对象
* @param appId string
* @param subAppId string
* @param params SubAppUpdateParams
* @return SubApp
 */
func (c Client) Update(appId, subAppId string, params xpay.SubAppUpdateParams) (*xpay.SubApp, error) {
	paramsString, _ := xpay.JsonEncode(params)
	if xpay.LogLevel > 2 {
		log.Printf("params of update SubApp  to xpay is :\n %v\n ", string(paramsString))
	}

	subApp := &xpay.SubApp{}

	err := c.backend.Call("PUT", fmt.Sprintf("/apps/%s/sub_apps/%s", appId, subAppId), nil, paramsString, subApp)
	if err != nil {
		if xpay.LogLevel > 0 {
			log.Printf("%v\n", err)
		}
		return nil, err
	}
	return subApp, err
}

/*
* 删除子商户对象
* @param appId string
* @param subAppId string
* @return DeleteResult
 */
func (c Client) Delete(appId, subAppId string) (*xpay.DeleteResult, error) {
	result := &xpay.DeleteResult{}

	err := c.backend.Call("DELETE", fmt.Sprintf("/apps/%s/sub_apps/%s", appId, subAppId), nil, nil, result)
	if err != nil {
		if xpay.LogLevel > 0 {
			log.Printf("%v\n", err)
		}
		return nil, err
	}
	return result, err
}
