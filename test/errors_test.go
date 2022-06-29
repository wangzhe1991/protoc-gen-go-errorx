package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrors(t *testing.T) {
	e := ErrorOK("操作成功")
	assert.True(t, IsOK(e))
	assert.Equal(t, 100002, BizErrorCode(e))

	e1 := OK
	assert.True(t, IsOK(e1))
	assert.Equal(t, 100002, BizErrorCode(e1))
}
