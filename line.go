package main

import (
	"time"
)

type line struct{}

func newLine() line {
	return line{}
}

func (line) move() {
	time.Sleep(lineMoveTime)
}
