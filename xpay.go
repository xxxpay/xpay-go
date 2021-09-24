package xpay

import (
	"bytes"
	"net/url"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

const (
	// 当前版本的api生成生成时间
	apiVersion = "2021-06-30"
	// httpclient等待时间
	defaultHTTPTimeout = 80 * time.Second
)

var (
	// 当前版本的api地址
	APIBase = "https://api.xpay.ucfish.com/xpay/v2"
	// 默认错误信息返回语言
	AcceptLanguage = "zh-CN"
	// xpay api统一需要通过Authentication（http BasicAuth），需要在调用时赋值
	Key string
	// loglevel 是 debug 模式开关.
	// 0: no logging
	// 1: errors only
	// 2: errors + informational (default)
	// 3: errors + informational + debug
	LogLevel          = 2
	AccountPrivateKey string
	OsInfo            string
)

// 定义统一后端处理接口
type Backend interface {
	Call(method, path string, body *url.Values, params []byte, v interface{}) error
}

// 获取当前sdk的版本
func Version() string {
	return "3.2.1"
}

func init() {
	var uname string
	switch runtime.GOOS {
	case "windows":
		uname = "windows"
	default:
		cmd := exec.Command("uname", "-a")
		cmd.Stdin = strings.NewReader("some input")
		var out bytes.Buffer
		var stderr bytes.Buffer
		cmd.Stdout = &out
		cmd.Stderr = &stderr
		_ = cmd.Run()
		uname = out.String()
	}
	m := map[string]interface{}{
		"lang":             "golang",
		"lang_version":     runtime.Version(),
		"bindings_version": Version(),
		"publisher":        "xpay",
		"uname":            uname,
	}
	content, _ := JsonEncode(m)
	OsInfo = string(content)
}
