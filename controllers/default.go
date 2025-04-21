package controllers

import (
	"fmt"
	"net/http"
	"rateserver/logpack"
	"rateserver/services"
)

type MainController struct {
	BaseController
}

var BinanceAPI = "https://api.binance.com/api/v3/"
var MexcAPI = "https://api.mexc.com/api/v3/"

func (this *MainController) GetBinancePrice() {
	var symbol string
	if err := this.Ctx.Input.Bind(&symbol, "symbol"); err != nil {
		logpack.Error(err.Error(), "GetBinancePrice")
		this.Data["json"] = map[string]string{"status": "false", "errMsg": "Get Symbol failed"}
		this.ServeJSON()
		return
	}
	query := map[string]string{
		"symbol": symbol,
	}
	url := fmt.Sprintf("%sticker/24hr", BinanceAPI)
	req := &services.ReqConfig{
		Method:  http.MethodGet,
		HttpUrl: url,
		Payload: query,
	}
	var result interface{}
	if err := services.HttpRequest(req, &result); err != nil {
		this.Data["json"] = map[string]string{"status": "false", "errMsg": "Get Binance Price data failed"}
		this.ServeJSON()
		return
	}
	this.Data["json"] = result
	this.ServeJSON()
}

func (this *MainController) GetBinanceDepth() {
	var symbol, limit string
	symbolErr := this.Ctx.Input.Bind(&symbol, "symbol")
	limitErr := this.Ctx.Input.Bind(&limit, "limit")
	if symbolErr != nil || limitErr != nil {
		logpack.Error("Get param error", "GetBinanceDepth")
		this.Data["json"] = map[string]string{"status": "false", "errMsg": "Get Param failed"}
		this.ServeJSON()
		return
	}
	query := map[string]string{
		"symbol": symbol,
		"limit":  limit,
	}
	url := fmt.Sprintf("%sdepth", BinanceAPI)
	req := &services.ReqConfig{
		Method:  http.MethodGet,
		HttpUrl: url,
		Payload: query,
	}
	var result interface{}
	if err := services.HttpRequest(req, &result); err != nil {
		logpack.Error("Get Binance Depth error", "GetBinanceDepth")
		this.Data["json"] = map[string]string{"status": "false", "errMsg": "Get Binance Depth data failed"}
		this.ServeJSON()
		return
	}
	this.Data["json"] = result
	this.ServeJSON()
}

func (this *MainController) GetBinanceCandleStick() {
	var symbol, interval string
	symbolErr := this.Ctx.Input.Bind(&symbol, "symbol")
	intervalErr := this.Ctx.Input.Bind(&interval, "interval")
	if symbolErr != nil || intervalErr != nil {
		logpack.Error("Get param error", "GetBinanceCandleStick")
		this.Data["json"] = map[string]string{"status": "false", "errMsg": "Get Param failed"}
		this.ServeJSON()
		return
	}
	query := map[string]string{
		"symbol":   symbol,
		"interval": interval,
	}
	url := fmt.Sprintf("%sklines", BinanceAPI)
	req := &services.ReqConfig{
		Method:  http.MethodGet,
		HttpUrl: url,
		Payload: query,
	}
	var result interface{}
	if err := services.HttpRequest(req, &result); err != nil {
		logpack.Error("Get Binance Candle Stick error", "GetBinanceCandleStick")
		this.Data["json"] = map[string]string{"status": "false", "errMsg": "Get Binance Candle Stick data failed"}
		this.ServeJSON()
		return
	}
	this.Data["json"] = result
	this.ServeJSON()
}

func (this *MainController) GetMexcPrice() {
	var symbol string
	if err := this.Ctx.Input.Bind(&symbol, "symbol"); err != nil {
		logpack.Error(err.Error(), "GetMexcPrice")
		this.Data["json"] = map[string]string{"status": "false", "errMsg": "Get Symbol failed"}
		this.ServeJSON()
		return
	}
	query := map[string]string{
		"symbol": symbol,
	}
	url := fmt.Sprintf("%sticker/24hr", MexcAPI)
	req := &services.ReqConfig{
		Method:  http.MethodGet,
		HttpUrl: url,
		Payload: query,
	}
	var result interface{}
	if err := services.HttpRequest(req, &result); err != nil {
		this.Data["json"] = map[string]string{"status": "false", "errMsg": "Get Mexc Price data failed"}
		this.ServeJSON()
		return
	}
	this.Data["json"] = result
	this.ServeJSON()
}

func (this *MainController) GetMexcDepth() {
	var symbol, limit string
	symbolErr := this.Ctx.Input.Bind(&symbol, "symbol")
	limitErr := this.Ctx.Input.Bind(&limit, "limit")
	if symbolErr != nil || limitErr != nil {
		logpack.Error("Get param error", "GetMexcDepth")
		this.Data["json"] = map[string]string{"status": "false", "errMsg": "Get Param failed"}
		this.ServeJSON()
		return
	}
	query := map[string]string{
		"symbol": symbol,
		"limit":  limit,
	}
	url := fmt.Sprintf("%sdepth", MexcAPI)
	req := &services.ReqConfig{
		Method:  http.MethodGet,
		HttpUrl: url,
		Payload: query,
	}
	var result interface{}
	if err := services.HttpRequest(req, &result); err != nil {
		logpack.Error("Get Mexc Depth error", "GetMexcDepth")
		this.Data["json"] = map[string]string{"status": "false", "errMsg": "Get Mexc Depth data failed"}
		this.ServeJSON()
		return
	}
	this.Data["json"] = result
	this.ServeJSON()
}

func (this *MainController) GetMexcCandleStick() {
	var symbol, interval string
	symbolErr := this.Ctx.Input.Bind(&symbol, "symbol")
	intervalErr := this.Ctx.Input.Bind(&interval, "interval")
	if symbolErr != nil || intervalErr != nil {
		logpack.Error("Get param error", "GetMexcCandleStick")
		this.Data["json"] = map[string]string{"status": "false", "errMsg": "Get Param failed"}
		this.ServeJSON()
		return
	}
	query := map[string]string{
		"symbol":   symbol,
		"interval": interval,
	}
	url := fmt.Sprintf("%sklines", MexcAPI)
	req := &services.ReqConfig{
		Method:  http.MethodGet,
		HttpUrl: url,
		Payload: query,
	}
	var result interface{}
	if err := services.HttpRequest(req, &result); err != nil {
		logpack.Error("Get Mexc Candle Stick error", "GetMexcCandleStick")
		this.Data["json"] = map[string]string{"status": "false", "errMsg": "Get Mexc Candle Stick data failed"}
		this.ServeJSON()
		return
	}
	this.Data["json"] = result
	this.ServeJSON()
}
