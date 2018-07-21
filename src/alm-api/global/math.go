package glb

import (
	"strings"
	"github.com/satori/go.uuid"
)

/*
	数学库
 */
 
//?:
func If(condition bool, trueExpression, falseExpression interface{}) interface{} {
	if condition {
		return trueExpression
	}
	return falseExpression
}

func UUID() string {
	return strings.Replace(uuid.Must(uuid.NewV4()).String(), "-", "", 4)
}
