/***************************************************************************
 *
 * Copyright (c) 2020 nonstriater, Inc. All Rights Reserved
 *
 * @desc
 * @author <ranwenjie@qq.com>
 * @version 2017-10-09
 **************************************************************************/

package logger

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

type RotateState int

const (
	RotateMinute RotateState = iota
	RotateHour
	RotateDay
	RotateNone // todo 不回滚
)

func (r *RotateState) check(now, last time.Time) bool {
	switch {
	case *r == RotateMinute && now.Minute() != last.Minute():
		return true
	case *r == RotateHour && now.Hour() != last.Hour():
		return true
	case *r == RotateDay && now.Day() != last.Day():
		return true
	}
	return false
}

func (r *RotateState) suffixName(t time.Time) string {
	switch {
	case *r == RotateMinute:
		return fmt.Sprintf("%02d%02d%02d", t.Day(), t.Hour(), t.Minute())
	case *r == RotateHour:
		return fmt.Sprintf("%02d%02d", t.Day(), t.Hour())
	case *r == RotateDay:
		return fmt.Sprintf("%02d", t.Day())
	}
	return ""
}

type LogLevel int

const (
	DebugLog LogLevel = iota
	InfoLog
	WarningLog
	ErrorLog
	FatalLog
	NumsLevel = 5
)

var LogLevelName = []string{
	DebugLog:   "DEBUG",
	InfoLog:    "INFO",
	WarningLog: "WARN",
	ErrorLog:   "ERROR",
	FatalLog:   "FATAL",
}

func (s *LogLevel) getLevel() LogLevel {
	return *s
}

func (s *LogLevel) getLevelName(level LogLevel) (string, bool) {
	for i, name := range LogLevelName {
		if LogLevel(i) == level {
			return name, true
		}
	}
	return "", false
}

func (s *LogLevel) logLevelByName(v string) (int, bool) {
	v = strings.ToUpper(v)
	for i, name := range LogLevelName {
		if name == v {
			return i, true
		}
	}
	return 0, false
}

func (s *LogLevel) checkLevelValid(level LogLevel) bool {
	if level >= DebugLog && level < NumsLevel {
		return true
	}
	return false
}

type asyncBuffer struct {
	time time.Time
	buf  *bytes.Buffer
}

type AsyncLog struct {
	writer       *bufio.Writer
	logdir       string // /home/log/
	logname      string // passport
	needfilenum  bool   // add logger.go:80
	logfd        *os.File
	lasttime     time.Time
	buffer       chan *asyncBuffer
	loglevel     LogLevel
	rotatestatus RotateState
	stackdepth   int // stack depth
}

var (
	gPid     = strconv.Itoa(os.Getpid())
	gHost, _ = os.Hostname()
)

const (
	bufferSize   = 32 * 1024
	ioBufferSize = 32 * 1024
)

func NewAsyncLog(logdir, logname string, level LogLevel,
	rotate RotateState, filenumber bool, depth int) (*AsyncLog, error) {
	now := time.Now()
	l := &AsyncLog{
		logdir:       logdir,
		logname:      logname,
		needfilenum:  filenumber,
		lasttime:     time.Now(),
		buffer:       make(chan *asyncBuffer, bufferSize),
		loglevel:     level,
		rotatestatus: rotate,
		stackdepth:   depth,
	}

	fd, _, err := l.createFile(now)
	if err != nil {
		return nil, err
	}
	l.logfd = fd
	l.writer = bufio.NewWriterSize(fd, ioBufferSize)

	go l.consumerLog()
	return l, nil
}

func Debug(format string, args ...interface{}) {
	glog.write(DebugLog, []byte(fmt.Sprintf(format, args...)))
}

func Info(format string, args ...interface{}) {
	glog.write(InfoLog, []byte(fmt.Sprintf(format, args...)))
}

func Warn(format string, args ...interface{}) {
	glog.write(WarningLog, []byte(fmt.Sprintf(format, args...)))
}

func Error(format string, args ...interface{}) {
	glog.write(ErrorLog, []byte(fmt.Sprintf(format, args...)))
}

func Fatal(format string, args ...interface{}) {
	glog.write(FatalLog, []byte(fmt.Sprintf(format, args...)))
}

//-------------------------Private-------------------------------------
var bufferPool = sync.Pool{
	New: func() interface{} {
		return &asyncBuffer{buf: new(bytes.Buffer)}
	},
}

func (l *AsyncLog) write(level LogLevel, msg []byte) error {
	if !l.loglevel.checkLevelValid(level) {
		return errors.New("log: log level invalid")
	}

	if level < l.loglevel.getLevel() {
		return nil
	}

	now := time.Now()
	pool := bufferPool.Get().(*asyncBuffer)
	pool.buf.Reset()
	pool.time = now

	l.header(pool.buf, now, level, l.stackdepth)
	pool.buf.Write(msg)
	if len(msg) == 0 || msg[len(msg)-1] != '\n' {
		pool.buf.WriteByte('\n')
	}

	l.producerLog(pool)

	return nil
}

func (l *AsyncLog) consumerLog() {
	for {
		select {
		case msg := <-l.buffer:
			l.rotateFile(msg.time)
			l.writer.Write(msg.buf.Bytes()) // ignore error
			bufferPool.Put(msg)
		case <-time.After(500 * time.Microsecond):
			l.writer.Flush()
		}
	}
}

func (l *AsyncLog) producerLog(msg *asyncBuffer) {
	select {
	case l.buffer <- msg:
		return
	case <-time.After(10 * time.Millisecond):
		// only wakeup
	}
	return
}

func (l *AsyncLog) shortFile(file string) string {
	short := file
	for i := len(file) - 1; i > 0; i-- {
		if file[i] == '/' {
			short = file[i+1:]
			break
		}
	}
	return short
}

func (l *AsyncLog) header(buf *bytes.Buffer, t time.Time, level LogLevel, depth int) {
	file, line := "???", -1
	if l.needfilenum {
		var ok bool
		_, file, line, ok = runtime.Caller(2 + depth)
		if !ok {
			file = "???"
			line = -1
		} else {
			file = l.shortFile(file)
		}
	}

	l.formatHeader(buf, t, file, level, line)
}

const digits = "0123456789"

func nDigits(tmp []byte, n, i, d int, pad byte) {
	j := n - 1
	for ; j >= 0 && d > 0; j-- {
		tmp[i+j] = digits[d%10]
		d /= 10
	}
	for ; j >= 0; j-- {
		tmp[i+j] = pad
	}
}

func (l *AsyncLog) formatHeader(result *bytes.Buffer, now time.Time, file string, level LogLevel, line int) {
	//[ERROR][2016-03-11T10:37:40.003+0800][dlog/dlog_test.go:207::dlog.TestDefaultLogger]

	result.WriteByte('[')

	// level
	logname, _ := l.loglevel.getLevelName(level)
	result.WriteString(logname)
	result.WriteString("][")

	var tmp [32]byte

	year, month, day := now.Date()
	hour, minute, second := now.Clock()
	// 2017-10-15T10:15:15.142+0800
	nDigits(tmp[:], 4, 0, year, '0')
	tmp[4] = '-'
	nDigits(tmp[:], 2, 5, int(month), '0')
	tmp[7] = '-'
	nDigits(tmp[:], 2, 8, day, '0')
	tmp[10] = 'T'
	nDigits(tmp[:], 2, 11, hour, '0')
	tmp[13] = ':'
	nDigits(tmp[:], 2, 14, minute, '0')
	tmp[16] = ':'
	nDigits(tmp[:], 2, 17, second, '0')
	tmp[19] = '.'
	nDigits(tmp[:], 3, 20, now.Nanosecond()/1000000, '0')
	result.Write(tmp[:23])
	result.WriteString(now.Format("-0700"))
	result.WriteString("][")

	result.WriteString(file)
	tmp[0] = ':'
	nDigits(tmp[:], 4, 1, line, '0')
	result.Write(tmp[:5])
	result.WriteString("] ")
}

func (l *AsyncLog) logName(t time.Time) string {
	head := fmt.Sprintf("%s.log.%04d%02d",
		l.logname,
		t.Year(),
		t.Month(),
	)

	suffix := l.rotatestatus.suffixName(t)
	return head + suffix
}

func (l *AsyncLog) createFile(t time.Time) (*os.File, string, error) {
	if l.logname == "" || l.logdir == "" {
		return nil, "", errors.New("log: log name or log dir is empty")
	}

	logname := ""
	if l.logdir[len(l.logdir)-1] == '/' {
		logname = l.logdir + l.logName(t)
	} else {
		logname = l.logdir + "/" + l.logName(t)
	}

	fd, err := os.OpenFile(logname, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return nil, "", err
	}
	return fd, logname, nil
}

func (l *AsyncLog) rotateFile(t time.Time) {
	if l.logfd != nil {
		if l.rotatestatus.check(t, l.lasttime) {
			fd, _, err := l.createFile(t)
			if err != nil {
				l.exit(err)
				return
			}

			l.writer.Flush()
			l.logfd.Sync()
			l.logfd.Close()

			l.writer = bufio.NewWriterSize(fd, ioBufferSize)
			l.logfd = fd
			l.lasttime = t
		}
	}
}

func (l *AsyncLog) exit(err error) {
	fmt.Fprintf(os.Stderr, "log: exiting because of error: %s\n", err)
	os.Exit(2)
}

var glog *AsyncLog

func InitLog(dir, name string) {
	var err error
	glog, err = NewAsyncLog(dir, name, DebugLog, RotateHour, true, 1)
	if err != nil {
		panic(fmt.Sprintf("init log: path %v err %v", dir+name, err))
	}
}

func InitCustomLog(dir, name string, depth int) {
	if !dirIsExist(dir) {
		os.Mkdir(dir, 0775)
	}
	var err error
	glog, err = NewAsyncLog(dir, name, DebugLog, RotateHour, true, depth)
	if err != nil {
		panic(fmt.Sprintf("init log: path %v err %v", dir+name, err))
	}
}

func dirIsExist(dir string) bool {
	fir, err := os.Stat(dir)
	if err != nil {
		if os.IsExist(err) {
			return fir.IsDir()
		}
		return false
	}
	return fir.IsDir()
}
