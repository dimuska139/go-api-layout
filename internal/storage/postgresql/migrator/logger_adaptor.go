package migrator

import (
	"context"
	"fmt"

	"github.com/dimuska139/urlshortener/pkg/logging"
)

type LoggerAdaptor struct {
}

func NewLoggerAdaptor() *LoggerAdaptor {
	return &LoggerAdaptor{}
}

func (l *LoggerAdaptor) Fatal(v ...any) {
	logging.Fatal(context.Background(), "", v...)
}

func (l *LoggerAdaptor) Fatalf(format string, v ...any) {
	logging.Fatal(context.Background(), fmt.Sprintf(format, v...))
}

func (l *LoggerAdaptor) Print(v ...any) {
	logging.Fatal(context.Background(), "", v...)
}

func (l *LoggerAdaptor) Println(v ...any) {
	logging.Info(context.Background(), "", v...)
}

func (l *LoggerAdaptor) Printf(format string, v ...any) {
	logging.Info(context.Background(), fmt.Sprintf(format, v...))
}
