package main

import "math/rand"

func Strategy(input chan Bar, output chan interface{}) {
	for msg := range input {
		qty := rand.Int63n(200) - 100
		if qty != 0 {
			if rand.Float64() > 0.5 {
				// Market Order
				output <- MarketOrder{symbol: "TSLA", qty: qty}
			} else {
				// Limit Order
				output <- LimitOrder{symbol: "TSLA", qty: qty, price: msg.c}
			}
		}
	}
}
