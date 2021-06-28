package main

import "math/rand"

type Order interface {
	getSymbol() string
	getQty() int64
}

type MarketOrder struct {
	symbol string
	qty    int64
}

func (o MarketOrder) getSymbol() string {
	return o.symbol
}

func (o MarketOrder) getQty() int64 {
	return o.qty
}

type LimitOrder struct {
	symbol string
	qty    int64
	price  float64
}

func (o LimitOrder) getSymbol() string {
	return o.symbol
}

func (o LimitOrder) getQty() int64 {
	return o.qty
}

type Status struct {
	status string
	symbol string
	qty    int64
	price  float64
}

type simBroker struct {
	orders []Order
}

func BrokerExecutor(input chan Order, output chan interface{}) {
	broker := simBroker{orders: []Order{}}

	for msg := range input {
		broker.orders = append(broker.orders, msg)
		output <- Status{status: "PENDING"}

		for _, order := range broker.orders {
			switch o := order.(type) {
			case MarketOrder:
				output <- Status{status: "FILLED", symbol: o.getSymbol(), qty: o.getQty(), price: rand.Float64()}
			case LimitOrder:
				output <- Status{status: "FILLED", symbol: o.getSymbol(), qty: o.getQty(), price: o.price}
			}
		}
		broker.orders = []Order{}
	}
}
