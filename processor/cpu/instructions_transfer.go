package cpu

// Transfer A to X
func Tax(cpu *CPU, _ Operand) {
	cpu.irx = cpu.a

	cpu.SetFlag(cpu.irx == 0, Zero)
	cpu.SetFlag(cpu.irx&0x80 != 0, Negative)
}

// Transfer A to Y
func Tay(cpu *CPU, _ Operand) {
	cpu.iry = cpu.a

	cpu.SetFlag(cpu.iry == 0, Zero)
	cpu.SetFlag(cpu.iry&0x80 != 0, Negative)
}

// Transfer X to A
func Txa(cpu *CPU, _ Operand) {
	cpu.a = cpu.irx

	cpu.SetFlag(cpu.a == 0, Zero)
	cpu.SetFlag(cpu.a&0x80 != 0, Negative)
}

// Transfer Y to A
func Tya(cpu *CPU, _ Operand) {
	cpu.a = cpu.iry

	cpu.SetFlag(cpu.a == 0, Zero)
	cpu.SetFlag(cpu.a&0x80 != 0, Negative)
}
