package cpu

import (
	"fmt"
)

type Memory interface {
	Read(addr uint16) *byte
	Write(addr uint16, value byte)
	WriteBlock(startAddr uint16, block []byte)
	SaveToFile()
}

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

	Mem Memory

	cycles uint64

	nmi bool
	irq bool
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

func NewCPU(mem Memory) *CPU {
	cpu := &CPU{
		Mem: mem,
	}
	return cpu
}

func (cpu *CPU) Step() {
	opcode := *cpu.Mem.Read(cpu.pc)
	cpu.pc++

	inst, ok := OpcodeTable[opcode]
	if !ok {
		panic(fmt.Sprintf(
			"unknown opcode %02X at %04X",
			opcode,
			cpu.pc,
		))
	}

	operand := inst.AddressingMode(cpu)

	fmt.Printf(
		"OPC=%02X, X=%02X Y=%02X A=%02X SP=%02X PC=%04X P=%02X C=%d\n",
		opcode,
		cpu.irx,
		cpu.iry,
		cpu.a,
		cpu.sp,
		cpu.pc,
		cpu.ps,
		cpu.cycles*3,
	)

	cpu.cycles += uint64(inst.Cycles)

	inst.Execute(cpu, operand)

	// interrupt polling
	if cpu.nmi {
		cpu.nmi = false
		cpu.Interrupt(NMI)
	} else if cpu.irq {
		cpu.irq = false
		cpu.Interrupt(IRQ)
	}
}

func PrintCpuState(cpu *CPU) {
	fmt.Printf(
		"PC:%04X A:%02X X:%02X Y:%02X SP:%02X P:%02X\n",
		cpu.pc,
		cpu.a,
		cpu.irx,
		cpu.iry,
		cpu.sp,
		cpu.ps,
	)

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
	addr := 0x0100 | uint16(cpu.sp)
	cpu.Mem.Write(addr, value)
	cpu.sp--
}

func (cpu *CPU) pop() byte {
	cpu.sp++
	addr := 0x0100 | uint16(cpu.sp)
	return *cpu.Mem.Read(addr)
}

func (cpu *CPU) FetchByte() byte {
	value := *cpu.Mem.Read(cpu.pc)
	cpu.pc++

	return value
}

func (cpu *CPU) FetchWord() uint16 {
	lo := uint16(cpu.FetchByte())
	hi := uint16(cpu.FetchByte())

	return lo | (hi << 8)
}
