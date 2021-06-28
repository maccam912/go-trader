package main

import (
	"fmt"
)

type Position struct {
	symbol string
	qty    int64
	basis  float64
}

func Portfolio(input chan Status, output chan interface{}) {
	cash := 10000.0

	positions := make(map[string]*Position)

	for msg := range input {
		if msg.status == "FILLED" {
			// Filled order, add it to positions
			pos, ok := positions[msg.symbol]
			if ok {
				// Add or subtract from position
				pos.qty += msg.qty
				if pos.qty == 0 {
					delete(positions, msg.symbol)
				} else {
					pos.basis = ((pos.basis * float64(pos.qty)) + (msg.price * float64(msg.qty))) / float64(pos.qty+msg.qty)
				}
				cash -= float64(msg.qty) * msg.price
			} else {
				positions[msg.symbol] = &Position{symbol: msg.symbol, qty: msg.qty, basis: msg.price}
				cash -= float64(msg.qty) * msg.price
			}
		}
		fmt.Printf("Cash: %v\n", cash)
		fmt.Printf("Positions: %v\n\n\n", positions["TSLA"])
	}
}
