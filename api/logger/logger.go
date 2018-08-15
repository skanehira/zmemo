package logger

import (
	"log"

	"github.com/comail/colog"
)

const (
	Linfo  = "info: %"
	Lwarn  = "warn: %"
	Lerror = "error: %"
	Ldebug = "debug: %"
)

func Init() {
	// ログの設定方法
	// log.Printf("trace: this is a trace log.")
	// log.Printf("debug: this is a debug log.")
	// log.Printf("info: this is an info log.")
	// log.Printf("warn: this is a warning log.")
	// log.Printf("error: this is an error log.")
	// log.Printf("alert: this is an alert log.")
	// log.Printf("this is a default level log.")

	// ログ設定
	// colog.SetDefaultLevel(colog.LDebug)
	colog.SetMinLevel(colog.LTrace)
	colog.SetFormatter(&colog.StdFormatter{
		Colors: true,
		Flag:   log.Ldate | log.Ltime | log.Lshortfile,
	})
	colog.Register()

}

func LogInfo(message string) {
	log.Print("info: " + message)
}

func LogInfof(format string, args ...interface{}) {
	log.Printf(format, args)
}
