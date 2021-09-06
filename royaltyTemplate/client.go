package royaltyTemplate

import (
	"fmt"
	"log"
	"net/url"

	"github.com/xxxpay/xpay-go"
)

type Client struct {
	B   xpay.Backend
	Key string
}

func getC() Client {
	return Client{xpay.GetBackend(xpay.APIBackend), xpay.Key}
}

// 创建分润模板
func New(params *xpay.RoyaltyTmplParams) (*xpay.RoyaltyTmpl, error) {
	return getC().New(params)
}

func (c Client) New(params *xpay.RoyaltyTmplParams) (*xpay.RoyaltyTmpl, error) {
	paramsString, _ := xpay.JsonEncode(params)
	if xpay.LogLevel > 2 {
		log.Printf("params of create royalty_template request to xpay is :\n %v\n ", string(paramsString))
	}

	royaltyTemplate := &xpay.RoyaltyTmpl{}

	err := c.B.Call("POST", "/royalty_templates", c.Key, nil, paramsString, royaltyTemplate)
	if err != nil {
		if xpay.LogLevel > 0 {
			log.Printf("%v\n", err)
		}
		return nil, err
	}
	return royaltyTemplate, err

}

//查询指定的分润模板
func Get(royaltyTmplId string) (*xpay.RoyaltyTmpl, error) {
	return getC().Get(royaltyTmplId)
}

func (c Client) Get(royaltyTmplId string) (*xpay.RoyaltyTmpl, error) {
	var body *url.Values
	body = &url.Values{}
	royaltyTmpl := &xpay.RoyaltyTmpl{}

	err := c.B.Call("GET", fmt.Sprintf("/royalty_templates/%s", royaltyTmplId), c.Key, body, nil, royaltyTmpl)
	if err != nil {
		if xpay.LogLevel > 0 {
			log.Printf("Get royalty Template error: %v\n", err)
		}
	}
	return royaltyTmpl, err
}

//更新分润模板
func Update(royaltyTmplId string, params *xpay.RoyaltyTmplUpdateParams) (*xpay.RoyaltyTmpl, error) {
	return getC().Update(royaltyTmplId, params)
}

func (c Client) Update(royaltyTmplId string, params *xpay.RoyaltyTmplUpdateParams) (*xpay.RoyaltyTmpl, error) {
	paramsString, _ := xpay.JsonEncode(params)
	if xpay.LogLevel > 2 {
		log.Printf("params of update royalty template to xpay is :\n %v\n ", string(paramsString))
	}

	royaltyTmpl := &xpay.RoyaltyTmpl{}

	err := c.B.Call("PUT", fmt.Sprintf("/royalty_templates/%s", royaltyTmplId), c.Key, nil, paramsString, royaltyTmpl)
	if err != nil {
		if xpay.LogLevel > 0 {
			log.Printf("%v\n", err)
		}
		return nil, err
	}
	return royaltyTmpl, err
}

//删除分润模板

func Delete(royaltyTmplId string) (*xpay.DeleteResult, error) {
	return getC().Delete(royaltyTmplId)
}

func (c Client) Delete(royaltyTmplId string) (*xpay.DeleteResult, error) {
	result := &xpay.DeleteResult{}

	err := c.B.Call("DELETE", fmt.Sprintf("/royalty_templates/%s", royaltyTmplId), c.Key, nil, nil, result)
	if err != nil {
		if xpay.LogLevel > 0 {
			log.Printf("Delete Royalty Template error: %v\n", err)
		}
	}
	return result, err
}

//查询分润模板列表
func List(params *xpay.PagingParams) (*xpay.RoyaltyTmplList, error) {
	return getC().List(params)
}
func (c Client) List(params *xpay.PagingParams) (*xpay.RoyaltyTmplList, error) {
	body := &url.Values{}
	params.Filters.AppendTo(body)

	royaltyTmplList := &xpay.RoyaltyTmplList{}
	err := c.B.Call("GET", "/royalty_templates", c.Key, body, nil, royaltyTmplList)
	return royaltyTmplList, err
}
