package user

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
	err := c.backend.Call("POST", fmt.Sprintf("/apps/%s/users", appId), nil, paramsString, user)
	return user, err
}

func (c Client) Get(appId, userId string) (*xpay.User, error) {
	user := &xpay.User{}

	err := c.backend.Call("GET", fmt.Sprintf("/apps/%s/users/%s", appId, userId), nil, nil, user)
	return user, err
}
func (c Client) List(appId string, params *xpay.PagingParams) (*xpay.UserList, error) {
	body := &url.Values{}
	params.Filters.AppendTo(body)

	userList := &xpay.UserList{}
	err := c.backend.Call("GET", fmt.Sprintf("/apps/%s/users", appId), body, nil, userList)
	return userList, err
}

func (c Client) Update(appId, userId string, params map[string]interface{}) (*xpay.User, error) {
	paramsString, _ := xpay.JsonEncode(params)
	if xpay.LogLevel > 2 {
		log.Printf("params of update user  to xpay is :\n %v\n ", string(paramsString))
	}

	user := &xpay.User{}

	err := c.backend.Call("PUT", fmt.Sprintf("/apps/%s/users/%s", appId, userId), nil, paramsString, user)
	if err != nil {
		if xpay.LogLevel > 0 {
			log.Printf("%v\n", err)
		}
		return nil, err
	}
	return user, err
}
