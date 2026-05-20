package processor

import (
	"sync"
)

type RAM struct {
	memory [65536]byte
}

var lock = &sync.Mutex{}
var ram *RAM

func GetRam() *RAM {
	if ram == nil {
		lock.Lock()
		defer lock.Unlock()
		if ram == nil {
			ram = &RAM{}
		}
	}

	return ram
}

func (ram *RAM) Read(address uint16) *byte {
	return &ram.memory[address]
}

func (ram *RAM) Write(address uint16, value byte) {
	ram.memory[address] = value
}

func (ram *RAM) WriteBlock(start_addr uint16, block []byte) {
	// not realized
}
