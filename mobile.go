package mobile

func START(os, port string) {
	flagOS = os
	flagPORT = port
	go startWebserver()
}
