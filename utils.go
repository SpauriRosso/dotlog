package dotlog

// Colors code
const (
	Reset    = "\033[0m"
	White    = "\033[47m"
	BgRed    = "\033[41m"
	BgGreen  = "\033[42m"
	BgYellow = "\033[43m"
	BgBlue   = "\033[44m"
)

// Log levels
type LogLevel int

const (
	DEBU = iota
	INFO
	WARN
	ERRO
)
