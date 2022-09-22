package pyd2

import (
	"context"
	"net/url"

	json "github.com/json-iterator/go"
)

// BeforeNotify
// @param body
// @date 2022-05-17 19:27:06
func BeforeNotify[T any](body []byte) (*T, error) {
	req := &NotifyRequest{}
	if err := json.Unmarshal(body, req); err != nil {
		return nil, err
	}

	bizContent, err := url.QueryUnescape(req.BizContent)
	if err != nil {
		return nil, err
	}

	result := new(T)

	return result, json.Unmarshal([]byte(bizContent), result)
}

// NotifyPreCredit
// @param ctx
// @param body
// @date 2022-05-17 19:27:02
func NotifyPreCredit(ctx context.Context, body []byte) error {
	return nil
}

// NotifyCredit
// @param ctx
// @param body
// @date 2022-05-17 19:27:01
func NotifyCredit(ctx context.Context, body []byte) error {

	return nil
}

// NotifyApply
// @param ctx
// @param body
// @date 2022-05-17 19:27:00
func NotifyApply(ctx context.Context, body []byte) error {
	return nil
}
