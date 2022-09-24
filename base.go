package pyd2

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	json "github.com/json-iterator/go"

	"github.com/zrb-channel/utils"
	"github.com/zrb-channel/utils/aesutil"
	"github.com/zrb-channel/utils/hash"
	log "github.com/zrb-channel/utils/logger"
	"github.com/zrb-channel/utils/rsautil"
	"net/url"
	"strconv"
	"strings"
	"time"
)

var (
	AesIV    = []byte{0x38, 0x37, 0x36, 0x35, 0x34, 0x33, 0x32, 0x31, 0x68, 0x67, 0x66, 0x65, 0x64, 0x63, 0x62, 0x61}
	SubAesIV = []byte("hanasian12345678")
)

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

// NewServiceRequest
// @param conf
// @param serviceID
// @param id
// @param msg
// @date 2022-09-24 01:06:03
func NewServiceRequest(conf *Config, serviceID string, id string, msg interface{}) (*ServiceBaseRequest, error) {
	base := &ServiceBaseRequest{
		AppId:     conf.SubConfig.AppId,
		RequestId: id,
		Timestamp: strconv.FormatInt(time.Now().UnixNano()/1e6, 10),
		Channel:   conf.SubConfig.Channel,
	}

	base.SetServiceID(serviceID)

	if err := base.Sign(conf, msg); err != nil {
		log.WithError(err).Error("消息签名失败")
		return nil, err
	}

	return base, nil
}

// SetAk
// @param v
// @date 2022-09-24 01:06:01
func (req *ServiceBaseRequest) SetAk(v string) {
	req.Ak = v
}

// SetServiceID
// @param id
// @date 2022-09-24 01:06:00
func (req *ServiceBaseRequest) SetServiceID(id string) {
	req.ServiceId = id
}

// SetSignture
// @param v
// @date 2022-09-24 01:05:59
func (req *ServiceBaseRequest) SetSignture(v string) {
	req.Signture = v
}

// SetMessage
// @param message
// @date 2022-09-24 01:05:58
func (req *ServiceBaseRequest) SetMessage(message string) {
	req.Message = message
}

// Sign
// @param conf
// @param msg
// @date 2022-09-24 01:05:57
func (req *ServiceBaseRequest) Sign(conf *Config, msg interface{}) error {
	publicKey, err := PublicKeyFrom64(conf.SubConfig.PublicKey)
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
	encryptMessage, err = aesutil.EncryptToBase64(message, []byte(aesKey), SubAesIV)
	if err != nil {
		return err
	}

	req.SetAk(ak)

	sign := hash.SHA256String(encryptMessage + conf.SubConfig.AppSecret + req.Timestamp + req.ServiceId + conf.SubConfig.Channel)

	encryptMessage = url.QueryEscape(formatString(encryptMessage))
	req.SetMessage(encryptMessage)
	req.SetSignture(strings.ToUpper(sign))
	return nil
}

// formatString
// @param sourceStr
// @date 2022-09-24 01:05:56
func formatString(sourceStr string) string {
	if sourceStr == "" {
		return ""
	}
	return strings.ReplaceAll(strings.ReplaceAll(sourceStr, "\\r", ""), "\\n", "")
}

// PublicKeyFrom64
// @param key
// @date 2022-09-24 01:05:55
func PublicKeyFrom64(key string) (*rsa.PublicKey, error) {
	b, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return nil, err
	}
	return PublicKeyFrom(b)
}

// PublicKeyFrom
// @param key
// @date 2022-09-24 01:05:54
func PublicKeyFrom(key []byte) (*rsa.PublicKey, error) {
	if pub, err := x509.ParsePKIXPublicKey(key); err != nil {
		return nil, err
	} else {
		return pub.(*rsa.PublicKey), nil
	}
}
