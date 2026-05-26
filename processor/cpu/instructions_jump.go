package cpu

// Return from Interrupt
func Rti(cpu *CPU, _ Operand) {
	cpu.ps = cpu.pop()

	cpu.ps |= (1 << 5)
	cpu.ps &^= (1 << 4)

	lo := cpu.pop()
	hi := cpu.pop()

	cpu.pc = uint16(lo) | uint16(hi)<<8
}

// Return from Subroutine
func Rts(cpu *CPU, _ Operand) {
	lo := cpu.pop()
	hi := cpu.pop()
	cpu.pc = uint16(lo) | uint16(hi)<<8
	cpu.pc += 1
}

// Jump
func Jmp(cpu *CPU, op Operand) {
	cpu.pc = op.Addr
}

// Jump to Subroutine
func Jsr(cpu *CPU, op Operand) {
	tmp_pc := cpu.pc - 1 // Return address on the stack points 1 byte before the start of the next instruction, rather than directly at the instruction.
	// This is because RTS increments the program counter before the next instruction is fetched.

	lo := byte(tmp_pc & 0x00FF)
	hi := byte(tmp_pc >> 8)
	cpu.push(hi)
	cpu.push(lo)

	cpu.pc = op.Addr
}

// Break
func Brk(cpu *CPU, _ Operand) {
	cpu.Interrupt(BRK)
}
