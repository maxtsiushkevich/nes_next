package processor

type ram struct {
	memory [65536]byte
}

func InitRam() *ram {
	return &ram{}
}

func (ram *ram) ReadByte(address uint16) byte {
	return ram.memory[address]
}

func (ram *ram) WriteByte(address uint16, value byte) {
	ram.memory[address] = value
}

func (ram *ram) ReadWord(address uint16) uint16 {
	lo := uint16(ram.memory[address])
	hi := uint16(ram.memory[address+1])
	return (hi << 8) + lo
}
