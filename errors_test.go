package errors

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 这里定义错误码
var (
	ConfigurationNotValid = 1
)

func Test_Example(t *testing.T) {
	Register(DefaultCoder{ConfigurationNotValid, 500, "ConfigurationNotValid error"})

	err := WithCode(ConfigurationNotValid, "test")
	// err -> 错误码 -> Coder -> 错误码对应的信息
	coder := ParseCoder(err)
	assert.Equal(t, coder.Code(), ConfigurationNotValid)
	assert.Equal(t, coder.String(), "ConfigurationNotValid error")
}
