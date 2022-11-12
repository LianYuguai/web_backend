package consts

const (
	CODE_OK       = "0"
	CODE_AUTH     = "0000001"
	CODE_PARAMS   = "0000002"
	CODE_INTERNAL = "0000003"
)

var CODE_MSG = map[string]string{
	CODE_OK:       "",
	CODE_AUTH:     "登录信息失效",
	CODE_PARAMS:   "参数错误",
	CODE_INTERNAL: "服务器内部错误",
}
