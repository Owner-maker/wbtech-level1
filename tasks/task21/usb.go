package main

type Usb interface {
	connectWithUsbCable()
	disconnectWithUsbCable()
}
