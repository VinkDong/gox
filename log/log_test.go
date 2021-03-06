package log

import (
	"testing"
	"time"
)

func TestInfo(t *testing.T) {
	Info("this is info")
}

func TestError(t *testing.T) {
	Error("this is error")
}

func TestSuccess(t *testing.T) {
	Success("this is success")
}

func TestSuccessf(t *testing.T) {
	Successf("this is success with args type: %s", "string")
}

func TestWarn(t *testing.T) {
	Warn("this is warn")
}

func TestWarnf(t *testing.T) {
	Warnf("this is warn with args type: %s", "string")
}

func TestLock(t *testing.T) {
	Info("this is first log next lock log")
	Lock()
	go func() {
		for {
			Info("this is sub task")
			time.Sleep(time.Second * 1)
		}
	}()
	time.Sleep(time.Second * 1)
	Unlock()
	Info("this is last log")
	Lock()
	time.Sleep(time.Second * 3)
	Unlock()
	time.Sleep(time.Second * 5)
}

func TestWrite(t *testing.T) {
	writeEnabled = true
	SetFilename("/tmp/gox.log")
	Write("hahaha")
	Write(123)
	Write('a')
}

func TestErrorLine(t *testing.T) {
	ErrorLine("this is log with file info")
}
