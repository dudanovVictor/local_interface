package main

import (
	"fmt"
	"time"
)

type slog struct {
	arr []string
}

func (l *slog) Logf(s string, a ...interface{}) {
	timeLabel := time.Now().Format(time.RFC3339Nano)
	l.arr = append(l.arr, timeLabel+" "+fmt.Sprintf(s, a...))
}
func (l *slog) All() []string {
	return l.arr
}

func NewSimpleLog() *slog {
	return &slog{arr: []string{}}
}
