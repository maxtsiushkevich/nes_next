package cpu

import (
	"NES_NEXT/utils"
)

// Bitwise OR
func Ora(cpu *CPU, op Operand) {
	cpu.a |= *op.Value

	cpu.SetFlag(cpu.a == 0, Zero)
	cpu.SetFlag(cpu.a&0x80 != 0, Negative)
}

// Bitwise XOR
func Eor(cpu *CPU, op Operand) {
	cpu.a ^= *op.Value

	cpu.SetFlag(cpu.a == 0, Zero)
	cpu.SetFlag(cpu.a&0x80 != 0, Negative)
}

// Bitwise And
func And(cpu *CPU, op Operand) {
	cpu.a &= *op.Value

	cpu.SetFlag(cpu.a == 0, Zero)
	cpu.SetFlag(cpu.a&0x80 != 0, Negative)
}

// Bit Test
func Bit(cpu *CPU, op Operand) {
	result := cpu.a & *op.Value

	cpu.SetFlag(result == 0, Zero)
	cpu.SetFlag(utils.GetBit(*op.Value, 6) != 0, Overflow)
	cpu.SetFlag(utils.GetBit(*op.Value, 7) != 0, Negative)
}
