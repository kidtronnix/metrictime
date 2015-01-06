package main

import (
	"log"
	"math"
	"time"
)

func main() {
	y, m, d := time.Now().Date()
	now := time.Now().Unix()
	startDay := time.Date(y, m, d, 0, 0, 0, 0, time.Local).Unix()
	secs := now - startDay
	metSecs := float64(secs) / 0.864

	metH := math.Trunc(metSecs / (100 * 100))
	metM := math.Trunc((metSecs - metH*100*100) / (100))
	metS := math.Trunc((metSecs - metH*100*100 - metM*100))
	log.Printf("or %02v:%02v:%02v (metric time)", metH, metM, metS)
}
