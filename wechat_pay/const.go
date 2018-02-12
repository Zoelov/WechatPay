package wechat_pay

const (
	// PayTypeCreditCard 刷卡支付
	PayTypeCreditCard = iota + 1

	// PayTypeScanQR 扫码支付
	PayTypeScanQR

	// PayTypePublicAccount 公众号支付
	PayTypePublicAccount

	// PayTypeApp app支付
	PayTypeApp
)

const (
	// DevMod 开发模式
	DevMod = iota
	// ProMod 线上模式
	ProMod
)

const (
	URLCreditCardPaying = "https://api.mch.weixin.qq.com/pay/micropay"
)
