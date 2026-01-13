package gerror

// 代码是通用错误代码接口定义。
type Exception interface {
	// Code返回当前错误代码的整数。
	Code() int

	// Message返回当前错误代码的简短消息。
	Message() string

	// Detail返回当前错误代码的详细信息，
	// 其主要被设计为错误代码的扩展字段。
	Detail() any

	String() string

	// 为了适配error
	Error() string

	I18n() string
}

// ================================================================================================================
// 常见错误代码定义。
// 框架保留了内部错误代码, 并且进行调整
// ================================================================================================================

var (
	CodeNil                       = localCode{code: -1, message: "", detail: nil, i18n: ""}                                      // 未指定错误代码。
	CodeOK                        = localCode{code: 200000, message: "OK", detail: nil, i18n: ""}                                // 没关系。
	CodeInternalError             = localCode{code: 200050, message: "Internal Error", detail: nil, i18n: ""}                    // 内部发生错误。
	CodeValidationFailed          = localCode{code: 200051, message: "Validation Failed", detail: nil, i18n: ""}                 // 数据验证失败。
	CodeDbOperationError          = localCode{code: 200052, message: "Database Operation Error", detail: nil, i18n: ""}          // 数据库操作错误。
	CodeInvalidParameter          = localCode{code: 200053, message: "Invalid Parameter", detail: nil, i18n: ""}                 // 当前操作的给定参数无效。
	CodeMissingParameter          = localCode{code: 200054, message: "Missing Parameter", detail: nil, i18n: ""}                 // 当前操作的参数缺失。
	CodeInvalidOperation          = localCode{code: 200055, message: "Invalid Operation", detail: nil, i18n: ""}                 // 该函数不能这样使用。
	CodeInvalidConfiguration      = localCode{code: 200056, message: "Invalid Configuration", detail: nil, i18n: ""}             // 该配置对于当前操作无效。
	CodeMissingConfiguration      = localCode{code: 200057, message: "Missing Configuration", detail: nil, i18n: ""}             // 当前操作缺少配置。
	CodeNotImplemented            = localCode{code: 200058, message: "Not Implemented", detail: nil, i18n: ""}                   // 该操作尚未实施。
	CodeNotSupported              = localCode{code: 200059, message: "Not Supported", detail: nil, i18n: ""}                     // 该操作尚不受支持。
	CodeOperationFailed           = localCode{code: 200060, message: "Operation Failed", detail: nil, i18n: ""}                  // 我试过了，但我不能给你你想要的。
	CodeNotAuthorized             = localCode{code: 200061, message: "Not Authorized", detail: nil, i18n: ""}                    // 未经授权。
	CodeSecurityReason            = localCode{code: 200062, message: "Security Reason", detail: nil, i18n: ""}                   // 安全原因。
	CodeServerBusy                = localCode{code: 200063, message: "Server Is Busy", detail: nil, i18n: ""}                    // 服务器正忙，请稍后再试。
	CodeUnknown                   = localCode{code: 200064, message: "Unknown Error", detail: nil, i18n: ""}                     // 未知错误。
	CodeNotFound                  = localCode{code: 200065, message: "Not Found", detail: nil, i18n: ""}                         // 资源不存在。
	CodeInvalidRequest            = localCode{code: 200066, message: "Invalid Request", detail: nil, i18n: ""}                   // 无效请求。
	CodeNecessaryPackageNotImport = localCode{code: 200067, message: "Necessary Package Not Import", detail: nil, i18n: ""}      // 它需要必要的包导入。
	CodeInternalPanic             = localCode{code: 200068, message: "Internal Panic", detail: nil, i18n: ""}                    // 内部发生了恐慌。
	CodeUnmarshalingPanic         = localCode{code: 200069, message: "Unmarshaling failed", detail: nil, i18n: ""}               // 反序列化失败。
	CodeRequestException          = localCode{code: 200070, message: "Request exception", detail: nil, i18n: ""}                 // 发请求失败
	CodeParameterFailure          = localCode{code: 200071, message: "Failed to receive parameters", detail: nil, i18n: ""}      // 接收参数失败
	CodeFileWriteError            = localCode{code: 200200, message: "File write failed", detail: nil, i18n: ""}                 // 文件写入失败。
	CodeFileTruncateError         = localCode{code: 200201, message: "File truncate failed", detail: nil, i18n: ""}              // 文件清空失败。
	CodeCreateDirectoryError      = localCode{code: 200202, message: "Directory create failed", detail: nil, i18n: ""}           // 文件清空失败。
	CodeGetFileSizeError          = localCode{code: 200203, message: "Get file size failed", detail: nil, i18n: ""}              // 获取文件大小失败。
	CodeFileNotExistsError        = localCode{code: 200204, message: "File not exists", detail: nil, i18n: ""}                   // 文件不存在。
	CodeFileEncodeError           = localCode{code: 200205, message: "File encode failed", detail: nil, i18n: ""}                // 文件编码不兼容。
	CodeFileCreateSysError        = localCode{code: 200206, message: "File create in sys dir failed", detail: nil, i18n: ""}     // 不能在/sys目录下创建文件。
	CodeFileCreateError           = localCode{code: 200207, message: "File create failed", detail: nil, i18n: ""}                // 创建文件失败。
	CodeDirCreateError            = localCode{code: 200208, message: "Dir create failed", detail: nil, i18n: ""}                 // 创建文件夹失败。
	CodeSensitiveError            = localCode{code: 200209, message: "Sensitive dir/file operate failed", detail: nil, i18n: ""} // 重要目录/文件操作失败。
	CodeMemoryError               = localCode{code: 200210, message: "Memory dir failed", detail: nil, i18n: ""}                 // 内存盘目录失败。
	CodeDirDeleteError            = localCode{code: 200211, message: "Dir delete failed", detail: nil, i18n: ""}                 // 删除文件夹失败。
	CodeFileDeleteError           = localCode{code: 200212, message: "File delete failed", detail: nil, i18n: ""}                // 删除文件失败。
	CodeMvSelfError               = localCode{code: 200213, message: "Mv self failed", detail: nil, i18n: ""}
	CodeRecoveryFileError         = localCode{code: 200214, message: "Recovery file failed", detail: nil, i18n: ""}          // 移动文件失败。
	CodeDelRecycleFileError       = localCode{code: 200215, message: "Del recycle file failed", detail: nil, i18n: ""}       // 删除回收站文件。
	CodeAddShortcutRootError      = localCode{code: 200216, message: "Root dir cannot add shortcuts", detail: nil, i18n: ""} // 根目录不允许创建快捷方式。
	CodeAddShortcutExistsError    = localCode{code: 200217, message: "shortcuts exists", detail: nil, i18n: ""}              // 快捷方式已存在。
	CodeAddShortcutError          = localCode{code: 200218, message: "Add shortcuts failed", detail: nil, i18n: ""}          // 创建快捷方式失败。
	CodeResponseReadError         = localCode{code: 200219, message: "Read response failed", detail: nil, i18n: ""}          // 创建快捷方式失败。

	CodeUserHomeError            = localCode{code: 200100, message: "Failed to retrieve the user's home directory", detail: nil, i18n: ""} // 获取用户主目录失败， 注意：用户相关的错误从100开始
	CodeBusinessValidationFailed = localCode{code: 200300, message: "Business Validation Failed", detail: nil, i18n: ""}                   // 业务验证失败。
	CodeBadRequest               = localCode{code: 200400, message: "Bad Request", detail: nil, i18n: ""}

	CodeTerminal = localCode{code: 200600, message: "terminal gmc", detail: nil, i18n: ""} // gmc。

	CodeCatalogueNotExist = localCode{code: 200700, message: "The catalogue does not exist", detail: nil, i18n: ""} // 200700 目录不存在。2007xx code表示

	InvalidToken = localCode{code: 200410, message: "invalid token", detail: nil, i18n: ""} // gmc。
)

// WithCodeOptionFunc defines a function type for configuring localCode options.
// These functions are used to modify localCode fields during initialization.
type WithCodeOptionFunc func(*localCode)

// New creates a new Exception (localCode) with the given code, message and detail.
// This is a direct constructor that doesn't use the options pattern.
func New(code int, message string, detail any) Exception {
	return localCode{
		code:    code,
		message: message,
		detail:  detail,
	}
}

// WithCodeFunc returns a WithCodeOptionFunc that sets the error code.
// Used with NewCode for option-based initialization.
func WithCodeFunc(code int) WithCodeOptionFunc {
	return func(lc *localCode) {
		lc.code = code
	}
}

// WithCodeMessageFunc returns a WithCodeOptionFunc that sets the error message.
// Used with NewCode for option-based initialization.
func WithCodeMessageFunc(message string) WithCodeOptionFunc {
	return func(lc *localCode) {
		lc.message = message
	}
}

// WithCodeI18nFunc returns a WithCodeOptionFunc that sets the i18n template.
// Used with NewCode for option-based initialization.
func WithCodeI18nFunc(i18n string) WithCodeOptionFunc {
	return func(lc *localCode) {
		lc.i18n = i18n
	}
}

// NewCode creates a new Exception (localCode) using the options pattern.
// Accepts one or more WithCodeOptionFunc to configure the error.
func NewCode(opts ...WithCodeOptionFunc) Exception {
	l := &localCode{}
	for _, opt := range opts {
		opt(l)
	}
	return l
}

// WithCode creates a new Exception based on an existing one with updated detail.
// Preserves the original code and message while replacing the detail.
func WithCode(code Exception, detail any) Exception {
	return localCode{
		code:    code.Code(),
		message: code.Message(),
		detail:  detail,
	}
}

// WithMessage creates a new Exception based on an existing one with updated message.
// Preserves the original code and detail while replacing the message.
func WithMessage(code Exception, message string) Exception {
	return localCode{
		code:    code.Code(),
		message: message,
		detail:  code.Detail(),
	}
}

// WithI18n creates a new Exception based on an existing one with updated i18n template.
// Preserves the original code, message and detail while adding/replacing i18n.
func WithI18n(code Exception, i18n string) Exception {
	return localCode{
		code:    code.Code(),
		message: code.Message(),
		detail:  code.Detail(),
		i18n:    i18n,
	}
}

// WithMessageErr creates a new Exception with message from an error and optional i18n.
// Uses the error's message if available, preserves code and detail from the original.
func WithMessageErr(code Exception, err error, i18n string) Exception {
	var msg string
	if err != nil {
		msg = err.Error()
	}
	return localCode{
		code:    code.Code(),
		message: msg,
		detail:  code.Detail(),
		i18n:    i18n,
	}
}

// RaiseInternalError is a convenience function for creating internal server errors.
// Combines a system error with an i18n template using CodeInternalError as base.
func RaiseInternalError(err error, i18n string) (error Exception) {
	return WithMessageErr(CodeInternalError, err, i18n)
}
