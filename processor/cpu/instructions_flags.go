package cpu

// Clear Carry
func Clc(cpu *CPU, _ Operand) {
	cpu.SetFlag(false, Carry)
}

// Set Carry
func Sec(cpu *CPU, _ Operand) {
	cpu.SetFlag(true, Carry)
}

// Clear Interrupt
func Cli(cpu *CPU, _ Operand) {
	cpu.SetFlag(false, InterruptDisable)
}

// Set Interrupt
func Sei(cpu *CPU, _ Operand) {
	cpu.SetFlag(true, InterruptDisable)
}

// Clear Overflow
func Clv(cpu *CPU, _ Operand) {
	cpu.SetFlag(false, Overflow)
}

// Clear Decimal
func Cld(cpu *CPU, _ Operand) {
	cpu.SetFlag(false, Decimal)
}

// Set Decimal
func Sed(cpu *CPU, _ Operand) {
	cpu.SetFlag(true, Decimal)
}
