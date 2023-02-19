package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/fatih/color"
)

type Logger struct {
	App        string
	Module     string
	DateFormat string
	JSON       bool
}

type ProdLevelLog struct {
	App     string        `json:"app"`
	Module  string        `json:"module"`
	Message string        `json:"message"`
	Date    time.Time     `json:"date"`
	Args    []interface{} `json:"args"`
}

type colors func(format string, a ...interface{}) string

func (l *Logger) debugLevelLog(message string, args []interface{}, color colors) {
	date := time.Now()
	format := color("[%s] | [%s] - [%s] - Message: %s - Args: %s \n")
	data, _ := json.Marshal(args)
	cmd := fmt.Sprintf(format, date.Format(l.DateFormat), l.App, l.Module, message, string(data))
	io.WriteString(os.Stdout, cmd)
}

func (l *Logger) prodLevelLog(message string, args []interface{}) {
	logObject := &ProdLevelLog{
		App:     l.App,
		Module:  l.Module,
		Message: message,
		Date:    time.Now(),
		Args:    args,
	}
	out, err := json.Marshal(logObject)
	if err != nil {
		fmt.Println(err, message)
	}
	fmt.Println(string(out))
}

func (l *Logger) log(message string, args []interface{}, color colors) {
	if l.JSON == false {
		l.debugLevelLog(message, args, color)
	} else {
		l.prodLevelLog(message, args)
	}
}

func (l *Logger) Trace(message string, args ...interface{}) {
	l.log(message, args, color.CyanString)
}

func (l *Logger) Info(message string, args ...interface{}) {
	l.log(message, args, color.GreenString)
}

func (l *Logger) Warning(message string, args ...interface{}) {
	l.log(message, args, color.YellowString)
}

func (l *Logger) Error(message string, args ...interface{}) {
	l.log(message, args, color.RedString)
}

func (l *Logger) Fatal(message string, args ...interface{}) {
	l.log(message, args, color.MagentaString)
	os.Exit(1)
}

func New(
	app string,
	module string,
	dateFormat string,
	json bool,
) *Logger {
	return &Logger{
		App:        app,
		Module:     module,
		DateFormat: dateFormat,
		JSON:       json,
	}
}
