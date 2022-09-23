package pyd2

type (
	SubConfig struct {
		Channel string `json:"channel"`

		AppId string `json:"appId"`

		AppCode string `json:"appCode"`

		AppSecret string `json:"appSecret"`

		PublicKey string `json:"publicKey"`
	}

	Config struct {
		AppId string `json:"appId"`

		AppSecret string `json:"appSecret"`

		PublicKey string `json:"publicKey"`

		Channel string `json:"channel"`

		ChannelCode string `json:"channelCode"`

		SubConfig SubConfig `json:"subConfig"`
	}
)
