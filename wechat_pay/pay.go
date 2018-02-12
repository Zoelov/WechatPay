package wechat_pay

type WechatPay interface {
	AppID() string
	SecretID() string
	MerchantID() string
	APIKey() string
	PayType() int
	RunMod() int
}

type payParam struct {
	appID          string `xml:"appid"`
	mchID          string `xml:"mch_id"`
	deviceInfo     string `xml:"device_info,omitempty"`
	nonceStr       string `xml:"nonce_str"`
	sign           string `xml:"sign"`
	body           string `xml:"body"`
	detail         string `xml:"detail,omitempty"`
	attach         string `xml:"attach,omitempty"`
	outTradeNo     string `xml:"out_trade_no"`
	totalFee       int    `xml:"total_fee"`
	spbillCreateIP string `xml:"spbill_create_ip"`
	goodsTag       string `xml:"goods_tag,omitempty"`
	limitPay       string `xml:"limit_pay,omitempty"`
	authCode       string `xml:"auth_code"`
	sceneInfo      string `xml:"scene_info,omitempty"`
}

func CommitPayment(instance WechatPay) {

}
