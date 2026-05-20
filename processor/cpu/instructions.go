package processor

import (
	ram "NES_NEXT/processor/memory"
	"NES_NEXT/utils"
)

// Clear Carry
func Clc(cpu *CPU, _ *byte) {
	cpu.SetFlag(false, Carry)
}

// Set Carry
func Sec(cpu *CPU, _ *byte) {
	cpu.SetFlag(true, Carry)
}

// Clear Interrupt
func Cli(cpu *CPU, _ *byte) {
	cpu.SetFlag(false, InterruptDisable)
}

// Set Interrupt
func Sei(cpu *CPU, _ *byte) {
	cpu.SetFlag(true, InterruptDisable)
}

// Clear Overflow
func Clv(cpu *CPU, _ *byte) {
	cpu.SetFlag(false, Overflow)
}

// Clear Decimal
func Cld(cpu *CPU, _ *byte) {
	cpu.SetFlag(false, Decimal)
}

// Set Decimal
func Sed(cpu *CPU, _ *byte) {
	cpu.SetFlag(true, Decimal)
}

// Add with Carry
func Adc(cpu *CPU, value *byte) {
	carry := cpu.GetFlag(Carry)

	sum := uint16(cpu.a) +
		uint16(*value) +
		uint16(carry)

	result := byte(sum)

	// Carry
	cpu.SetFlag(sum > 0xFF, Carry)

	// Zero
	cpu.SetFlag(result == 0, Zero)

	// Overflow
	cpu.SetFlag(((result^cpu.a)&(result^*value)&0x80) != 0, Overflow)

	// Negative
	cpu.SetFlag(result&0x80 != 0, Negative)

	cpu.a = result
}

// Bitwise And
func And(cpu *CPU, value *byte) {
	cpu.a &= *value

	// Zero
	cpu.SetFlag(cpu.a == 0, Zero)
	// Negative
	cpu.SetFlag(cpu.a&0x80 != 0, Negative)
}

// Arithmetic Shift Left
func Asl(cpu *CPU, value *byte) {
	// Carry
	cpu.SetFlag(*value&0x80 != 0, Carry)

	*value = *value << 1

	// Zero
	cpu.SetFlag(*value == 0, Zero)

	// Negative
	cpu.SetFlag(*value&0x80 != 0, Negative)
}

func Bit(cpu *CPU, value *byte) {
	result := cpu.a & *value

	// Zero
	cpu.SetFlag(result == 0, Zero)

	// Overflow
	cpu.SetFlag(utils.GetBit(*value, 6) != 0, Overflow)

	// Negative
	cpu.SetFlag(utils.GetBit(*value, 7) != 0, Negative)
}

// Break
func Brk(cpu *CPU, _ *byte) {
	// Store PC to stack
	tmp_pc := cpu.pc
	lo := byte(tmp_pc & 0x00FF)
	hi := byte(tmp_pc >> 8)
	cpu.push(hi) // hi
	cpu.push(lo) // lo

	// Store flags to stack
	var flags byte = cpu.ps
	flags |= (1 << 5) | (1 << 4)
	cpu.push(flags) // flags NV11DIZC

	// Interrupt Disable
	cpu.SetFlag(true, InterruptDisable)

	// Transfer control to interrupt vector
	ram := ram.GetRam()
	lo = *ram.Read(0xFFFE)
	hi = *ram.Read(0xFFFF)
	cpu.pc = uint16(lo) | uint16(hi)<<8
}

func branch(cpu *CPU, offset byte) {
	rel := int8(offset)

	cpu.pc = uint16(int32(cpu.pc) + int32(rel))
}

// Branch if Carry Clear
func Bcc(cpu *CPU, value *byte) {
	if cpu.GetFlag(Carry) == 0 {
		branch(cpu, *value)
	}
}

// Branch if Carry Set
func Bcs(cpu *CPU, value *byte) {
	if cpu.GetFlag(Carry) == 1 {
		branch(cpu, *value)
	}
}

// Branch if Equal
func Beq(cpu *CPU, value *byte) {
	if cpu.GetFlag(Zero) == 1 {
		branch(cpu, *value)
	}
}

// Branch if Minus
func Bmi(cpu *CPU, value *byte) {
	if cpu.GetFlag(Negative) == 1 {
		branch(cpu, *value)
	}
}

// Branch if Not Equal
func Bne(cpu *CPU, value *byte) {
	if cpu.GetFlag(Zero) == 0 {
		branch(cpu, *value)
	}
}

// Branch if Plus
func Bpl(cpu *CPU, value *byte) {
	if cpu.GetFlag(Negative) == 0 {
		branch(cpu, *value)
	}
}

// Branch if Overflow Clear
func Bvc(cpu *CPU, value *byte) {
	if cpu.GetFlag(Overflow) == 0 {
		branch(cpu, *value)
	}
}

// Branch if Overflow Set
func Bvs(cpu *CPU, value *byte) {
	if cpu.GetFlag(Overflow) == 1 {
		branch(cpu, *value)
	}
}

func compare(cpu *CPU, register byte, value byte) {
	res := register - value

	// Negative
	cpu.SetFlag(utils.GetBit(res, 7) != 0, Negative)

	// Carry
	cpu.SetFlag(register >= value, Carry)

	// Zero
	cpu.SetFlag(res == 0, Zero)
}

// Compare A
func Cmp(cpu *CPU, value *byte) {
	compare(cpu, cpu.a, *value)
}

// Compare X
func Cpx(cpu *CPU, value *byte) {
	compare(cpu, cpu.irx, *value)
}

// Compare Y
func Cpy(cpu *CPU, value *byte) {
	compare(cpu, cpu.iry, *value)
}

// No operation
func Nop(_ *CPU, _ *byte) {}

// Bitwise OR
func Ora(cpu *CPU, value *byte) {
	cpu.a |= *value

	// Zero
	cpu.SetFlag(cpu.a == 0, Zero)
	// Negative
	cpu.SetFlag(cpu.a&0x80 != 0, Negative)
}

// Bitwise XOR
func Eor(cpu *CPU, value *byte) {
	cpu.a ^= *value

	// Zero
	cpu.SetFlag(cpu.a == 0, Zero)
	// Negative
	cpu.SetFlag(cpu.a&0x80 != 0, Negative)
}

// Load A
func Lda(cpu *CPU, value *byte) {
	cpu.a = *value

	// Zero
	cpu.SetFlag(cpu.a == 0, Zero)

	// Negative
	cpu.SetFlag(cpu.a&0x80 != 0, Negative)
}

// Store A
func Sta(cpu *CPU, value *byte) {
	*value = cpu.a
}

// Load X
func Ldx(cpu *CPU, value *byte) {
	cpu.irx = *value

	// Zero
	cpu.SetFlag(cpu.irx == 0, Zero)

	// Negative
	cpu.SetFlag(cpu.irx&0x80 != 0, Negative)
}

// Store X
func Stx(cpu *CPU, value *byte) {
	*value = cpu.irx
}

// Load Y
func Ldy(cpu *CPU, value *byte) {
	cpu.iry = *value

	// Zero
	cpu.SetFlag(cpu.iry == 0, Zero)

	// Negative
	cpu.SetFlag(cpu.iry&0x80 != 0, Negative)
}

// Store Y
func Sty(cpu *CPU, value *byte) {
	*value = cpu.iry
}

// Transfer A to X
func Tax(cpu *CPU, value *byte) {
	cpu.irx = cpu.a

	cpu.SetFlag(cpu.irx == 0, Zero)

	// Negative
	cpu.SetFlag(cpu.irx&0x80 != 0, Negative)
}

// Transfer A to Y
func Tay(cpu *CPU, value *byte) {
	cpu.iry = cpu.a

	cpu.SetFlag(cpu.iry == 0, Zero)

	// Negative
	cpu.SetFlag(cpu.iry&0x80 != 0, Negative)
}

// Transfer X to A
func Txa(cpu *CPU, value *byte) {
	cpu.a = cpu.irx

	cpu.SetFlag(cpu.a == 0, Zero)

	// Negative
	cpu.SetFlag(cpu.a&0x80 != 0, Negative)
}

// Transfer Y to A
func Tya(cpu *CPU, value *byte) {
	cpu.a = cpu.iry

	cpu.SetFlag(cpu.a == 0, Zero)

	// Negative
	cpu.SetFlag(cpu.a&0x80 != 0, Negative)
}

// Transfer X to Stack Pointer
func Txs(cpu *CPU, value *byte) {
	cpu.sp = cpu.irx
}

// Transfer Stack Pointer to X
func Tsx(cpu *CPU, value *byte) {
	cpu.irx = cpu.sp

	cpu.SetFlag(cpu.irx == 0, Zero)

	// Negative
	cpu.SetFlag(cpu.irx&0x80 != 0, Negative)
}

// Push A
func Pha(cpu *CPU, value *byte) {

}

// Pull A
func Pla(cpu *CPU, value *byte) {

}

// Push Processor Status
func Php(cpu *CPU, value *byte) {

}

// Pull Processor Status
func Plp(cpu *CPU, value *byte) {

}
