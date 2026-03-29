package logger

import (
	"fmt"
	"log"
	"os"
)

type Logger struct {
	Info *log.Logger
	Error *log.Logger
	Debug *log.Logger
}

var AppLogger *Logger

func init(){
	AppLogger = loggerFunction()
}

func loggerFunction() *Logger{

	file , err := os.OpenFile("./logger/applog.log",os.O_CREATE|os.O_WRONLY|os.O_APPEND,0666)
	if(err != nil){
		fmt.Print("Couldn't Create/Write/Append log file : ",err)
		os.Exit(1)
	}

	infoLog := log.New(file,"INFO : ",log.Ldate|log.Ltime|log.Lshortfile)
	errorLog := log.New(file,"ERROR : ",log.Ldate|log.Ltime|log.Lshortfile)
	debugLog := log.New(file,"DEBUG : ",log.Ldate|log.Ltime|log.Lshortfile)

	
	return &Logger{
		Info: infoLog,
		Error: errorLog,
		Debug: debugLog,
	}
}