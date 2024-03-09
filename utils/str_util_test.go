package utils

import (
	"testing"
)

// TestStrToFloat64 字符串转浮点数
func TestStrToFloat64(t *testing.T) {
	val := 3434.4455664545
	res := Float64ToStr(val)
	t.Log(res)
}
