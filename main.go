package main

import (
	"NES_NEXT/processor/memory"
	"fmt"
)

func main() {
	fmt.Println("NES_Next Emulator")
	ram := processor.InitRam()
	ram.WriteByte(0x00f1, 0x35)
	ram.WriteByte(0x00f2, 0x13)
	fmt.Printf("%x", ram.ReadWord(0x00f1))
}
