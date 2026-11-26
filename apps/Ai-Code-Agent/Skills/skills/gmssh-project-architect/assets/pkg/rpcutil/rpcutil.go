package rpcutil

import (
	"encoding/json"
	"fmt"

	"github.com/DemonZack/simplejrpc-go/core/gi18n"
	"github.com/DemonZack/simplejrpc-go/net/gsock"
)

// SetLanguage 设置请求语言
func SetLanguage(req *gsock.Request) {
	if req.RawRequest().Params == nil {
		return
	}
	var args struct {
		Lang string `json:"lang"`
	}
	_ = json.Unmarshal(*req.RawRequest().Params, &args)
	if args.Lang != "" {
		gi18n.Instance().SetLanguage(args.Lang)
	}
}

// ParseParams 解析请求参数(自动设置语言)
func ParseParams(req *gsock.Request, args any) error {
	SetLanguage(req)
	if req.RawRequest().Params == nil {
		return fmt.Errorf(gi18n.Instance().T("InvalidParams"))
	}
	if err := json.Unmarshal(*req.RawRequest().Params, args); err != nil {
		return fmt.Errorf(gi18n.Instance().T("InvalidParams"))
	}
	return nil
}
