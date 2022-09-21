package pyd2

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"net/url"
	"strconv"
	"strings"
	"time"

	log "github.com/zrb-channel/utils/logger"

	"github.com/zrb-channel/utils/aesutil"

	json "github.com/json-iterator/go"
	"github.com/zrb-channel/utils"
	"github.com/zrb-channel/utils/hash"
	"github.com/zrb-channel/utils/rsautil"
)

var AesIV = []byte{0x38, 0x37, 0x36, 0x35, 0x34, 0x33, 0x32, 0x31, 0x68, 0x67, 0x66, 0x65, 0x64, 0x63, 0x62, 0x61}

// NewRequest
// @param serviceCode
// @param data
// @date 2022-09-21 17:59:12
func NewRequest(conf *Config, serviceCode string, data any) (*BaseRequest, error) {

	value, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	aesKey := []byte(utils.RandString(16))
	var enData string
	if enData, err = aesutil.EncryptToBase64(value, aesKey, AesIV); err != nil {
		return nil, err
	}

	publicKey, err := PublicKeyFrom64(conf.PublicKey)
	if err != nil {
		return nil, err
	}

	var enKeyBytes []byte
	if enKeyBytes, err = rsautil.PublicEncrypt(publicKey, aesKey); err != nil {
		return nil, err
	}

	timestamp := strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
	sign := hash.SHA256String(enData + conf.AppId + serviceCode + timestamp + conf.AppSecret)

	base := &BaseRequest{
		EncryptedAesKey: base64.StdEncoding.EncodeToString(enKeyBytes),
		Data:            enData,
		ServiceCode:     serviceCode,
		AppCode:         conf.AppId,
		Sign:            sign,
		Timestamp:       timestamp,
	}
	return base, nil
}

// SubMonth
// @param t1
// @param t2
// @date 2022-09-21 18:08:22
func SubMonth(t1, t2 time.Time) (month int) {
	y1 := t1.Year()
	y2 := t2.Year()
	m1 := int(t1.Month())
	m2 := int(t2.Month())
	d1 := t1.Day()
	d2 := t2.Day()

	yearInterval := y1 - y2
	// 如果 d1的 月-日 小于 d2的 月-日 那么 yearInterval-- 这样就得到了相差的年数
	if m1 < m2 || m1 == m2 && d1 < d2 {
		yearInterval--
	}
	// 获取月数差值
	monthInterval := (m1 + 12) - m2
	if d1 < d2 {
		monthInterval--
	}
	monthInterval %= 12
	month = yearInterval*12 + monthInterval
	return
}

func NewServiceRequest(conf *Config, serviceID string, id string, msg interface{}) (*ServiceBaseRequest, error) {
	base := &ServiceBaseRequest{
		AppId:     conf.AppId,
		RequestId: id,
		Timestamp: strconv.FormatInt(time.Now().UnixNano()/1e6, 10),
		Channel:   conf.Channel,
	}

	base.SetServiceID(serviceID)

	if err := base.Sign(conf, msg); err != nil {
		log.WithError(err).Error("消息签名失败")
		return nil, err
	}

	return base, nil
}

func (req *ServiceBaseRequest) SetAk(v string) {
	req.Ak = v
}

func (req *ServiceBaseRequest) SetServiceID(id string) {
	req.ServiceId = id
}

func (req *ServiceBaseRequest) SetSignture(v string) {
	req.Signture = v
}

func (req *ServiceBaseRequest) SetMessage(message string) {
	req.Message = message
}

func (req *ServiceBaseRequest) Sign(conf *Config, msg interface{}) error {
	publicKey, err := PublicKeyFrom64(conf.PublicKey)
	if err != nil {
		return err
	}

	aesKey := utils.RandString(16)

	ak, err := rsautil.PublicEncryptToBase64(publicKey, []byte(aesKey))
	if err != nil {
		return err
	}

	message, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	var encryptMessage string
	encryptMessage, err = aesutil.EncryptToBase64(message, []byte(aesKey), AesIV)
	if err != nil {
		return err
	}

	req.SetAk(ak)

	sign := hash.SHA256String(encryptMessage + conf.AppSecret + req.Timestamp + req.ServiceId + conf.Channel)

	encryptMessage = url.QueryEscape(formatString(encryptMessage))
	req.SetMessage(encryptMessage)
	req.SetSignture(strings.ToUpper(sign))
	return nil
}

func formatString(sourceStr string) string {
	if sourceStr == "" {
		return ""
	}
	return strings.ReplaceAll(strings.ReplaceAll(sourceStr, "\\r", ""), "\\n", "")
}

func PublicKeyFrom64(key string) (*rsa.PublicKey, error) {
	b, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return nil, err
	}
	return PublicKeyFrom(b)
}

func PublicKeyFrom(key []byte) (*rsa.PublicKey, error) {
	if pub, err := x509.ParsePKIXPublicKey(key); err != nil {
		return nil, err
	} else {
		return pub.(*rsa.PublicKey), nil
	}
}
