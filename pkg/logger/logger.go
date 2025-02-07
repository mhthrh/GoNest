package logger

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"runtime"
	"strconv"
	"time"
)

type Message struct {
	time    time.Time
	logType Types
	Key     string
	Value   []interface{}
}
type Detail struct {
	methodName string
	start      time.Time
	end        time.Time
	Messages   []Message
}
type Log struct {
	id            uuid.UUID
	transactionId string
	timeStart     time.Time
	timeEnd       time.Time
	serviceName   string
	details       []Detail
}

func Initialize(transactionId string) Log {
	return Log{
		id:            uuid.New(),
		transactionId: transactionId,
		timeStart:     time.Now(),
		serviceName:   "",
		details:       nil,
	}
}
func (l *Log) Start() {
	if l.details != nil {
		l.details[len(l.details)-1].end = time.Now()
	}
	pc, _, _, _ := runtime.Caller(1)
	detail := Detail{
		methodName: runtime.FuncForPC(pc).Name(),
		start:      time.Now(),
		end:        time.Time{},
		Messages:   nil,
	}
	l.details = append(l.details, detail)
}

func (l *Log) Add(msg string, lType Types, values ...interface{}) {
	var r []interface{}
	for _, value := range values {
		r = append(r, value)
	}
	message := Message{
		time:    time.Now(),
		logType: lType,
		Key:     msg,
		Value:   r,
	}
	l.details[len(l.details)-1].Messages = append(l.details[len(l.details)-1].Messages, message)
}
func (l *Log) Done() {
	l.timeEnd = time.Now()
}

func (l *Log) Export() string {
	jsonByte, err := json.Marshal(&l)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(jsonByte))
	return string(jsonByte)
}

func conv(i interface{}) string {
	switch i.(type) {
	case string:
		return i.(string)
	case int:
		return strconv.Itoa(i.(int))
	case int64:
		return strconv.FormatInt(i.(int64), 10)
	case float64:
		return strconv.FormatFloat(i.(float64), 'f', -1, 64)
	case bool:
		return strconv.FormatBool(i.(bool))
	case []byte:
		return string(i.([]byte))
	case time.Time:
		return i.(time.Time).Format("2006-01-02 15:04:05")
	case error:
		return i.(error).Error()
	case time.Duration:
		return i.(time.Duration).String()
	default:
		return fmt.Sprintf("%v", i)
	}

}

func LogConfig() zapcore.Core {
	config := zapcore.EncoderConfig{
		TimeKey:       "timestamp",
		LevelKey:      "level",
		NameKey:       "logger",
		CallerKey:     "caller",
		MessageKey:    "msg",
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   zapcore.CapitalLevelEncoder,
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
		},
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	return zapcore.NewCore(
		zapcore.NewJSONEncoder(config),
		zapcore.AddSync(zapcore.Lock(os.Stdout)),
		zap.InfoLevel,
	)
}
