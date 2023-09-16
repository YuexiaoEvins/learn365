package lib

type RetCode int

const (
	RetCodeSuccess            RetCode = 0    // 成功。
	RetCodeWarningCallTimeout         = 1001 // 调用超时警告。
	RetCodeErrorCall                  = 2001 // 调用错误。
	RetCodeErrorResponse              = 2002 // 响应内容错误。
	RetCodeErrorCalee                 = 2003 // 被调用方（被测软件）的内部错误。
	RetCodeFatalCall                  = 3001 // 调用过程中发生了致命错误！
)

func GetRetCodePlain(code RetCode) string {
	var codePlain string
	switch code {
	case RetCodeSuccess:
		codePlain = "Success"
	case RetCodeWarningCallTimeout:
		codePlain = "Call Timeout Warning"
	case RetCodeErrorCall:
		codePlain = "Call Error"
	case RetCodeErrorResponse:
		codePlain = "Response Error"
	case RetCodeErrorCalee:
		codePlain = "Callee Error"
	case RetCodeFatalCall:
		codePlain = "Call Fatal Error"
	default:
		codePlain = "Unknown result code"
	}
	return codePlain
}

const (
	// STATUS_ORIGINAL 代表原始。
	STATUS_ORIGINAL uint32 = 0
	// STATUS_STARTING 代表正在启动。
	STATUS_STARTING = 1
	// STATUS_STARTED 代表已启动。
	STATUS_STARTED = 2
	// STATUS_STOPPING 代表正在停止。
	STATUS_STOPPING = 3
	// STATUS_STOPPED 代表已停止。
	STATUS_STOPPED = 4
)
