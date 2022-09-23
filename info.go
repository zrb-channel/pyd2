package pyd2

import (
	"context"
	"errors"
	"fmt"
	json "github.com/json-iterator/go"
	"github.com/zrb-channel/utils"
	log "github.com/zrb-channel/utils/logger"
	"go.uber.org/zap"
)

// QueryInfo
// @param ctx
// @param conf
// @param req
// @date 2022-09-22 18:45:07
func QueryInfo(ctx context.Context, conf *Config, req *QueryInfoRequest) (*QueryInfoResponse, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	req.Channel = conf.ChannelCode

	body, err := NewRequest(conf, "S000705", req)

	fields := map[string]any{
		"req":  req,
		"body": body,
	}

	if err != nil {
		log.WithError(err).Error("[浦慧税贷]-[查询贷款信息]-创建请求失败", zap.Any("data", fields))
		return nil, err
	}

	addr := fmt.Sprintf(Addr, "S000705")

	resp, err := utils.Request(ctx).SetHeaders(headers).SetBody(body).Post(addr)
	if err != nil {
		log.WithError(err).Error("[浦慧税贷]-[查询贷款信息]-请求失败", zap.Any("data", fields))
		return nil, err
	}

	result := &BaseResponse[string]{}
	if err = json.Unmarshal(resp.Body(), result); err != nil {
		log.WithError(err).Error("[浦慧税贷]-[查询贷款信息]-响应数据解析未BaseResponse失败", zap.Any("data", fields))
		return nil, err
	}

	if result.Code != "000000" {
		log.WithError(err).Error("[浦慧税贷]-[查询贷款信息]-code有误", zap.Any("data", fields))
		return nil, errors.New("[浦慧税贷]-[查询贷款信息]code有误")
	}

	data := &QueryInfoResponse{}
	if err = json.Unmarshal([]byte(result.Data), data); err != nil {
		log.WithError(err).Error("[浦慧税贷]-[查询贷款信息]-响应数据解析未CreateOrderResponse失败", zap.Any("data", fields))
		return nil, err
	}

	return data, nil
}
