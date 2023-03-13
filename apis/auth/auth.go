// Copyright 2020 FastWeGo
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package auth 登录
package auth

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/jackluo2012/microapp"
)

const (
	apiCode2Session = "/api/apps/jscode2session"
)

/*
code2Session

通过login接口获取到登录凭证后，开发者可以通过服务器发送请求的方式获取 session_key 和 openId。

See: https://microapp.bytedance.com/docs/zh-CN/mini-app/develop/server/log-in/code-2-session

GET https://developer.toutiao.com/api/apps/jscode2session
*/
// RespCode2Session struct
type RespCode2Session struct {
	microapp.Error

	Openid          string `json:"openid"`           // 用户在当前小程序的 ID，如果请求时有 code 参数才会返回
	SessionKey      string `json:"session_key"`      // 会话密钥，如果请求时有 code 参数才会返回
	AnonymousOpenid string `json:"anonymous_openid"` // 匿名用户在当前小程序的 ID，如果请求时有 anonymous_code 参数才会返回
}

func Code2Session(ctx *microapp.MicroApp, params url.Values) (RespCode2Session, error) {
	params.Add("appid", ctx.Config.AppId)
	params.Add("secret", ctx.Config.AppSecret)
	xresp := RespCode2Session{}
	resp, err := ctx.Client.HTTPGet(apiCode2Session + "?" + params.Encode())
	if err != nil {
		return xresp, err
	}

	err = json.Unmarshal(resp, &resp)
	if err != nil {
		return xresp, err
	}

	if xresp.Code != 0 {
		err = fmt.Errorf("Code2Session error : errcode=%v , errmsg=%v", xresp.Code, xresp.Msg)
		return xresp, err
	}
	return xresp, err
}
