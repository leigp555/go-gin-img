package utils

type ErrCode map[int]string

var ErrMap = ErrCode{
	1000: "系统错误",
	1001: "无权访问",
	1002: "服务器繁忙",
	1003: "内容不存在",
	1004: "身份过期",
	1005: "用户或密码不正确",
}
