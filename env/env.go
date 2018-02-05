package env

import "github.com/astaxie/beego"

// RunEnvInit 环境
func RunEnvInit() {
	err := beego.LoadAppConfig("ini", "../config/app.conf")
	if err != nil {
		panic(err)
	}
}
