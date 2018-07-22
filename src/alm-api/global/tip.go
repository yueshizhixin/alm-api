package glb

import "errors"

const (
	TIP_FAIL    = "操作失败"
	TIP_SUCCESS = "操作成功"
)

var (
	ERR_FAIL = errors.New(TIP_FAIL)
)
