package main

import (
    "bytes"
    "fmt"
    "log"
    "time"
)

type Logger struct {
    buf bytes.Buffer
}

func (l *Logger) Now() time.Time {
    current_time := time.Now().UTC()
    return current_time
}

func (l *Logger) Debug(logmsg interface{}) {
    if (l.buf.Len() > 0) {
        l.buf.Next(l.buf.Len())
    }
    logger_ := log.New(&l.buf, "[Debug] | ", log.Lshortfile)
    logger_.Print(logmsg)
    fmt.Printf("%s | %s", l.Now(), &l.buf)
}

func (l *Logger) Error(logmsg interface{}) {
    if (l.buf.Len() > 0) {
        l.buf.Next(l.buf.Len())
    }
    logger_ := log.New(&l.buf, "[Error] | ", log.Lshortfile)
    logger_.Print(logmsg)
    fmt.Printf("%s | %s", l.Now(), &l.buf)
}
