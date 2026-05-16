package processor

// Clear Carry
func (cpu *cpu) Clc() {
	cpu.SetFlag(false, Carry)
}

// Set Carry
func (cpu *cpu) Sec() {
	cpu.SetFlag(true, Carry)
}

// Clear Interrupt
func (cpu *cpu) Cli() {
	cpu.SetFlag(false, InterruptDisable)
}

// Set Interrupt
func (cpu *cpu) Sei() {
	cpu.SetFlag(true, InterruptDisable)
}

// Set Overflow
func (cpu *cpu) Clv() {
	cpu.SetFlag(true, Overflow)
}

// Clear Overflow
func (cpu *cpu) Cld() {
	cpu.SetFlag(false, Overflow)
}

// Set Decimal
func (cpu *cpu) Sed() {
	cpu.SetFlag(true, Decimal)
}

// Add with Carry
func (cpu *cpu) Adc(operand byte) {
	result := uint16(cpu.GetFlag(Carry)) + uint16(operand) + uint16(cpu.a)

	if result > 0xFF {
		cpu.SetFlag(true, Carry)
	}
	if result == 0 {
		cpu.SetFlag(true, Zero)
	}
	if (byte(result)^cpu.a)&(byte(result)^operand)&0x80 == 1 {
		cpu.SetFlag(true, Overflow)
	}
	if byte(result)&(1<<7) == 0b10000000 {
		cpu.SetFlag(true, Negative)
	}

	cpu.a = byte(result)

}

// Bitwise And
func (cpu *cpu) And(operand byte) {
	result := cpu.a & operand
	if result == 0 {
		cpu.SetFlag(true, Zero)
	}
	if result&(1<<7) == 0b10000000 {
		cpu.SetFlag(true, Negative)
	}
}
