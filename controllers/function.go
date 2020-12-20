package controllers

import "github.com/astaxie/beego"

type FunctionController struct {
	beego.Controller
}

const KEY = "RaW#XhH2aVgo!Iy1"

type JsonStruct struct {
	Code  int         `json:"code"`
	Msg   interface{} `json:"msg"`
	Items interface{} `json:"items"`
	Count int64       `json:"count"`
}

type JsonEmptyStruct struct {
	Code  int         `json:"code"`
	Items interface{} `json:"items"`
	Count int64       `json:"count"`
}

func ReturnError(code int, msg interface{}) (json *JsonStruct) {
	json = &JsonStruct{Code: code, Msg: msg}
	return
}

func ReturnSuccess(code int, msg interface{}, item interface{}, count int64) (json *JsonStruct) {
	json = &JsonStruct{Code: code, Msg: msg, Items: item, Count: count}
	return
}

func ReturnEmpty(code int, item interface{}, count int64) (json *JsonEmptyStruct) {
	json = &JsonEmptyStruct{Code: code, Items: item, Count: count}
	return
}
