package routers

import (
	"rateserver/controllers"

	beego "github.com/beego/beego/v2/adapter"
)

func init() {
	beego.Router("/api/v3/ticker/24hr", &controllers.MainController{}, "get:GetBinancePrice")
	beego.Router("/api/v3/depth", &controllers.MainController{}, "get:GetBinanceDepth")
	beego.Router("/api/v3/klines", &controllers.MainController{}, "get:GetBinanceCandleStick")
	beego.Router("/api/mexc/v3/ticker/24hr", &controllers.MainController{}, "get:GetMexcPrice")
	beego.Router("/api/mexc/v3/depth", &controllers.MainController{}, "get:GetMexcDepth")
	beego.Router("/api/mexc/v3/klines", &controllers.MainController{}, "get:GetMexcCandleStick")
}
