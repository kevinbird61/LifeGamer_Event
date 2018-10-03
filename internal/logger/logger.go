package logger

import (
	"log"
	"io"
	"os"
)

// class: LG_Logger
type LG_Logger struct {
	Info		*log.Logger 
	Warning 	*log.Logger 
	Error		*log.Logger

	infoURL		string 
	warnURL		string 
	errURL		string
}

// Initial
func (lg *LG_Logger) Init(){
	lg.errURL = os.TempDir() + "/lg.e.errors.log"
	lg.infoURL= os.TempDir() + "/lg.e.info.log"
	lg.warnURL= os.TempDir() +"/lg.e.warn.log"

	errFile, err := os.OpenFile(lg.errURL, os.O_CREATE | os.O_WRONLY | os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Open errors.log Failed : ", err)
	}
	infoFile, err := os.OpenFile(lg.infoURL, os.O_CREATE | os.O_WRONLY | os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Open info.log Failed : ", err)
	}
	warnFile, err := os.OpenFile(lg.warnURL, os.O_CREATE | os.O_WRONLY | os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Open warn.log Failed : ", err)
	}

	lg.Info = log.New(io.MultiWriter(os.Stdout, infoFile), "[LifeGamer Event Engine][Info]", log.Ldate | log.Ltime | log.Lshortfile)
	lg.Error = log.New(io.MultiWriter(os.Stderr, errFile), "[LifeGamer Event Engine][Error]", log.Ldate | log.Ltime | log.Lshortfile)
	lg.Warning = log.New(io.MultiWriter(os.Stdout, warnFile), "[LifeGamer Event Engine][Warning]", log.Ldate | log.Ltime | log.Lshortfile)
}

/*
	Default usage: 

	lg.<*log.Logger>.Println(str1, str2, ...)
*/