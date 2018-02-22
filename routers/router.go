package routers

import (
	"Email/controllers"
	"github.com/astaxie/beego"
)

func init() {
		beego.Router("/", &controllers.MainController{},"get:Email")
		beego.Router("/uplaod_user", &controllers.MainController{},"post:Uplaoduser")
		beego.Router("/sendmail", &controllers.MainController{},"post:SendMail")


}