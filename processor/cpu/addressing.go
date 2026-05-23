package processor

import (
	ram "NES_NEXT/processor/memory"
)

type Operand struct {
	Value *byte
	Addr  uint16
}

func Implied() Operand {
	return Operand{}
}

func Accumulator() Operand {
	return Operand{
		Value: &cpu.a,
		Addr:  0,
	}
}

func Absolute() Operand {
	ram := ram.GetRam()
	cpu := GetCPU()

	address := cpu.FetchWord()
	value := ram.Read(uint16(address))

	return Operand{
		Value: value,
		Addr:  address,
	}
}

func Immediate() Operand {
	val := GetCPU().FetchByte()

	return Operand{
		Value: &val,
		Addr:  0,
	}
}

func XIndexedAbsolute() Operand {
	ram := ram.GetRam()
	cpu := GetCPU()

	address := cpu.FetchWord()
	address += uint16(cpu.irx)

	value := ram.Read(address)

	return Operand{
		Value: value,
		Addr:  address,
	}
}

func YIndexedAbsolute() Operand {
	ram := ram.GetRam()
	cpu := GetCPU()

	address := cpu.FetchWord()
	address += uint16(cpu.iry)

	value := ram.Read(address)

	return Operand{
		Value: value,
		Addr:  address,
	}
}

func AbsoluteIndirect() Operand {
	ram := ram.GetRam()
	cpu := GetCPU()

	ptr := cpu.FetchWord()

	lo := ram.Read(ptr)

	// Emulate bug
	hiAddr := (ptr & 0xFF00) | ((ptr + 1) & 0x00FF)
	hi := ram.Read(hiAddr)

	effectiveAddress := uint16(*lo) | (uint16(*hi) << 8)
	value := ram.Read(uint16(effectiveAddress))

	return Operand{
		Value: value,
		Addr:  effectiveAddress,
	}
}

func ZeroPage() Operand {
	ram := ram.GetRam()
	cpu := GetCPU()

	address := cpu.FetchByte()
	value := ram.Read(uint16(address))
	return Operand{
		Value: value,
		Addr:  uint16(address),
	}
}

func XIndexedZeroPage() Operand {
	ram := ram.GetRam()
	cpu := GetCPU()

	address := cpu.FetchByte() + cpu.irx
	value := ram.Read(uint16(address))

	return Operand{
		Value: value,
		Addr:  uint16(address),
	}
}

func XIndexedZeroPageIndirect() Operand {
	ram := ram.GetRam()
	cpu := GetCPU()

	base := cpu.FetchByte()
	ptr := base + cpu.irx // wrap 0xFF -> 0x00

	lo := *ram.Read(uint16(ptr))
	hi := *ram.Read(uint16(ptr + 1))

	addr :=
		uint16(lo) |
			(uint16(hi) << 8)

	return Operand{
		Value: ram.Read(addr),
		Addr:  addr,
	}
}

func YIndexedZeroPage() Operand {
	ram := ram.GetRam()
	cpu := GetCPU()

	address := cpu.FetchByte() + cpu.iry
	value := ram.Read(uint16(address))

	return Operand{
		Value: value,
		Addr:  uint16(address),
	}
}

func YIndexedZeroPageIndirect() Operand {
	ram := ram.GetRam()
	cpu := GetCPU()

	base := cpu.FetchByte()
	ptr := base + cpu.iry // wrap 0xFF -> 0x00

	lo := *ram.Read(uint16(ptr))
	hi := *ram.Read(uint16(ptr + 1))

	addr :=
		uint16(lo) |
			(uint16(hi) << 8)

	return Operand{
		Value: ram.Read(addr),
		Addr:  addr,
	}
}

func Relative() Operand {
	cpu := GetCPU()
	value := cpu.FetchByte()
	return Operand{
		Value: &value,
		Addr:  0,
	}
}
