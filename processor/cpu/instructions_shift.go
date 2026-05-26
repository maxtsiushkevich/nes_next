package cpu

import (
	"NES_NEXT/utils"
)

// Arithmetic Shift Left
func Asl(cpu *CPU, op Operand) {
	cpu.SetFlag(*op.Value&0x80 != 0, Carry)
	*op.Value = *op.Value << 1

	cpu.SetFlag(*op.Value == 0, Zero)
	cpu.SetFlag(*op.Value&0x80 != 0, Negative)
}

// Rotate Left
func Rol(cpu *CPU, op Operand) {
	oldCarry := cpu.GetFlag(Carry)

	cpu.SetFlag(utils.GetBit(*op.Value, 7) != 0, Carry)

	result := (*op.Value << 1) + oldCarry
	*op.Value = result

	cpu.SetFlag(result == 0, Zero)
	cpu.SetFlag(result&0x80 != 0, Negative)
}

// Rotate Right
func Ror(cpu *CPU, op Operand) {
	oldCarry := cpu.GetFlag(Carry)

	cpu.SetFlag(utils.GetBit(*op.Value, 0) != 0, Carry)

	result := (*op.Value >> 1) | (oldCarry << 7)
	*op.Value = result

	cpu.SetFlag(result == 0, Zero)
	cpu.SetFlag(result&0x80 != 0, Negative)
}

// Logical Shift Right
func Lsr(cpu *CPU, op Operand) {
	cpu.SetFlag(utils.GetBit(*op.Value, 0) != 0, Carry)

	*op.Value >>= 1

	cpu.SetFlag(*op.Value == 0, Zero)
	cpu.SetFlag(false, Negative)
}
