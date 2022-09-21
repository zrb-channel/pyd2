package pyd2

import (
	"context"
	"net/url"

	json "github.com/json-iterator/go"
)

type OrderStatus struct {
	Text  string
	State uint8
}

var (
	// // 状态0审批中1审批成功2佣金待提现3不符合支付要求4审批失败5佣金已提现 6.已拒绝
	//001-审批中
	//002-审批失败
	//003-审批通过
	//004-审批异常
	//B01-申请成功
	//B02-申请失败

	//A01-审批中
	//A02-审批失败
	//A03-审批通过
	//A04-审批作废
	//B01-申请成功
	//B02-申请失败
	preApplyStatusMapping = map[string]*OrderStatus{
		"001": {Text: "预授信审批中", State: 0},
		"002": {Text: "预授信审批失败", State: 6},
		"003": {Text: "预授信审批通过", State: 0},
		"004": {Text: "预授信审批异常", State: 4},
		"B01": {Text: "预授信审批成功", State: 0},
		"B02": {Text: "预授信审批失败", State: 6},
	}

	creditStatusMapping = map[string]*OrderStatus{
		"A01": {Text: "授信中", State: 0},
		"A02": {Text: "授信失败", State: 6},
		"A03": {Text: "授信通过", State: 1},
		"A04": {Text: "授信作废", State: 6},
		"B01": {Text: "授信成功", State: 1},
		"B02": {Text: "授信失败", State: 6},
	}

	applyStatusMapping = map[string]*OrderStatus{
		"A01": {Text: "放款中...", State: 0},
		"A02": {Text: "放款失败", State: 6},
		"A03": {Text: "放款通过", State: 1},
		"A04": {Text: "授信作废", State: 6},
		"B01": {Text: "放款成功", State: 1},
		"B02": {Text: "放款失败", State: 6},
	}
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
