package logger_test

import (
	"github.com/mhthrh/GoNest/pkg/logger"
	"testing"
)

var (
	transactionId = "ISMS000000123987"
	log           logger.Log
)

func init() {
	log = logger.Initialize(transactionId)
}
func TestLog_Start(t *testing.T) {
	log.Start()
}

func TestLog_Add_1(t *testing.T) {
	log.Start()
	log.Add("message-1", logger.Types(0), "value", 12)
	log.Add("message-1", logger.Types(2), "value", 12, []string{"1", "2", "3"})
}

func TestLog_Add_2(t *testing.T) {
	log.Start()
	log.Add("message-1", logger.Types(0), "value", 12)
	log.Add("message-1", logger.Types(2), "value", 12, []string{"1", "2", "3"})
}
func TestLog_Done(t *testing.T) {
	log.Done()
}

func TestLog_Export(t *testing.T) {
	log.Export()
}
