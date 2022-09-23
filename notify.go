package pyd2

import (
	"context"
	"net/url"

	json "github.com/json-iterator/go"
)

type NotifyHandlers interface {
	OnPreCredit(ctx context.Context, req *NotifyPreApply) error

	OnCredit(ctx context.Context, req *NotifyCreditRequest) error

	OnApply(ctx context.Context, req *NotifyApplyRequest) error
}

var notifyHandlers NotifyHandlers

func RegisterNotifyHandlers(handlers NotifyHandlers) {
	notifyHandlers = handlers
}

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
	if err := ctx.Err(); err != nil {
		return err
	}

	v, err := BeforeNotify[NotifyPreApply](body)
	if err != nil {
		return err
	}
	if notifyHandlers == nil {
		return nil
	}

	return notifyHandlers.OnPreCredit(ctx, v)
}

// NotifyCredit
// @param ctx
// @param body
// @date 2022-05-17 19:27:01
func NotifyCredit(ctx context.Context, body []byte) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	v, err := BeforeNotify[NotifyCreditRequest](body)
	if err != nil {
		return err
	}

	if notifyHandlers == nil {
		return nil
	}

	return notifyHandlers.OnCredit(ctx, v)
}

// NotifyApply
// @param ctx
// @param body
// @date 2022-05-17 19:27:00
func NotifyApply(ctx context.Context, body []byte) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	v, err := BeforeNotify[NotifyApplyRequest](body)
	if err != nil {
		return err
	}

	if notifyHandlers == nil {
		return nil
	}

	return notifyHandlers.OnApply(ctx, v)
}
