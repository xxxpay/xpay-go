package xpay

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// ApiBackend api相关的后端类型
type ApiBackend struct {
	URL        string
	HTTPClient *http.Client
}

// Call 后端处理请求方法
func (s *ApiBackend) Call(method, path, key string, form *url.Values, params []byte, v interface{}) error {
	var body io.Reader
	if strings.ToUpper(method) == "POST" || strings.ToUpper(method) == "PUT" {
		body = bytes.NewBuffer(params)

		if CustomQuery != nil && len(CustomQuery) > 0 {
			path += "?" + CustomQuery.Encode()
		}
	}

	if strings.ToUpper(method) == "GET" || strings.ToUpper(method) == "DELETE" {
		qs := make(url.Values)

		if CustomQuery != nil && len(CustomQuery) > 0 {
			for k, values := range CustomQuery {
				for _, val := range values {
					qs.Add(k, val)
				}
			}
		}

		if form != nil && len(*form) > 0 {
			for k, values := range *form {
				for _, val := range values {
					qs.Add(k, val)
				}
			}
		}

		if len(qs) > 0 {
			data := qs.Encode()
			path += "?" + data
		}
	}

	req, err := s.NewRequest(method, path, key, "application/json", body, params)

	if err != nil {
		return err
	}

	if err = s.Do(req, v); err != nil {
		return err
	}

	return nil
}

// NewRequest 建立http请求对象
func (s *ApiBackend) NewRequest(method, path, key, contentType string, body io.Reader, params []byte) (*http.Request, error) {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	path = s.URL + path
	req, err := http.NewRequest(method, path, body)
	if LogLevel > 2 {
		log.Printf("Request to xpay is : \n %v\n", req)
	}

	if err != nil {
		if LogLevel > 0 {
			log.Printf("Cannot create xpay request: %v\n", err)
		}
		return nil, err
	}
	var data string
	if strings.ToUpper(method) == "POST" || strings.ToUpper(method) == "PUT" {
		data = string(params)
	}
	requestTime := fmt.Sprintf("%d", time.Now().Unix())
	req.Header.Set("X-Request-Timestamp", requestTime)
	uri := req.URL.RequestURI()
	dataToBeSign := data + uri + requestTime

	log.Printf("RSA signature data: %s", data)
	log.Printf("RSA signature uri: %s", uri)
	log.Printf("RSA signature time: %s", requestTime)

	if len(AccountPrivateKey) > 0 {
		sign, err := GenSign([]byte(dataToBeSign), []byte(AccountPrivateKey))
		if err != nil {
			if LogLevel > 0 {
				log.Printf("Cannot create RSA signature: %v\n", err)
			}
			return nil, err
		}
		encodeSign := base64.StdEncoding.EncodeToString(sign)
		req.Header.Add("X-Signature", encodeSign)
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", key))
	req.Header.Add("XPay-Version", apiVersion)
	req.Header.Add("User-Agent", "xpay go sdk version:"+Version())
	req.Header.Add("XPay-Client-User-Agent", OsInfo)
	req.Header.Add("Content-Type", contentType)
	req.Header.Add("Accept-Language", AcceptLanguage)

	return req, nil
}

// Do 处理 http 请求
func (s *ApiBackend) Do(req *http.Request, v interface{}) error {
	if LogLevel > 1 {
		log.Printf("Requesting %v %v \n", req.Method, req.URL.String())
	}
	retryTimes := 1
	start := time.Now()

	var reqBody []byte
	var err error
	if req.Body != nil {
		reqBody, err = ioutil.ReadAll(req.Body)
		if err != nil {
			return err
		}
	}

	for i := 0; i <= retryTimes; i++ {
		req.Body = ioutil.NopCloser(bytes.NewBuffer(reqBody))
		res, err := s.HTTPClient.Do(req)
		if LogLevel > 0 {
			log.Printf("Request to xpay completed in %v\n", time.Since(start))
		}
		if err != nil {
			if LogLevel > 0 {
				log.Printf("Request to xpay failed: %v\n", err)
			}
			return err
		}
		defer res.Body.Close()
		if res.StatusCode == 502 {
			continue
		}

		resBody, err := ioutil.ReadAll(res.Body)
		if err != nil {
			if LogLevel > 0 {
				log.Printf("Cannot parse xpay response: %v\n", err)
			}
			return err
		}

		if res.StatusCode >= 400 {
			var errMap map[string]interface{}
			if err := JsonDecode(resBody, &errMap); err != nil {
				return err
			}

			if e, ok := errMap["error"]; !ok {
				err := errors.New(string(resBody))
				if LogLevel > 0 {
					log.Printf("Unparsable error returned from xpay: %v\n", err)
				}
				return err
			} else {
				root := e.(map[string]interface{})
				err := &Error{
					Type:           ErrorType(root["type"].(string)),
					Msg:            root["message"].(string),
					HTTPStatusCode: res.StatusCode,
				}

				if code, found := root["code"]; found {
					err.Code = ErrorCode(code.(string))
				}

				if param, found := root["param"]; found {
					err.Param = param.(string)
				}

				if LogLevel > 0 {
					log.Printf("Error encountered from xpay: %v\n", err)
				}
				return err
			}
		}

		if LogLevel > 2 {
			log.Printf("resBody from xpay API: \n%v\n", string(resBody))
		}

		if v != nil {
			return JsonDecode(resBody, v)
		}
		return nil
	}
	return nil
}
