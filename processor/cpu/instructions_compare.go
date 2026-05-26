package cpu

import (
	"NES_NEXT/utils"
)

func compare(cpu *CPU, register byte, value byte) {
	res := register - value

	cpu.SetFlag(utils.GetBit(res, 7) != 0, Negative)
	cpu.SetFlag(register >= value, Carry)
	cpu.SetFlag(res == 0, Zero)
}

// Compare A
func Cmp(cpu *CPU, op Operand) {
	compare(cpu, cpu.a, *op.Value)
}

// Compare X
func Cpx(cpu *CPU, op Operand) {
	compare(cpu, cpu.irx, *op.Value)
}

// Compare Y
func Cpy(cpu *CPU, op Operand) {
	compare(cpu, cpu.iry, *op.Value)
}
