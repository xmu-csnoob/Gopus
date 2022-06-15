package log

import "log"

type Logger struct {
}

func (logger *Logger) PrintInfo(info string) {
	log.Printf("\033[1;34;40m%s\033[0m", "INFO : "+info)
}
