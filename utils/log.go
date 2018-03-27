package utils

import (
	"os"
	"log"

)

var (
	Log	*log.Logger
)

// NewLog creates a file on the specified path and writes all logs there
func NewLog(logpath string){
	println("Log file: " + logpath)

	file, err := os.Create(logpath)
	if err != nil{
		panic(err)
	}
	Log = log.New(file, "", log.LstdFlags|log.Lshortfile)
}