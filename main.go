package main

func main() {
	recv := make(chan interface{})

	go DataHandler(recv)

	status_chan := make(chan Status)
	go Portfolio(status_chan, recv)

	order_chan := make(chan Order)
	go BrokerExecutor(order_chan, recv)

	bar_chan := make(chan Bar)
	go Strategy(bar_chan, recv)

	for msg := range recv {
		switch msg := msg.(type) {
		case Status:
			status_chan <- msg
		case Order:
			order_chan <- msg
		case Bar:
			bar_chan <- msg
		}
	}
}
