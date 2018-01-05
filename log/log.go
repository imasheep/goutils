// log.go
package log

import (
	"errors"
	"fmt"
	g "goutils"
	"os"
	"path/filepath"
	"time"

	"github.com/virsheep/stdlog"
)

var logLevelMap = map[int]string{
	1: "Info",
	2: "Warn",
	3: "Error",
	4: "Fatal",
}

var logPool = make(map[string]*LogInstance)
var LogPool *map[string]*LogInstance

func New(flag string) (l *LogInstance) {
	l = (*LogPool)[flag]
	return
}

func (this *LogInstance) log(level int, content ...interface{}) (err error) {

	this.LogPath = g.If(this.LogPath == "", "/tmp", this.LogPath).(string)
	this.Flag = g.If(this.Flag == "", "sniper", this.Flag).(string)
	this.TimeFlag = g.If(this.TimeFlag == "", "20060102", this.TimeFlag).(string)

	logNameTS := time.Now().Format(this.TimeFlag)

	fs, err := os.OpenFile(
		filepath.Join(this.LogPath, this.Flag+"."+logNameTS+".log"),
		os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	defer fs.Close()
	if err != nil {
		return
	}

	loger := stdlog.New(fs, "", stdlog.LstdFlags)
	loger.Println("["+logLevelMap[level]+"]", content)

	return
}

type LogInstance struct {
	LogPath  string
	Flag     string
	TimeFlag string
}

func (this *LogInstance) Info(content ...interface{}) (err error) {

	err = this.log(1, fmt.Sprint(content...))
	if err != nil {
		return
	}

	return
}

func (this *LogInstance) Warn(content ...interface{}) (err error) {

	err = this.log(2, fmt.Sprint(content...))
	if err != nil {
		return
	}

	return
}

func (this *LogInstance) Error(content ...interface{}) (err error) {

	err = this.log(3, fmt.Sprint(content...))
	if err != nil {
		return
	}

	return
}

func (this *LogInstance) Fatal(content ...interface{}) (err error) {

	err = this.log(4, fmt.Sprint(content...))
	if err != nil {
		return
	}

	return
}

func regLogInstance(logInstance LogInstance) (err error) {

	if logInstance.Flag == "" {
		return errors.New("logInstance.flag can not be empty")
	}

	logPool[logInstance.Flag] = &logInstance

	return
}

func regLogInstanceMulti(logInstances []LogInstance) (err error) {

	for _, logInstance := range logInstances {

		err = regLogInstance(logInstance)
		if err != nil {
			return
		}

	}

	LogPool = &logPool

	return
}
func LogInit(logInstances []LogInstance) {
	regLogInstanceMulti(logInstances)
}
