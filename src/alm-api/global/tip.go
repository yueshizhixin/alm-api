package glb

import "errors"

const (
	TIP_FAIL      = "操作失败"
	TIP_SUCCESS   = "操作成功"
	TIP_NOT_EXIST = "无该记录"
)

var (
	ERR_FAIL error
)

func init() {
	ERR_FAIL = errors.New(TIP_FAIL)
}
