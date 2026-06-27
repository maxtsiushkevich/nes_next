package cpu

func pageCrossed(a, b uint16) bool {
	return (a & 0xFF00) != (b & 0xFF00)
}

type Operand struct {
	Value *byte
	Addr  uint16
}

func Implied(cpu *CPU) Operand {
	return Operand{}
}

func Accumulator(cpu *CPU) Operand {
	return Operand{
		Value: nil,
		Addr:  0,
	}
}

func Absolute(cpu *CPU) Operand {
	address := cpu.FetchWord()
	value := cpu.Mem.Read(uint16(address))

	return Operand{
		Value: value,
		Addr:  address,
	}
}

func Immediate(cpu *CPU) Operand {
	val := cpu.FetchByte()

	return Operand{
		Value: &val,
		Addr:  0,
	}
}

func XIndexedAbsolute(cpu *CPU) Operand {
	base := cpu.FetchWord()
	addr := base + uint16(cpu.irx)

	if pageCrossed(base, addr) {
		cpu.cycles++
	}

	value := cpu.Mem.Read(addr)

	return Operand{
		Value: value,
		Addr:  addr,
	}
}

func YIndexedAbsolute(cpu *CPU) Operand {
	base := cpu.FetchWord()
	addr := base + uint16(cpu.iry)

	if pageCrossed(base, addr) {
		cpu.cycles++
	}

	value := cpu.Mem.Read(addr)

	return Operand{
		Value: value,
		Addr:  addr,
	}
}

func AbsoluteIndirect(cpu *CPU) Operand {
	ptr := cpu.FetchWord()
	lo := cpu.Mem.Read(ptr)

	// Emulate bug
	hiAddr := (ptr & 0xFF00) | ((ptr + 1) & 0x00FF)
	hi := cpu.Mem.Read(hiAddr)

	effectiveAddress := uint16(*lo) | (uint16(*hi) << 8)
	value := cpu.Mem.Read(uint16(effectiveAddress))

	return Operand{
		Value: value,
		Addr:  effectiveAddress,
	}
}

func ZeroPage(cpu *CPU) Operand {
	address := cpu.FetchByte()
	value := cpu.Mem.Read(uint16(address))
	return Operand{
		Value: value,
		Addr:  uint16(address),
	}
}

func XIndexedZeroPage(cpu *CPU) Operand {
	address := cpu.FetchByte() + cpu.irx
	value := cpu.Mem.Read(uint16(address))

	return Operand{
		Value: value,
		Addr:  uint16(address),
	}
}

func XIndexedZeroPageIndirect(cpu *CPU) Operand {
	base := cpu.FetchByte()
	ptr := base + cpu.irx

	lo := *cpu.Mem.Read(uint16(ptr))
	hi := *cpu.Mem.Read(uint16(byte(ptr + 1)))

	addr := uint16(lo) | (uint16(hi) << 8)

	return Operand{
		Value: cpu.Mem.Read(addr),
		Addr:  addr,
	}
}

func YIndexedZeroPage(cpu *CPU) Operand {
	address := cpu.FetchByte() + cpu.iry
	value := cpu.Mem.Read(uint16(address))

	return Operand{
		Value: value,
		Addr:  uint16(address),
	}
}

func YIndexedZeroPageIndirect(cpu *CPU) Operand {
	base := cpu.FetchByte()

	lo := *cpu.Mem.Read(uint16(base))
	hi := *cpu.Mem.Read(uint16(byte(base + 1)))

	baseAddr := uint16(lo) | (uint16(hi) << 8)
	addr := baseAddr + uint16(cpu.iry)

	if pageCrossed(baseAddr, addr) {
		cpu.cycles++
	}

	return Operand{
		Value: cpu.Mem.Read(addr),
		Addr:  addr,
	}
}

func Relative(cpu *CPU) Operand {
	value := cpu.FetchByte()
	return Operand{
		Value: &value,
		Addr:  0,
	}
}
