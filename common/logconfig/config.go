package logconfig

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"runtime"
	"strings"
)

func InitialLogConfig(opts ...Option) {
	//default
	config := &Config{
		reportCaller:    true,
		level:           logrus.InfoLevel,
		timestampFormat: "2006-01-02 15:04:05.000",
		levelDesc:       []string{"PANIC", "FATAL", "ERROR", "WARN", "INFO", "DEBUG"},
		project:         "unknown-project",
		module:          "unknown-module",
		version:         "unknown-version",
		env:             "unknown-env",
		debug:           false,
	}
	for _, opt := range opts {
		opt(config)
	}
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(config.level)
	logrus.SetReportCaller(config.reportCaller)
	logrus.SetFormatter(&CustomizeFormatter{
		TimestampFormat: "2006-01-02 15:04:05.000",
		LevelDesc:       []string{"PANIC", "FATAL", "ERROR", "WARN", "INFO", "DEBUG"},
		Project:         config.project,
		Module:          config.module,
		Version:         config.version,
		Env:             config.env,
		Debug:           config.debug,
	})
}

type CustomizeFormatter struct {
	TimestampFormat string
	LevelDesc       []string
	Project         string
	Module          string
	Version         string
	Env             string
	Debug           bool
}

func (cf *CustomizeFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	timestamp := fmt.Sprintf(entry.Time.Format(cf.TimestampFormat))
	if cf.Debug {
		return []byte(fmt.Sprintf("[%s] %s\n", timestamp, entry.Message)), nil
	}
	return []byte(fmt.Sprintf("[%s] [%s] [%s] [%s] [%s] [%s] [%s] [%s] - %s\n", timestamp, cf.Project, cf.Module, cf.Env, cf.Version, goroutineId(), cf.LevelDesc[entry.Level], findCaller(entry.Caller), entry.Message)), nil
}

func goroutineId() string {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	goroutineArray := strings.Fields(string(buf[:n]))
	if goroutineArray == nil || len(goroutineArray) < 2 {
		return "goroutine-999"
	}
	return fmt.Sprintf("%s-%s", strings.Fields(string(buf[:n]))[0], strings.Fields(string(buf[:n]))[1])
}

func findCaller(f *runtime.Frame) string {
	if f != nil {
		filename := path.Base(f.File)
		funcArray := strings.Split(f.Function, ".")
		if funcArray == nil || len(funcArray) == 1 {
			return fmt.Sprintf("%s/%s():%d", filename, f.Function, f.Line)
		}
		funcName := funcArray[len(funcArray)-1]
		return fmt.Sprintf("%s/%s():%d", filename, funcName, f.Line)
	}
	result := ""
	for i := 1; i <= 50; i++ {
		if pc, file, line, ok := runtime.Caller(i); ok {
			funcName := runtime.FuncForPC(pc).Name()
			result = fmt.Sprintf("%s:%s:%d", path.Base(file), path.Base(funcName), line)
			if !strings.Contains(funcName, "logrus") && !strings.Contains(funcName, "logrus-client") {
				break
			}
		}
	}
	return result
}
