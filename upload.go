package pyd2

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/zrb-channel/utils"

	log "github.com/zrb-channel/utils/logger"

	json "github.com/json-iterator/go"
	"go.uber.org/zap"
)

// UploadInvoice 发票数据上传
// @param ctx
// @param order
// @date 2022-05-17 19:23:57
func UploadInvoice(ctx context.Context, conf *Config, req *UploadInvoiceRequest) (*BaseResponse[string], error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	req.Channel = conf.Channel

	addr := fmt.Sprintf(Addr, "S000707")
	body, err := NewRequest(conf, "S000707", req)
	fields := map[string]any{
		"req":  req,
		"body": body,
	}

	if err != nil {
		log.WithError(err).Error("[浦慧税贷]-[发票数据上传]-创建请求失败", zap.Any("data", fields))
		return nil, err
	}

	resp, err := utils.Request(ctx).SetHeaders(headers).SetBody(body).Post(addr)
	if err != nil {
		log.WithError(err).Error("[浦慧税贷]-[发票数据上传]-请求失败", zap.Any("data", fields))
		return nil, errors.New("请重试")
	}

	if resp.StatusCode() != http.StatusOK {
		fields["status"] = resp.StatusCode()
		log.WithError(err).Error("[浦慧税贷]-[发票数据上传]-响应状态码有误", zap.Any("data", fields))
		return nil, errors.New("请重试")
	}

	fields["resp"] = resp.String()
	result := &BaseResponse[string]{}
	if err = json.Unmarshal(resp.Body(), result); err != nil {
		log.WithError(err).Error("[浦慧税贷]-[发票数据上传]-响应数据解析未BaseResponse失败", zap.Any("data", fields))
		return nil, errors.New("请重试")
	}

	fields["result"] = result

	log.Info("[浦慧税贷]-[发票数据上传]-成功", zap.Any("data", fields))
	return result, nil
}
