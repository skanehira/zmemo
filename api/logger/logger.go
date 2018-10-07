package logger

import (
	"log"

	"github.com/comail/colog"
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

func Info(message string) {
	log.Print("info: " + message)
}

func Infof(format string, args ...interface{}) {
	log.Printf("info: "+format, args)
}

func Warn(err error) {
	format := "warn: %+v"
	log.Printf(format, err)
}

func Error(err error) {
	format := "error: %+v"
	log.Printf(format, err)
}

func Errorf(format string, i interface{}) {
	if err, ok := i.(error); ok {
		format = "error: %+v"
		log.Printf(format, err)
	} else {
		format = "error: " + format
		log.Printf(format, i)
	}
}
