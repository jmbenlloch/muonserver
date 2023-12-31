package main

import (
	"fmt"
	"log"
)

func decodeFrame(recvChannel chan Frame, sendChannel chan *Frame) {
	for {
		frame := <-recvChannel
		log.Printf("length %d, %s", len(frame.Payload), frame.Command)
		switch frame.Command {
		case FEB_SET_RECV:
			log.Println("get device ok")
			deviceFound(frame, sendChannel)
		case FEB_WR_PMR:
			log.Println("pmr config received")
			pmr_config_ack(frame, sendChannel)
		case FEB_WR_SCR:
			log.Println("scr config received")
			scr_config_ack(frame, sendChannel)
		case FEB_WR_FIL:
			log.Println("fil config received")
			fil_config_ack(frame, sendChannel)
		case FEB_GEN_INIT:
			log.Println("run control received")
			run_control(frame, sendChannel)
		case FEB_GET_RATE:
			log.Println("get rate received")
			get_rate(frame, sendChannel)
		case FEB_RD_CDR:
			log.Println("data read received")
			read_data(frame, sendChannel)
		case FEB_OK:
			log.Println("feb ok")
		case FEB_DATA_CDR:
			log.Println("data cdr")
		case FEB_EOF_CDR:
			log.Println("End of data")
		case FEB_OK_SCR:
			log.Println("CITIROC slow control OK")
		case FEB_OK_PMR:
			log.Println("CITIROC probe OK")
		case FEB_OK_FIL:
			log.Println("FPGA input logic OK")
		default:
			log.Fatalf("Unkown response command %s", frame.Command)
		}
	}
}

func deviceFound(frame Frame, sendChannel chan *Frame) {
	payload := make([]byte, 2+6)           // register + mac address
	copy(payload[0:2], []byte{0x00, 0x00}) // register
	copy(payload[2:], "FEB_rev3_FLX7.003")
	newFrame, _ := buildFrame(frame.Destination, frame.Source, FEB_OK, payload)
	sendChannel <- newFrame
}

func pmr_config_ack(frame Frame, sendChannel chan *Frame) {
	payload := make([]byte, 2+6)           // register + mac address
	copy(payload[0:2], []byte{0x00, 0x00}) // register
	newFrame, _ := buildFrame(frame.Destination, frame.Source, FEB_OK_PMR, payload)
	sendChannel <- newFrame
}

func scr_config_ack(frame Frame, sendChannel chan *Frame) {
	payload := make([]byte, 2+6)           // register + mac address
	copy(payload[0:2], []byte{0x00, 0x00}) // register
	newFrame, _ := buildFrame(frame.Destination, frame.Source, FEB_OK_SCR, payload)
	sendChannel <- newFrame
}

func fil_config_ack(frame Frame, sendChannel chan *Frame) {
	payload := make([]byte, 2+6)           // register + mac address
	copy(payload[0:2], []byte{0x00, 0x00}) // register
	newFrame, _ := buildFrame(frame.Destination, frame.Source, FEB_OK_FIL, payload)
	sendChannel <- newFrame
}

func run_control(frame Frame, sendChannel chan *Frame) {
	fmt.Println(frame)
	//log.Printf("[%s] %x", frame.Source.String(), frame.Payload)

	switch frame.Payload[0] {
	case 0:
		fmt.Println("stop run")
	case 1:
		fmt.Println("reset run")
	case 2:
		fmt.Println("start run")
	}

	payload := make([]byte, 2+6)           // register + mac address
	copy(payload[0:2], []byte{0x00, 0x00}) // register
	newFrame, _ := buildFrame(frame.Destination, frame.Source, FEB_OK, payload)
	sendChannel <- newFrame
}

func get_rate(frame Frame, sendChannel chan *Frame) {
	payload := make([]byte, 2+6)           // register + mac address
	copy(payload[0:2], []byte{0x00, 0x00}) // register
	copy(payload[2:], "FEB_rev3_FLX7.003")
	newFrame, _ := buildFrame(frame.Destination, frame.Source, FEB_OK, payload)
	sendChannel <- newFrame
}

func read_data(frame Frame, sendChannel chan *Frame) {
	payload := make([]byte, 2+6)           // register + mac address
	copy(payload[0:2], []byte{0x00, 0x00}) // register
	copy(payload[2:], "FEB_rev3_FLX7.003")
	newFrame, _ := buildFrame(frame.Destination, frame.Source, FEB_OK, payload)
	sendChannel <- newFrame
}
