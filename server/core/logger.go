package core

import (
	"log"
	"os"
)

type Logger struct {
	warnings    *log.Logger
	information *log.Logger
	errors      *log.Logger
	fatal       *log.Logger
}

func NewLogger() *Logger {
	file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	logger := &Logger{}
	logger.warnings = log.New(file, "[WARN] ", log.Ldate|log.Ltime)
	logger.information = log.New(file, "[INFO] ", log.Ldate|log.Ltime)
	logger.errors = log.New(file, "[ERROR] ", log.Ldate|log.Ltime)
	logger.fatal = log.New(file, "[FATAL] ", log.Ldate|log.Ltime)

	return logger
}

func (l *Logger) Warn(msg string) {
	l.warnings.Println(msg)
	log.Println(l.warnings.Prefix() + " " + msg)
}

func (l *Logger) Info(msg string) {
	l.information.Println(msg)
	log.Println(l.information.Prefix() + " " + msg)
}

func (l *Logger) Error(msg string) {
	l.errors.Println(msg)
	log.Println(l.errors.Prefix() + " " + msg)
}

func (l *Logger) Fatal(msg string) {
	l.fatal.Fatalln(msg)
	log.Println(l.fatal.Prefix() + " " + msg)
}
