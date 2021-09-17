package main

func main() {
	client := &client{}
	mac := &mac{}
	client.insertSquareUsbInComputer(mac)
	windows := &windows{}
	adapter := &windowsAdapter{windows}
	client.insertSquareUsbInComputer(adapter)

}
