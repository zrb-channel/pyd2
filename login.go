package pyd2

import (
	"context"
	"errors"
	"net/http"

	"github.com/zrb-channel/utils"
	log "github.com/zrb-channel/utils/logger"

	"github.com/go-resty/resty/v2"
	json "github.com/json-iterator/go"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
)

// TaxLogin 生成H5页面地址链接
// @param ctx
// @param order *model.MemberApply
// @date 2022-05-17 18:13:56
func TaxLogin(ctx context.Context, conf *Config, req *LoginRequest) (string, error) {
	if err := ctx.Err(); err != nil {
		return "", err
	}

	id := uuid.NewV4().String()

	body, err := NewServiceRequest(conf, "tax090302", id, req)
	fields := map[string]any{
		"req":  req,
		"body": body,
	}

	if err != nil {
		log.WithError(err).Error("[浦慧税贷]-[生成H5页面地址链接]-创建请求失败", zap.Any("data", fields))
		return "", errors.New(err.Error())
	}

	var resp *resty.Response
	resp, err = utils.Request(ctx).SetBody(body).SetHeaders(headers).Post(ServiceAddr + "/jd/tax/taxLogin")

	if err != nil {
		log.WithError(err).Error("[浦慧税贷]-[生成H5页面地址链接]-请求失败", zap.Any("data", fields))
		return "", errors.New(err.Error())
	}

	if resp.StatusCode() != http.StatusOK {
		fields["status"] = resp.StatusCode()
		log.WithError(err).Error("[浦慧税贷]-[生成H5页面地址链接]-响应状态码有误", zap.Any("data", fields))
		return "", errors.New("请求失败")
	}

	fields["resp"] = resp.String()
	result := &BaseResponse[*BaseData[*ResponseRecord[*LoginResponse]]]{}
	if err = json.Unmarshal(resp.Body(), result); err != nil {
		log.WithError(err).Error("[浦慧税贷]-[生成H5页面地址链接]-响应数据解析为BaseResponse失败", zap.Any("data", fields))
		return "", err
	}

	fields["result"] = result
	if result.Code != "000000" {
		log.WithError(err).Error("[浦慧税贷]-[生成H5页面地址链接]-code有误", zap.Any("data", fields))
		return "", errors.New(result.Message)
	}

	log.Info("[浦慧税贷]-[生成H5页面地址链接]-成功", zap.Any("data", fields))
	return result.Data.Records.Data.URL, nil
}

// PluginGatherCheck 判断是否有必要走RPA
// @param ctx context.Context
// @param order *model.MemberApply
// @date 2022-05-17 18:13:55
func PluginGatherCheck(ctx context.Context, conf *Config, req *PluginGatherCheckRequest) (*BaseResponse[*PluginGatherCheckResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	id := uuid.NewV4().String()

	body, err := NewServiceRequest(conf, "RPACJ090301", id, req)
	fields := map[string]any{
		"req":  req,
		"body": body,
	}

	if err != nil {
		log.WithError(err).Error("[浦慧税贷]-[判断是否有必要走RPA]-创建请求失败", zap.Any("data", fields))
		return nil, err
	}

	resp, err := utils.Request(ctx).SetBody(body).SetHeaders(headers).Post(ServiceAddr + "/rpa/upgrade/fpPluginGatherCheck")

	if err != nil {
		log.WithError(err).Error("[浦慧税贷]-[判断是否有必要走RPA]-请求失败", zap.Any("data", fields))
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		fields["status"] = resp.StatusCode()
		log.WithError(err).Error("[浦慧税贷]-[判断是否有必要走RPA]-响应状态码有误", zap.Any("data", fields))
		return nil, errors.New(resp.Status())
	}

	fields["resp"] = resp.String()
	result := &BaseResponse[*PluginGatherCheckResponse]{}
	if err = json.Unmarshal(resp.Body(), result); err != nil {
		log.WithError(err).Error("[浦慧税贷]-[判断是否有必要走RPA]-响应数据解析为BaseResponse失败", zap.Any("data", fields))
		return nil, err
	}

	fields["result"] = result

	log.WithError(err).Info("[浦慧税贷]-[判断是否有必要走RPA]-成功", zap.Any("data", fields))
	return result, nil
}

// CollectModelQuery 查询是否支持RPA采集
// @param ctx context.Context
// @param taxCode string
// @date 2022-05-17 18:13:52
func CollectModelQuery(ctx context.Context, conf *Config, req *CollectModelQueryRequest) (*BaseResponse[*CollectModelQueryResponse], error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	id := uuid.NewV4().String()

	body, err := NewServiceRequest(conf, "RPA082502", id, req)
	fields := map[string]any{
		"req":  req,
		"body": body,
	}

	if err != nil {
		log.WithError(err).Error("[浦慧税贷]-[查询是否支持RPA采集]-创建请求失败", zap.Any("data", fields))
		return nil, err
	}

	resp, err := utils.Request(ctx).SetBody(body).SetHeaders(headers).Post(ServiceAddr + "/collect/model/query")
	if err != nil {
		log.WithError(err).Error("[浦慧税贷]-[查询是否支持RPA采集]-请求失败", zap.Any("data", fields))
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		fields["status"] = resp.StatusCode()
		log.WithError(err).Error("[浦慧税贷]-[查询是否支持RPA采集]-响应状态码有误", zap.Any("data", fields))
		return nil, errors.New(resp.Status())
	}

	fields["resp"] = resp.String()

	result := &BaseResponse[*CollectModelQueryResponse]{}
	if err = json.Unmarshal(resp.Body(), result); err != nil {
		log.WithError(err).Error("[浦慧税贷]-[查询是否支持RPA采集]-响应数据解析为BaseResponse失败", zap.Any("data", fields))
		return nil, err
	}

	fields["result"] = result

	log.WithError(err).Info("[浦慧税贷]-[查询是否支持RPA采集]-成功", zap.Any("data", fields))

	return result, nil
}

// CollectRpaQuery H5发票关联查询
// @param ctx context.Context
// @param order *model.MemberApply
// @date 2022-05-17 18:13:51
func CollectRpaQuery(ctx context.Context, conf *Config, req *H5QueryRequest) (*BaseResponse[*H5QueryResponseResult], error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	id := uuid.NewV4().String()

	body, err := NewServiceRequest(conf, "RPA082501", id, req)

	fields := map[string]any{
		"req":  req,
		"body": body,
	}

	if err != nil {
		log.WithError(err).Error("[浦慧税贷]-[H5发票关联查询]-创建请求失败", zap.Any("data", fields))
		return nil, err
	}

	resp, err := utils.Request(ctx).SetBody(body).SetHeaders(headers).Post(ServiceAddr + "/longma/billHaierbillCorrelation/RPA")

	if err != nil {
		log.WithError(err).Error("[浦慧税贷]-[H5发票关联查询]-请求失败", zap.Any("data", fields))
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		fields["status"] = resp.StatusCode()
		log.WithError(err).Error("[浦慧税贷]-[H5发票关联查询]-响应状态码有误", zap.Any("data", fields))
		return nil, errors.New(resp.Status())
	}

	fields["resp"] = resp.String()
	result := &BaseResponse[*H5QueryResponseResult]{}
	if err = json.Unmarshal(resp.Body(), result); err != nil {
		log.WithError(err).Error("[浦慧税贷]-[H5发票关联查询]-响应数据解析为BaseResponse失败", zap.Any("data", fields))
		return nil, err
	}

	fields["result"] = result
	log.WithError(err).Info("[浦慧税贷]-[H5发票关联查询]-成功", zap.Any("data", fields))
	return result, nil
}

// CollectPluginQuery 插件发票关联查询
// @param ctx
// @param order *model.MemberApply
// @date 2022-05-17 23:49:26
func CollectPluginQuery(ctx context.Context, conf *Config, req *PluginQueryRequest) (*BaseResponse[any], error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	id := uuid.NewV4().String()

	body, err := NewServiceRequest(conf, "HRFPD083101", id, req)

	fields := map[string]any{"req": req, "body": body}

	if err != nil {
		log.WithError(err).Error("[浦慧税贷]-[插件发票关联查询]-创建请求失败", zap.Any("data", fields))
		return nil, err
	}

	resp, err := utils.Request(ctx).SetBody(body).SetHeaders(headers).Post(ServiceAddr + "/longma/billHaier/billCorrelation")

	if err != nil {
		log.WithError(err).Error("[浦慧税贷]-[插件发票关联查询]-请求失败", zap.Any("data", fields))
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		fields["status"] = resp.StatusCode()
		log.WithError(err).Error("[浦慧税贷]-[插件发票关联查询]-响应状态码有误", zap.Any("data", fields))
		return nil, errors.New(resp.Status())
	}

	fields["resp"] = resp.String()
	result := &BaseResponse[any]{}
	if err = json.Unmarshal(resp.Body(), result); err != nil {
		log.WithError(err).Error("[浦慧税贷]-[插件发票关联查询]-响应数据解析为BaseResponse失败", zap.Any("data", fields))
		return nil, err
	}

	fields["result"] = result

	log.WithError(err).Info("[浦慧税贷]-[插件发票关联查询]-成功", zap.Any("data", fields))
	return result, nil
}
