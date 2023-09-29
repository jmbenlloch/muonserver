package main

import "fmt"

func main() {
	iface := getNetworkInterface("wlp2s0")
	connection := createSocket(iface)
	sendFrameChannel := make(chan *Frame, 2000)
	recvFrameChannel := make(chan Frame, 2000)

	fmt.Println(iface)
	fmt.Println(connection)
	go sendFrameViaSocket(sendFrameChannel, connection)
	go receiveMessages(recvFrameChannel, connection, iface.MTU)
	go decodeFrame(recvFrameChannel, sendFrameChannel)

	for {
	}
}
