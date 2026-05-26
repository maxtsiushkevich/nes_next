package memory

type RAM struct {
	memory [65536]byte
}

func NewRAM() *RAM {
	return &RAM{}
}

func (ram *RAM) Read(address uint16) *byte {
	return &ram.memory[address]
}

func (ram *RAM) Write(address uint16, value byte) {
	ram.memory[address] = value
}

func (ram *RAM) WriteBlock(startAddr uint16, block []byte) {
	start := int(startAddr)

	if start >= len(ram.memory) {
		return
	}

	n := min(len(block), len(ram.memory)-start)

	copy(ram.memory[start:start+n], block[:n])
}
