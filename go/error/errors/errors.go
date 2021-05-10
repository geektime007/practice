package errors

import (
	err "github.com/geektime007/errors"
)

var (
	// 通用错误
	Nil             *err.Error = nil
	Unknown                    = err.NewError(1000, 500, "未知错误")
	BadFormat                  = err.NewError(1001, 400, "格式错误")
	RunCommandError            = err.NewError(1002, 500, "执行命令异常")
	Timeout                    = err.NewError(1003, 408, "超时")
	Unimplemented              = err.NewError(1004, 404, "未实现!")
)
