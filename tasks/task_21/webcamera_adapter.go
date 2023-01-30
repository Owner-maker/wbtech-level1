package main

// объявление структуры WebCameraAdapter -> содержит в себе поле типа WebCamera и имплементирует интерфейс usb
// 2 метода: подключение по USB и отключение по USB
// каждый из методов обращается к полю типа WebCamera и вызывает нужный метод, таким образом, данный структура типа WebCameraAdapter является
// своего рода посредником или прокси сущностью, предоствляющий функционал WebCamera, но посредством методов, имплементированных от другого интерфейса usb

type WebCameraAdapter struct {
	webCamera WebCamera
}

func (w WebCameraAdapter) connectWithUsbCable() {
	w.webCamera.connect()
}

func (w WebCameraAdapter) disconnectWithUsbCable() {
	w.webCamera.disconnect()
}
