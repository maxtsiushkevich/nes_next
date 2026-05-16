package processor

import "fmt"

type cpu struct {
	pc  int16 // Program Counter
	sp  byte  // Stack Pointer
	a   byte  // Accumulator
	irx byte  // Index Register X
	iry byte  // Index Register Y

	ps byte // Processor Status Register
	// Flags NV1B DIZC
	// N - Negative
	// V - Overflow
	// 1 - always 1, not affected
	// B - Break, not affected to CPU
	// D - Decimal
	// I - Interrupt Disable
	// Z - Zero
	// C - Carry
	// Initial state - 0b00100100

	tick byte
}

type CpuFlag int

const (
	Negative         CpuFlag = 7
	Overflow         CpuFlag = 6
	Break            CpuFlag = 4
	Decimal          CpuFlag = 3
	InterruptDisable CpuFlag = 2
	Zero             CpuFlag = 1
	Carry            CpuFlag = 0
)

func InitCpu() cpu {
	return cpu{
		ps: 0b00100100, // Processor Status Register initial value
	}
}

func (cpu *cpu) Step(opcode byte) {
	fmt.Printf("%x", opcode)
}

func PrintCpuState(cpu cpu) {
	fmt.Printf("%+v\n", cpu)
	fmt.Printf("%#08b\n", cpu.ps)

}

func (cpu *cpu) GetFlag(flag CpuFlag) byte {
	if cpu.ps&(1<<flag) != 0 {
		return 1
	}
	return 0
}

func (cpu *cpu) SetFlag(state bool, flag CpuFlag) {
	if state {
		cpu.ps = cpu.ps | (1 << flag)
	} else {
		cpu.ps = cpu.ps &^ (1 << flag)
	}
}
