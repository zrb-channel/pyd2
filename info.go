package pyd2

import (
	"context"
	"fmt"
	"time"

	"github.com/zrb-channel/utils"

	log "github.com/zrb-channel/utils/logger"

	json "github.com/json-iterator/go"

	"go.uber.org/zap"
)

func QueryInfo(ctx context.Context, conf *Config, req *QueryInfoRequest) {

	body, err := NewRequest(conf, "S000705", req)

	fields := map[string]any{
		"req":  req,
		"body": body,
	}

	if err != nil {
		log.WithError(err).Error("[浦慧税贷]-[查询贷款信息]-创建请求失败", zap.Any("data", fields))
		return
	}

	addr := fmt.Sprintf(Addr, "S000705")

	resp, err := utils.Request(ctx).SetHeaders(headers).SetBody(body).Post(addr)
	if err != nil {
		log.WithError(err).Error("[浦慧税贷]-[查询贷款信息]-请求失败", zap.Any("data", fields))
		return
	}

	result := &BaseResponse[string]{}
	if err = json.Unmarshal(resp.Body(), result); err != nil {
		log.WithError(err).Error("[浦慧税贷]-[查询贷款信息]-响应数据解析未BaseResponse失败", zap.Any("data", fields))
		return
	}

	if result.Code != "000000" {
		log.WithError(err).Error("[浦慧税贷]-[查询贷款信息]-code有误", zap.Any("data", fields))
		return
	}

	data := &QueryInfoResponse{}
	if err = json.Unmarshal([]byte(result.Data), data); err != nil {
		log.WithError(err).Error("[浦慧税贷]-[查询贷款信息]-响应数据解析未CreateOrderResponse失败", zap.Any("data", fields))
		return
	}

	if data.CertType != "1" {
		fmt.Println("还款方式有误")
		return
	}

	start, err := time.ParseInLocation("20060102", data.LoanStartDate, time.Local)
	end, err := time.ParseInLocation("20060102", data.LoanEndDate, time.Local)

	// 贷款利息=贷款金额×贷款时间×贷款利率
	subMonth := SubMonth(end, start)
	fmt.Println(subMonth)

}
