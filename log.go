package dotlog

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

//func CheckErr(infoMsg string, f interface{}, okMsg string) {
//	Log(INFO, infoMsg)
//	var err error
//
//	switch fn := f.(type) {
//	case func() error:
//		err = fn()
//	case func() (interface{}, error):
//		_, err = fn()
//	default:
//		Log(ERRO, "Unsupported function type")
//		os.Exit(1)
//	}
//
//	if err != nil {
//		_, file, line, _ := runtime.Caller(1)
//		errMsg := fmt.Sprintf("%s:%d - %v", file, line, err)
//		Log(ERRO, errMsg)
//		os.Exit(1)
//	} else {
//		Log(INFO, okMsg)
//	}
//}

func Debug(msg string) {
	Log(DEBU, msg)
}

func Info(msg string) {
	Log(INFO, msg)
}

func Warning(msg string) {
	Log(WARN, msg)
}

func Error(msg string) {
	Log(ERRO, msg)
}

func Log(level LogLevel, msg string) {
	timeNow := time.Now().Format("2006-02-01 - 15:04:05")
	color := fmt.Sprintf("%s%s%s%s", level.Color(), White, level, Reset)
	fmt.Printf("%s %s %s\n", timeNow, color, msg)
}

// String() is a method implementing Stringer which is useful to:
// - Provide a custom string representation of the type
// - Allow automatic string conversion in fmt.Print functions
func (l LogLevel) String() string {
	// [...] is a syntax that lets the compiler decide of the length of an array based on the number of elements
	return [...]string{"DEBU", "INFO", "WARN", "ERRO"}[l]
}

func (l LogLevel) Color() string {
	return [...]string{BgBlue, BgGreen, BgYellow, BgRed}[l]
}

func CheckFuncErr(infoMsg string, f interface{}, okMsg string) {
	Log(INFO, infoMsg)
	var err error

	switch fn := f.(type) {
	case func() error:
		err = fn()
	case func() (interface{}, error):
		_, err = fn()
	default:
		Log(ERRO, "Unsupported function type")
		os.Exit(1)
	}

	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		errMsg := fmt.Sprintf("%s:%d - %v", file, line, err)
		Log(ERRO, errMsg)
		os.Exit(1)
	} else {
		Log(INFO, okMsg)
	}
}

func CheckServer(infoMsg string, f func() error, okMsg string) {
	Log(INFO, infoMsg)

	go func() {
		err := f()
		if err != nil {
			_, file, line, _ := runtime.Caller(1)
			errMsg := fmt.Sprintf("%s:%d - %v", file, line, err)
			Log(ERRO, errMsg)
			//os.Exit(1)
		}
	}()
	time.Sleep(100 * time.Millisecond)
	Log(INFO, okMsg)
}
