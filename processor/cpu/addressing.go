package processor

import (
	ram "NES_NEXT/processor/memory"
)

func Implied(_ uint16) *byte {
	return nil
}

func Accumulator(_ uint16) *byte {
	return &GetCPU().a
}

func Immediate(address uint16) *byte {
	ram := ram.GetRam()
	value := ram.Read(uint16(address))
	return value
}

// NOT REALIZED

func Absolute(address uint16) *byte {
	return nil
}

func XIndexedAbsolute(address uint16) *byte {
	return nil
}

func YIndexedAbsolute(address uint16) *byte {
	return nil
}

func AbsoluteInderect(address uint16) *byte {
	return nil
}

func ZeroPage(address uint16) *byte {
	return nil
}

func XIndexedZeroPage(address uint16) *byte {
	return nil
}

func XIndexedZeroPageInderect(address uint16) *byte {
	return nil
}

func YIndexedZeroPage(address uint16) *byte {
	return nil
}

func YIndexedZeroPageInderect(address uint16) *byte {
	return nil
}

func Relative(address uint16) *byte {
	return nil
}
