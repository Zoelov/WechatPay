package main

import (
	"fmt"

	"WechatPay/env"

	"github.com/astaxie/beego"
)

func init() {
	env.RunEnvInit()
}

func main() {
	appID := beego.AppConfig.String("app_id")
	fmt.Println(appID)
}
