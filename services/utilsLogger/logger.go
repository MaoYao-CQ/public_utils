package utilsLogger

import (
	"fmt"
	"gopkg.in/natefinch/lumberjack.v2"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

var (
	LoggerTrace   *log.Logger // 记录所有日志
	LoggerInfo    *log.Logger // 重要的信息
	LoggerWarning *log.Logger // 需要注意的信息
	LoggerError   *log.Logger // 非常严重的问题
	LogPrefix    string      // 日志前辍
)

func GetSelfPath() string {
	root := filepath.Dir(os.Args[0])
	//fmt.Println(root)

	// 根据相对路径获取可执行文件的绝对路径
	root, err := filepath.Abs(root)
	if err != nil {
		log.Fatalln(err)
	}

	return root
}

func Init() {
	prefix := LogPrefix
	if "" != prefix {
		prefix += ": "
	}
	logWriter := &lumberjack.Logger{
		Filename:   filepath.Join(GetSelfPath(), "logs", "log.txt"),
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     28,   //days
		Compress:   true, // disabled by default
	}

	LoggerTrace = log.New(ioutil.Discard,
		prefix+"TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	LoggerInfo = log.New(logWriter,
		prefix+"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	LoggerWarning = log.New(logWriter,
		prefix+"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	LoggerError = log.New(logWriter,
		prefix+"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

func LogInfo(message string, err error) {
	fmt.Println(message, err)
	LoggerInfo.Println(message, err)
}

func LogWrong(message string, err error) {
	fmt.Println(message, err)
	LoggerWarning.Println(message, err)
}

func Logger(message string, err error) {
	fmt.Println(message, err)
	LoggerError.Println(message, err)
}
func LogTrace(message string, err error) {
	fmt.Println(message, err)
	LoggerTrace.Println(message, err)
}

func LoggerExit(message string, err error) {
	Logger(message, err)
	panic("stop on error")
}
