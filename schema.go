package pyd2

import (
	"github.com/shopspring/decimal"
	"github.com/zrb-channel/utils"
)

type (
	BaseRequest struct {
		Data            string `json:"bizContent"`
		AppCode         string `json:"appCode"`
		Timestamp       string `json:"timestamp"`
		ServiceCode     string `json:"serviceCode"`
		Sign            string `json:"sign"`
		EncryptedAesKey string `json:"encryptedAesKey"`
	}

	BaseResponse[T any] struct {
		Data    T      `json:"responseData"`
		Message string `json:"responseMessage"`
		Code    string `json:"responseCode"`
	}

	NotifyRequest struct {
		BizContent string `json:"bizContent"`
		Sign       string `json:"sign"`
		Timestamp  int64  `json:"timestamp"`
		RequestId  string `json:"requestId"`
	}

	QueryInfoRequest struct {
		OrderNo string `json:"extApplicationCode"`
		Channel string `json:"channelAgent"`
		TaxCode string `json:"taxCode"`
		BillNo  string `json:"billNo"`
	}

	UploadInvoiceRequest struct {
		OrderNo    string `json:"extApplicationCode"` // 渠道流水号(外部订单号)
		Channel    string `json:"channelAgent"`       // 渠道号
		TaxName    string `json:"companyName"`        // 公司名称
		TaxNo      string `json:"companyCreditCode"`  // 统一社会信用代码
		TaxAddress string `json:"companyAddress"`
	}

	UploadInvoiceResponse struct {
		LoanOrderCode string `json:"loanOrderCode"`
	}

	LoginRequest struct {
		RequestNo string `json:"requestNo"`
		YztCode   string `json:"yztCode"`   // 渠道编号,壹账通分配
		ClientID  string `json:"clientId"`  // 客户端ID
		TaxName   string `json:"taxName"`   // 公司名称
		TaxNo     string `json:"taxCode"`   // 统一社会信用代码
		OperName  string `json:"operName"`  // 法人
		OperNo    string `json:"operNo"`    // 法人证件号码
		ReturnUrl string `json:"returnUrl"` // 当selectType=1时必传:H5采集完成的回跳地 址址
	}

	LoginResponse struct {
		URL    string `json:"URL"`
		TaskId string `json:"taskId"`
	}

	PluginGatherCheckRequest struct {
		QueryType         string `json:"queryType"`         // 00:RPA前置申请 01:原标准产品 ---V1.0.1新增
		TranNo            string `json:"tranNo"`            // 订单流水号
		TaxDishId         string `json:"taxDishId"`         // 纳税人识别号
		ChannelId         string `json:"channelId"`         // 渠道编号
		CompanyCreditCode string `json:"companyCreditCode"` // 统一社会信用代码
		CompanyName       string `json:"companyName"`       // 公司名称
		InformStartDate   string `json:"informStartDate"`   // 查询开始时间
		InformEndDate     string `json:"informEndDate"`     // 查询结束时间
	}

	PluginGatherCheckResponse struct {
		Records         string `json:"records"`
		ResponseMessage string `json:"responseMessage"`
		ResponseCode    string `json:"responseCode"`
	}

	CollectModelQueryRequest struct {
		YztCode string `json:"yztCode"` // 渠道编号
		TaxCode string `json:"taxCode"` // 企业税号
	}

	CollectModelQueryResponse struct {
		Records         string `json:"records"`
		ResponseMessage string `json:"responseMessage"`
		ResponseCode    string `json:"responseCode"`
	}

	H5QueryRequest struct {
		TaskId            string  `json:"taskId"` // H5页面采集发票时的任务id
		YztCode           string  `json:"yztCode"`
		TaxDishId         *string `json:"taxDishId"`         // 纳税人识别号
		CompanyCreditCode string  `json:"companyCreditCode"` // 企业统一社会
		CompanyName       string  `json:"companyName"`       // 公司名称
		TranNo            string  `json:"tranNo"`            // 此入参是在弘昊项目引入， SMEB|2350为必填
		QueryType         string  `json:"queryType"`
	}

	H5QueryResponseResult struct {
		Records         string `json:"records"`
		ResponseMessage string `json:"responseMessage"`
		ResponseCode    string `json:"responseCode"`
	}
	H5QueryResponse struct {
		Result *H5QueryResponseResult `json:"item"`
		TaskID string                 `json:"taskId"`
	}

	PluginQueryRequest struct {
		RelationType       string `json:"relationType"`       // 关联类型 1:壹企数
		ChannelID          string `json:"channelId"`          // 渠道编号
		TaxNo              string `json:"taxDishId"`          // 纳税人识别号
		CompanyCreditCode  string `json:"companyCreditCode"`  // 企业统一社会信用代码
		CompanyName        string `json:"companyName"`        // 公司名称
		CompanyOwner       string `json:"companyOwner"`       // 企业法人
		CompanyOwnerIdCard string `json:"companyOwnerIdCard"` // 企业法人身份证件号
		CompanyOwnerMobile string `json:"companyOwnerMobile"` // 企业法人手机号码
		InformStartDate    string `json:"informStartDate"`    // 查询采集通知开始时间，格式：yyyy-MM-dd
		InformEndDate      string `json:"informEndDate"`      // 查询采集通知结束时间，格式：yyyy-MM-dd
		TranNo             string `json:"tranNo"`             // 订单流水号
	}

	NotifyPreApply struct {
		SupId              string            `json:"supId"`
		CertType           string            `json:"certType"`
		ApprStatus         string            `json:"apprStatus"`
		CustMobile         string            `json:"custMobile"`
		CertCode           string            `json:"certCode"`
		ExtApplicationCode string            `json:"extApplicationCode"`
		PlatTime           string            `json:"platTime"`
		CustName           string            `json:"custName"`
		RetCode            string            `json:"retCode"`
		PlatDate           string            `json:"platDate"`
		ApplyNoPre         string            `json:"applyNoPre"`
		LoanOrderCode      string            `json:"loanOrderCode"`
		ApprAmt            utils.TextDecimal `json:"apprAmt"`
	}

	NotifyResponse[T any] struct {
		Message string `json:"responseMessage"`
		Code    string `json:"responseCode"`
		Data    T      `json:"responseData"`
	}

	NotifyCreditRequest struct {
		CreditEndDate      string            `json:"creditEndDate"`
		SupId              string            `json:"supId"`
		RepayBankNo        string            `json:"repayBankNo"`
		ApprStatus         string            `json:"apprStatus"`
		CustMobile         string            `json:"custMobile"`
		ExtApplicationCode string            `json:"extApplicationCode"`
		SignStatus         string            `json:"signStatus"`
		RetCode            string            `json:"retCode"`
		TaxName            string            `json:"taxName"`
		PlatDate           string            `json:"platDate"`
		ApplyNoCredit      string            `json:"applyNoCredit"`
		CreditId           string            `json:"creditId"`
		CreditTerm         string            `json:"creditTerm"`
		LoanOrderCode      string            `json:"loanOrderCode"`
		RepayCardNo        string            `json:"repayCardNo"`
		PlatTime           string            `json:"platTime"`
		TaxCode            string            `json:"taxCode"`
		CustName           string            `json:"custName"`
		ApprAmt            decimal.Decimal   `json:"apprAmt"`
		ApplyAmt           decimal.Decimal   `json:"applyAmt"`
		CreditAmt          utils.TextDecimal `json:"creditAmt"`
		SignAmt            string            `json:"signAmt"`
		CreditStartDate    string            `json:"creditStartDate"`
	}

	NotifyApplyRequest struct {
		RepayCardNo        string          `json:"repayCardNo"`
		SupId              string          `json:"supId"`
		CertType           string          `json:"certType"`
		RepayBankNo        string          `json:"repayBankNo"`
		ApprStatus         string          `json:"apprStatus"`
		CustMobile         string          `json:"custMobile"`
		CertCode           string          `json:"certCode"`
		ExtApplicationCode string          `json:"extApplicationCode"`
		SignStatus         string          `json:"signStatus"`
		PlatTime           string          `json:"platTime"`
		CustName           string          `json:"custName"`
		RetCode            string          `json:"retCode"`
		ApprAmt            decimal.Decimal `json:"apprAmt"`
		ApplyAmt           decimal.Decimal `json:"applyAmt"`
		PlatDate           string          `json:"platDate"`
		CreditId           string          `json:"creditId"`
		ApplyNo            string          `json:"applyNo"`
		LoanOrderCode      string          `json:"loanOrderCode"`
		SignAmt            string          `json:"signAmt"`
		BillNo             string          `json:"billNo"`
	}

	CreateOrderRequest struct {
		OrderNo     string `json:"extApplicationCode"` // 外部订单号
		Channel     string `json:"channelAgent"`       // 渠道号
		Mobile      string `json:"mobile"`             // 申请人手机号码
		UserName    string `json:"userName"`           // 申请人姓名
		RequestID   string `json:"req"`                // 请求ID
		SupId       string `json:"supId"`              // 纳税人识别号
		SupPlatCode string `json:"supPlatCode"`        // 经营平台编号
		TaxCode     string `json:"taxCode"`            // 纳税人识别号
		TaxName     string `json:"taxName"`            // 企业名称
		CustName    string `json:"custName"`           // 企业代表人姓名
		CertType    string `json:"certType"`           // 法定代表人证件类型
		CertCode    string `json:"certCode"`           // 法定代表人证件号码
		CustMobile  string `json:"custMobile"`         // 法定代表人手机号码
		CustLevel   string `json:"custLevel"`          // 客户级别
		RegDate     string `json:"regDate"`            // 注册日期
		EnterType   string `json:"enterType"`          // 企业类型
		AreaType    string `json:"areaType"`           // 地区类别
		AreaCode    string `json:"areaCode"`           // 行政区划
		IndCode     string `json:"indCode"`            // 行业代码
		BizAddr     string `json:"bizAddr"`            // 经营地址
		RegAddr     string `json:"regAddr"`            // 注册地址
		LossFlag    string `json:"lossFlag"`           // 是否有失信被执行人信息
		LossAmt     string `json:"lossAmt"`            // 被执行人信息中执行标的（元）
		PunishAmt   string `json:"punishAmt"`          // 行政处罚历史- 处罚金额
		PledgeAmt   string `json:"pledgeAmt"`          // 股权出质历史- 出质金额
		LicType     string `json:"licType"`            // 经营者身份类型
		LicNo       string `json:"licNo"`              // 营业执照号
	}

	CreateOrderResponse struct {
		LoanOrderCode string `json:"loanOrderCode"`
	}

	QueryInfoResponse struct {
		CertType      string          `json:"certType"`
		LoanStartDate string          `json:"loanStartDate"`
		LoanEndDate   string          `json:"loanEndDate"`
		LoanAmt       decimal.Decimal `json:"loanAmt"`
		ExecRate      decimal.Decimal `json:"execRate"`
	}

	RedirectRequest struct {
		ExtApplicationCode string `json:"extApplicationCode"` // 渠道流水号(外部订单号)
		ChannelAgent       string `json:"channelAgent"`       // 渠道号
		SupId              string `json:"supId"`              // 用户ID
		SupPlatCode        string `json:"supPlatCode"`        // 经营平台编号
		RemoteUrl          string `json:"remoteUrl"`          // 渠道方回调地址	渠道方回调地址
	}

	JumpResponse struct {
		JumpUrl string `json:"jumpUrl"`
	}

	ServiceBaseRequest struct {
		ServiceId string `json:"serviceId"` // 服务ID
		AppId     string `json:"appId"`     // 应用ID    文本文件内容获取（key: appId）
		RequestId string `json:"requestId"` // 请求ID    长度最长为64位，数字和字符串的组合，客户端请求唯一标识（建议用UUID）, 客户端生成，每次接口请求的requestId都是不同的
		Timestamp string `json:"timestamp"` // 时间戳    long毫秒数（当前时间）, 也用于加签
		Channel   string `json:"channel"`   // 渠道    应用所属渠道
		Signture  string `json:"signture"`  // 加签内容    视接口需要，若接口需要加签则必填，若接口不需要加签则不需要。
		Ak        string `json:"ak"`        // Base64编码后的加密的AES秘钥    视接口需要，若接口需要加密则必填，若接口不需要加密则不需要
		Message   string `json:"message"`   // 业务参数    AES秘钥加密的业务入参，业务参数为appId的值
	}

	ServiceBaseResponse[T any] struct {
		Data    T      `json:"responseData"`
		Message string `json:"responseMessage"`
		Code    string `json:"responseCode"`
	}

	BaseData[T any] struct {
		Records   T      `json:"records"`
		RequestNo string `json:"requestNo"`
		Message   string `json:"responseMessage"`
		Code      string `json:"responseCode"`
	}

	ResponseRecord[T any] struct {
		Msg  string `json:"msg"`
		Code int    `json:"code"`
		Data T      `json:"data"`
	}
)
