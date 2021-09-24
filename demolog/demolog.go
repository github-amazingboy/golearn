package main

import (
	"log"
	"os"

	slog "github.com/sirupsen/logrus"
)

func main() {
	logrus()
	log.Println("这是一个普通的日志")

	v := "很普通的"
	log.Printf("这是一条%v的日志", v)

	logger := log.New(os.Stdout, "amazingboy ", log.Ldate|log.Llongfile)
	logger.Printf("这是一条%v的日志", v)

	log.Fatalln("这是一条fatal的日志")
	log.Panicln("这是一条panic的日志")
	log.Println("这是一个普通的日志")
}

func init() {
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.SetPrefix("hellohello")
	logFile, err := os.OpenFile("./xx.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("open file err=%v", err)
	}
	log.SetOutput(logFile)
}

func logrus() {
	slog.SetLevel(slog.TraceLevel)
	slog.SetFormatter(&slog.JSONFormatter{})
	slog.WithFields(slog.Fields{
		"animal": "walrus",
	}).Debug("with field")
	slog.Debug("this is logrus debug")
	slog.Info("this is logrus info")
	slog.Warn("this is logrus warn")
	slog.Error("this is logrus error")
}
