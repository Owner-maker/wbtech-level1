package main

// объявление структуры CardReader -> содержит в себе поле типа MicroSd и имплементирует интерфейс usb
// 2 метода: подключение по USB и отключение по USB
// каждый из методов обращается к полю типа MicroSd и вызывает нужный метод, таким образом, данный структура типа WebCameraAdapter является
// своего рода посредником или прокси сущностью, предоствляющий функционал MicroSd, но посредством методов, имплементированных от другого интерфейса usb

type CardReader struct {
	microSd MicroSd
}

func (c CardReader) connectWithUsbCable() {
	c.microSd.insert()
	c.microSd.copyData()
}

func (c CardReader) disconnectWithUsbCable() {
	c.microSd.disconnect()
}
