package errors

import (
	err "github.com/geektime007/errors"
)

var (
	// 通用错误
	Nil             *err.Error = nil
	Unknown                    = err.NewError(err.ErrorCodeUnknown, 500, "未知错误")
	BadFormat                  = err.NewError(err.ErrorCodeBadFormat, 400, "格式错误")
	RunCommandError            = err.NewError(err.ErrorRunCommandError, 500, "执行命令异常")
	Timeout                    = err.NewError(err.ErrorCodeTimeout, 408, "超时")
	Unimplemented              = err.NewError(err.ErrorCodeUnimplemented, 404, "未实现!")
)
