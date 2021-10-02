package main

func main() {
	hp := &hp{}
	epson := &epson{}
	mac := &mac{hp}
	mac.print()
	mac.setPrinter(epson)
	mac.print()

	windows := &windows{}
	windows.setPrinter(hp)
	windows.print()
	windows.setPrinter(epson)
	windows.print()
}
