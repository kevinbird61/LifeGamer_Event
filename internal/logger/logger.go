package logger

import (
	"log"
	"io"
)

type Logger struct {
	mu		sync.Mutex 	// ensure atomic writes
	prefix 	string		// prefix to write at the beginning of each line
	flag 	int 		// properties
	out		io.Writer 	// destination for output
	buf		[]byte		// for accumulating text to write
}

// initial function of Logging system
func (l *Logger) init() {
	// granularity
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	// prefix
	log.SetPrefix("[LifeGamer - Event Engine]")
}

// set prefix of log
func (l *Logger) set_prefix(s string) {
	log.SetPrefix("[%v]", s)
}

