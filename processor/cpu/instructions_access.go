package cpu

// Load A
func Lda(cpu *CPU, op Operand) {
	cpu.a = *op.Value

	cpu.SetFlag(cpu.a == 0, Zero)
	cpu.SetFlag(cpu.a&0x80 != 0, Negative)
}

// Load Y
func Ldy(cpu *CPU, op Operand) {
	cpu.iry = *op.Value

	cpu.SetFlag(cpu.iry == 0, Zero)
	cpu.SetFlag(cpu.iry&0x80 != 0, Negative)
}

// Load X
func Ldx(cpu *CPU, op Operand) {
	cpu.irx = *op.Value

	cpu.SetFlag(cpu.irx == 0, Zero)
	cpu.SetFlag(cpu.irx&0x80 != 0, Negative)
}

// Store A
func Sta(cpu *CPU, op Operand) {
	// ram := ram.GetRam()
	cpu.Mem.Write(op.Addr, cpu.a)
}

// Store X
func Stx(cpu *CPU, op Operand) {
	// ram := ram.GetRam()
	cpu.Mem.Write(op.Addr, cpu.irx)
}

// Store Y
func Sty(cpu *CPU, op Operand) {
	// ram := ram.GetRam()
	cpu.Mem.Write(op.Addr, cpu.iry)
}
