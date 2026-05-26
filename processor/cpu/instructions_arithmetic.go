package cpu

// Add with Carry
func Adc(cpu *CPU, op Operand) {
	carry := cpu.GetFlag(Carry)

	sum := uint16(cpu.a) +
		uint16(*op.Value) +
		uint16(carry)

	result := byte(sum)

	cpu.SetFlag(sum > 0xFF, Carry)
	cpu.SetFlag(result == 0, Zero)
	cpu.SetFlag(((result^cpu.a)&(result^*op.Value)&0x80) != 0, Overflow)
	cpu.SetFlag(result&0x80 != 0, Negative)

	cpu.a = result
}

// Subtract with Carry
func Sbc(cpu *CPU, op Operand) {
	carry := cpu.GetFlag(Carry)

	// Invert
	value := *op.Value ^ 0xFF

	sum := uint16(cpu.a) +
		uint16(value) +
		uint16(carry)

	result := byte(sum)

	cpu.SetFlag(sum > 0xFF, Carry)
	cpu.SetFlag(result == 0, Zero)
	cpu.SetFlag(
		((cpu.a^result)&(cpu.a^*op.Value)&0x80) != 0,
		Overflow,
	)
	cpu.SetFlag(result&0x80 != 0, Negative)

	cpu.a = result
}

// Increment Memory
func Inc(cpu *CPU, op Operand) {
	// ram := ram.GetRam()
	result := *cpu.Mem.Read(op.Addr)
	result++

	cpu.SetFlag(result == 0, Zero)
	cpu.SetFlag(result&0x80 != 0, Negative)

	cpu.Mem.Write(op.Addr, result)
}

// Decrement Memory
func Dec(cpu *CPU, op Operand) {
	// ram := ram.GetRam()
	result := *cpu.Mem.Read(op.Addr)
	result--

	cpu.SetFlag(result == 0, Zero)
	cpu.SetFlag(result&0x80 != 0, Negative)

	cpu.Mem.Write(op.Addr, result)
}

// Increment X
func Inx(cpu *CPU, _ Operand) {
	cpu.irx++

	cpu.SetFlag(cpu.irx == 0, Zero)
	cpu.SetFlag(cpu.irx&0x80 != 0, Negative)
}

// Decrement X
func Dex(cpu *CPU, _ Operand) {
	cpu.irx--

	cpu.SetFlag(cpu.irx == 0, Zero)
	cpu.SetFlag(cpu.irx&0x80 != 0, Negative)
}

// Increment Y
func Iny(cpu *CPU, _ Operand) {
	cpu.iry++

	cpu.SetFlag(cpu.iry == 0, Zero)
	cpu.SetFlag(cpu.iry&0x80 != 0, Negative)
}

// Decrement Y
func Dey(cpu *CPU, _ Operand) {
	cpu.iry--

	cpu.SetFlag(cpu.iry == 0, Zero)
	cpu.SetFlag(cpu.iry&0x80 != 0, Negative)
}
