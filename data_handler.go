package main

import (
	"math/rand"
	"time"
)

type Bar struct {
	symbol string
	o      float64
	h      float64
	l      float64
	c      float64
	v      int64
}

func DataHandler(output chan interface{}) {
	for {
		b := Bar{
			symbol: "TSLA",
			o:      rand.Float64(),
			h:      rand.Float64(),
			l:      rand.Float64(),
			c:      rand.Float64(),
			v:      1,
		}

		output <- b

		time.Sleep(1 * time.Second)
	}
}
