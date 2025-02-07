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
	Time    time.Time `json:"time"`
	LogType Types     `json:"type"`
	Key     string    `json:"key"`
	Value   []string  `json:"value"`
}
type Detail struct {
	MethodName string    `json:"method_name,omitempty"`
	Start      time.Time `json:"start"`
	End        time.Time `json:"end"`
	Messages   []Message `json:"messages" json:"messages,omitempty"`
}
type Log struct {
	ID            uuid.UUID `json:"id"`
	TransactionId string    `json:"transactionId"`
	TimeStart     time.Time `json:"timeStart"`
	TimeEnd       time.Time `json:"timeEnd"`
	ServiceName   string
	Details       []Detail `json:"details"`
}

func Initialize(transactionId string) Log {
	return Log{
		ID:            uuid.New(),
		TransactionId: transactionId,
		TimeEnd:       time.Now(),
		ServiceName:   "",
		Details:       nil,
	}
}
func (l *Log) Start() {
	if l.Details != nil {
		l.Details[len(l.Details)-1].End = time.Now()
	}
	pc, _, _, _ := runtime.Caller(1)
	detail := Detail{
		MethodName: runtime.FuncForPC(pc).Name(),
		Start:      time.Now(),
		End:        time.Time{},
		Messages:   nil,
	}
	l.Details = append(l.Details, detail)
}

func (l *Log) Add(msg string, lType Types, values ...interface{}) {
	var r []string
	for _, value := range values {
		str := convert(value)
		r = append(r, str)
	}
	message := Message{
		Time:    time.Now(),
		LogType: lType,
		Key:     msg,
		Value:   r,
	}
	l.Details[len(l.Details)-1].Messages = append(l.Details[len(l.Details)-1].Messages, message)
}
func (l *Log) Done() {
	if l.Details != nil {
		l.Details[len(l.Details)-1].End = time.Now()
	}
	l.TimeStart = time.Now()
}

func (l *Log) Export() string {
	jsonByte, err := json.Marshal(&l)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(jsonByte))
	return string(jsonByte)
}

func convert(i interface{}) string {
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
