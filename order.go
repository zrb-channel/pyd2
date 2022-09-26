package pyd2

import (
	"context"
	"errors"
	"fmt"
	json "github.com/json-iterator/go"
	"github.com/zrb-channel/utils"
	log "github.com/zrb-channel/utils/logger"
	"go.uber.org/zap"
	"net/http"
)

var (
	headers = map[string]string{
		"Content-Type": "application/json; charset=utf-8",
	}
)

// CreateOrder 订单信息提交
// @param ctx
// @param order
// @date 2022-05-17 19:27:40
func CreateOrder(ctx context.Context, conf *Config, req *CreateOrderRequest) (*CreateOrderResponse, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	req.Channel = conf.Channel

	body, err := NewRequest(conf, "S000701", req)

	fields := map[string]any{
		"req":  req,
		"body": body,
	}
	if err != nil {
		log.WithError(err).Error("[浦慧税贷]-[订单信息提交]-创建请求失败", zap.Any("data", fields))
		return nil, err
	}
	addr := fmt.Sprintf(Addr, "S000701")
	resp, err := utils.Request(ctx).SetHeaders(headers).SetBody(body).Post(addr)
	if err != nil {
		log.WithError(err).Error("[浦慧税贷]-[订单信息提交]-请求失败", zap.Any("data", fields))
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		fields["status"] = resp.StatusCode()
		log.WithError(err).Error("[浦慧税贷]-[订单信息提交]-响应状态码有误", zap.Any("data", fields))
		return nil, errors.New(resp.Status())
	}

	fields["resp"] = resp.String()
	result := &BaseResponse[string]{}
	if err = json.Unmarshal(resp.Body(), result); err != nil {
		log.WithError(err).Error("[浦慧税贷]-[订单信息提交]-响应数据解析未BaseResponse失败", zap.Any("data", fields))
		return nil, err
	}

	fields["result"] = result
	if result.Code != "000000" {
		log.WithError(err).Error("[浦慧税贷]-[订单信息提交]-code有误", zap.Any("data", fields))
		return nil, errors.New(parseMessage(result.Message))
	}

	data := &CreateOrderResponse{}
	if err = json.Unmarshal([]byte(result.Data), data); err != nil {
		log.WithError(err).Error("[浦慧税贷]-[订单信息提交]-响应数据解析未CreateOrderResponse失败", zap.Any("data", fields))
		return nil, err
	}

	return data, nil
}

// Redirect 获取跳转链接
// @param ctx
// @param order
// @date 2022-05-17 19:27:39
func Redirect(ctx context.Context, conf *Config, req *RedirectRequest) (string, error) {
	if err := ctx.Err(); err != nil {
		return "", err
	}

	req.ChannelAgent = conf.Channel

	addr := fmt.Sprintf(Addr, "S000706")

	body, err := NewRequest(conf, "S000706", req)

	fields := map[string]any{
		"req":  req,
		"body": body,
	}
	if err != nil {
		log.WithError(err).Error("[浦慧税贷]-[获取跳转链接]-创建请求失败", zap.Any("data", fields))
		return "", err
	}

	resp, err := utils.Request(ctx).SetHeaders(headers).SetBody(body).Post(addr)
	if err != nil {
		log.WithError(err).Error("[浦慧税贷]-[获取跳转链接]-请求失败", zap.Any("data", fields))
		return "", err
	}

	if resp.StatusCode() != http.StatusOK {
		fields["status"] = resp.StatusCode()
		log.WithError(err).Error("[浦慧税贷]-[获取跳转链接]-响应状态码有误", zap.Any("data", fields))
		return "", errors.New(resp.Status())
	}

	fields["resp"] = resp.String()
	result := &BaseResponse[string]{}
	if err = json.Unmarshal(resp.Body(), result); err != nil {
		log.WithError(err).Error("[浦慧税贷]-[获取跳转链接]-响应数据解析未BaseResponse失败", zap.Any("data", fields))
		return "", err
	}

	fields["result"] = result

	if result.Code == "000500" {
		return "", errors.New("未查询到相关订单信息")
	}

	if result.Code != "000000" {
		log.WithError(err).Error("[浦慧税贷]-[获取跳转链接]-code有误", zap.Any("data", fields))
		return "", errors.New(result.Message)
	}

	data := &JumpResponse{}
	if err = json.Unmarshal([]byte(result.Data), data); err != nil {
		log.WithError(err).Error("Redirect JumpResponse Unmarshal Error")
		return "", err
	}

	log.WithError(err).Info("[浦慧税贷]-[获取跳转链接]-成功", zap.Any("data", fields))
	return data.JumpUrl, nil
}

// parseMessage
// @param msg
// @date 2022-09-21 21:19:18
func parseMessage(msg string) string {
	var code int
	var message string

	if _, err := fmt.Sscanf(msg, "rescd:%d resmsg:%s", &code, &message); err != nil {
		return msg
	}

	if message == "" {
		return msg
	}
	return message
}
