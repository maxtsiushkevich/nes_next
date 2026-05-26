package cpu

func branch(cpu *CPU, offset byte) {
	rel := int8(offset)

	cpu.pc = uint16(int32(cpu.pc) + int32(rel))
}

// Branch if Carry Clear
func Bcc(cpu *CPU, op Operand) {
	if cpu.GetFlag(Carry) == 0 {
		branch(cpu, *op.Value)
	}
}

// Branch if Carry Set
func Bcs(cpu *CPU, op Operand) {
	if cpu.GetFlag(Carry) == 1 {
		branch(cpu, *op.Value)
	}
}

// Branch if Equal
func Beq(cpu *CPU, op Operand) {
	if cpu.GetFlag(Zero) == 1 {
		branch(cpu, *op.Value)
	}
}

// Branch if Minus
func Bmi(cpu *CPU, op Operand) {
	if cpu.GetFlag(Negative) == 1 {
		branch(cpu, *op.Value)
	}
}

// Branch if Not Equal
func Bne(cpu *CPU, op Operand) {
	if cpu.GetFlag(Zero) == 0 {
		branch(cpu, *op.Value)
	}
}

// Branch if Plus
func Bpl(cpu *CPU, op Operand) {
	if cpu.GetFlag(Negative) == 0 {
		branch(cpu, *op.Value)
	}
}

// Branch if Overflow Clear
func Bvc(cpu *CPU, op Operand) {
	if cpu.GetFlag(Overflow) == 0 {
		branch(cpu, *op.Value)
	}
}

// Branch if Overflow Set
func Bvs(cpu *CPU, op Operand) {
	if cpu.GetFlag(Overflow) == 1 {
		branch(cpu, *op.Value)
	}
}
