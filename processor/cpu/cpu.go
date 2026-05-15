package processor

import "fmt"

type processorStatus struct {
	carry             bool
	zero              bool
	interrupt_disable bool
	decimal_mode      bool
	break_command     bool
	overflow          bool
	negative          bool
}

type cpu struct {
	pc  int16           // Program Counter
	sp  int8            // Stack Pointer
	a   int8            // Accumulator
	irx int8            // Index Register X
	iry int8            // Index Register Y
	ps  processorStatus // Processor Status Register

	tick int8
}

func InitCpu() cpu {
	return cpu{
		ps: processorStatus{},
	}
}

func (cpu *cpu) Step(opcode byte) {
	fmt.Printf("%x", opcode)
}

func PrintCpuState(cpu cpu) {
	fmt.Printf("%+v", cpu)
}
