package user

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

func New(appId string, params *xpay.UserParams) (*xpay.User, error) {
	return getC().New(appId, params)
}

func (c Client) New(appId string, params *xpay.UserParams) (*xpay.User, error) {
	paramsString, errs := xpay.JsonEncode(params)
	if errs != nil {
		if xpay.LogLevel > 0 {
			log.Printf("UserParams Marshall Errors is : %q/n", errs)
		}
		return nil, errs
	}
	if xpay.LogLevel > 2 {
		log.Printf("params of create user is :\n %v\n ", string(paramsString))
	}

	user := &xpay.User{}
	err := c.B.Call("POST", fmt.Sprintf("/apps/%s/users", appId), c.Key, nil, paramsString, user)
	return user, err
}

func Get(appId, userId string) (*xpay.User, error) {
	return getC().Get(appId, userId)
}

func (c Client) Get(appId, userId string) (*xpay.User, error) {
	user := &xpay.User{}

	err := c.B.Call("GET", fmt.Sprintf("/apps/%s/users/%s", appId, userId), c.Key, nil, nil, user)
	return user, err
}

func List(appId string, params *xpay.PagingParams) (*xpay.UserList, error) {
	return getC().List(appId, params)
}
func (c Client) List(appId string, params *xpay.PagingParams) (*xpay.UserList, error) {
	body := &url.Values{}
	params.Filters.AppendTo(body)

	userList := &xpay.UserList{}
	err := c.B.Call("GET", fmt.Sprintf("/apps/%s/users", appId), c.Key, body, nil, userList)
	return userList, err
}

func Update(appId, userId string, params map[string]interface{}) (*xpay.User, error) {
	return getC().Update(appId, userId, params)
}

func (c Client) Update(appId, userId string, params map[string]interface{}) (*xpay.User, error) {
	paramsString, _ := xpay.JsonEncode(params)
	if xpay.LogLevel > 2 {
		log.Printf("params of update user  to xpay is :\n %v\n ", string(paramsString))
	}

	user := &xpay.User{}

	err := c.B.Call("PUT", fmt.Sprintf("/apps/%s/users/%s", appId, userId), c.Key, nil, paramsString, user)
	if err != nil {
		if xpay.LogLevel > 0 {
			log.Printf("%v\n", err)
		}
		return nil, err
	}
	return user, err
}
