package cpu

// Transfer X to Stack Pointer
func Txs(cpu *CPU, _ Operand) {
	cpu.sp = cpu.irx
}

// Transfer Stack Pointer to X
func Tsx(cpu *CPU, _ Operand) {
	cpu.irx = cpu.sp

	cpu.SetFlag(cpu.irx == 0, Zero)
	cpu.SetFlag(cpu.irx&0x80 != 0, Negative)
}

// Push A
func Pha(cpu *CPU, _ Operand) {
	cpu.push(cpu.a)
}

// Pull A
func Pla(cpu *CPU, _ Operand) {
	cpu.a = cpu.pop()

	cpu.SetFlag(cpu.a == 0, Zero)
	cpu.SetFlag(cpu.a&0x80 != 0, Negative)
}

// Push Processor Status
func Php(cpu *CPU, _ Operand) {
	var flags byte = cpu.ps
	flags |= (1 << 5) | (1 << 4)
	cpu.push(flags) // flags NV11DIZC
}

// Pull Processor Status
func Plp(cpu *CPU, _ Operand) {
	cpu.ps = cpu.pop()

	cpu.ps |= (1 << 5)
	cpu.ps &^= (1 << 4)
}
