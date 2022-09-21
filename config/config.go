package config

/**
const (
	Addr      = "https://smeb-stg1.jryzt.com/portal/portal/open/%s"
	AppSecret = "zssDauyhrPeUHELyUYE5Sg=="
	publicKey = "MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDHHJ0BL6pC2dOse2LZpn3i7gs1ibaAJLHOTlQ+UX0uK2+iRHt5XJLC/wFo4GyJhbJ1Hj4B+GMgBkj6fhG8o8sTMtg1YhP6Dn2iEKlyIlI4bFPov9jG+hsNtg3w2iWntMPqr2CNZ/LSazOF6lVtWtWYuJJvdWTSJSmLuouWMrf6/wIDAQAB"
	AppCode   = "FSQB_SYS"
)


const (
	Addr      = "https://e.jryzt.com/portal/portal/open/%s"
	AppSecret = "O5FLEkdkZaIQnB36TCzPfg=="
	publicKey = "MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCNaZDwtYsF1MK2fHi3QjK+Q5gEJjeFoRNc83mR1VcYreyHd8qu2Q+LnVvbG3h+ewTNkTq0NJWrRKJuqlIJkYk8XjMSJ27anW2zPVgxMg7T435cDIirD9yIHWaTmrYpOGWPzlLBnh83NWXC6Aj/OEJyUqZ4ooJktE3EuvROT1UpSwIDAQAB"
	AppCode   = "FSQB_SYS"
)

var AesIV = []byte{0x38, 0x37, 0x36, 0x35, 0x34, 0x33, 0x32, 0x31, 0x68, 0x67, 0x66, 0x65, 0x64, 0x63, 0x62, 0x61}

var PublicKey *rsa.PublicKey

func init() {
	var err error
	if PublicKey, err = PublicKeyFrom64(publicKey); err != nil {
		log.WithError(err).Fatal("加载公钥失败")
	}
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

*/
