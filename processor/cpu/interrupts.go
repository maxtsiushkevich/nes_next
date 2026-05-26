package cpu

type InterruptType int

const (
	NMI InterruptType = iota
	IRQ
	BRK
	RESET
)

func (cpu *CPU) Interrupt(kind InterruptType) {
	switch kind {
	case RESET:
		lo := *cpu.Mem.Read(0xFFFC)
		hi := *cpu.Mem.Read(0xFFFD)

		cpu.pc = uint16(lo) | uint16(hi)<<8
		cpu.sp = 0xFD
		cpu.ps = 0x24
		return
	}

	// IRQ маскируется флагом I
	if kind == IRQ && cpu.GetFlag(InterruptDisable) == 1 {
		return
	}

	// push PC
	hi := byte(cpu.pc >> 8)
	lo := byte(cpu.pc)

	cpu.push(hi)
	cpu.push(lo)

	// push status
	flags := cpu.ps

	// bit5 always set
	flags |= (1 << 5)

	// BRK only
	if kind == BRK {
		flags |= (1 << 4)
	} else {
		flags &^= (1 << 4)
	}

	cpu.push(flags)

	// disable IRQ
	cpu.SetFlag(true, InterruptDisable)

	var vector uint16

	switch kind {
	case NMI:
		vector = 0xFFFA

	case IRQ, BRK:
		vector = 0xFFFE
	}

	lo = *cpu.Mem.Read(vector)
	hi = *cpu.Mem.Read(vector + 1)

	cpu.pc = uint16(lo) | uint16(hi)<<8
}

func (cpu *CPU) RequestNMI() {
	cpu.nmi = true
}

func (cpu *CPU) RequestIRQ() {
	cpu.irq = true
}

func (cpu *CPU) Reset() {
	cpu.a = 0
	cpu.irx = 0
	cpu.iry = 0

	cpu.nmi = false
	cpu.irq = false

	cpu.Interrupt(RESET)
}
