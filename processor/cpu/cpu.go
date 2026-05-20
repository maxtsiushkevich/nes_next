package processor

import (
	ram "NES_NEXT/processor/memory"
	"fmt"
	"sync"
)

type CPU struct {
	pc  uint16 // Program Counter
	sp  byte   // Stack Pointer
	a   byte   // Accumulator
	irx byte   // Index Register X
	iry byte   // Index Register Y

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

	cycles byte
}

type CpuFlag byte

const (
	Negative         CpuFlag = 7
	Overflow         CpuFlag = 6
	Break            CpuFlag = 4
	Decimal          CpuFlag = 3
	InterruptDisable CpuFlag = 2
	Zero             CpuFlag = 1
	Carry            CpuFlag = 0
)

var lock = &sync.Mutex{}
var cpu *CPU

func GetCPU() *CPU {
	if cpu == nil {
		lock.Lock()
		defer lock.Unlock()
		if cpu == nil {
			cpu = &CPU{}
			cpu.Reset()
		}
	}

	return cpu
}

func (cpu *CPU) Reset() {
	cpu.a = 0
	cpu.irx = 0
	cpu.iry = 0

	cpu.ps = 0x24
	cpu.sp = 0xFD

	ram := ram.GetRam()
	cpu.pc = uint16(*ram.Read(0xFFFC)) |
		uint16(*ram.Read(0xFFFD))<<8
}

// func (cpu *CPU) Step() {
// 	ram := ram.GetRam()

// 	// Fecth opcode
// 	opcode := *ram.Read(cpu.pc)
// 	cpu.pc++

// 	inst := OpcodeTable[opcode]

// 	// Fetch operands
// 	// PC +1 or +2

// 	operand := inst.Addressing_mode(0x1100)

// 	// Fetch value from memory

// 	// EXECUTE
// 	inst.Execute(cpu, operand)
// }

func PrintCpuState(cpu *CPU) {
	fmt.Printf("%+v\n", cpu)
	fmt.Printf("%#08b\n", cpu.ps)

}

func (cpu *CPU) GetFlag(flag CpuFlag) byte {
	if cpu.ps&(1<<flag) != 0 {
		return 1
	}
	return 0
}

func (cpu *CPU) SetFlag(state bool, flag CpuFlag) {
	if state {
		cpu.ps = cpu.ps | (1 << flag)
	} else {
		cpu.ps = cpu.ps &^ (1 << flag)
	}
}

func (cpu *CPU) push(value byte) {
	ram := ram.GetRam()
	addr := 0x0100 | uint16(cpu.sp)
	ram.Write(addr, value)
	cpu.sp--
}

func (cpu *CPU) pop() byte {
	ram := ram.GetRam()
	cpu.sp++
	addr := 0x0100 | uint16(cpu.sp)
	return *ram.Read(addr)
}
