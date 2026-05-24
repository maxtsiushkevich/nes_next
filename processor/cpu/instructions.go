package processor

import (
	ram "NES_NEXT/processor/memory"
	"NES_NEXT/utils"
)

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

// Bitwise And
func And(cpu *CPU, op Operand) {
	cpu.a &= *op.Value

	cpu.SetFlag(cpu.a == 0, Zero)
	cpu.SetFlag(cpu.a&0x80 != 0, Negative)
}

// Arithmetic Shift Left
func Asl(cpu *CPU, op Operand) {
	cpu.SetFlag(*op.Value&0x80 != 0, Carry)
	*op.Value = *op.Value << 1

	cpu.SetFlag(*op.Value == 0, Zero)
	cpu.SetFlag(*op.Value&0x80 != 0, Negative)
}

func Bit(cpu *CPU, op Operand) {
	result := cpu.a & *op.Value

	cpu.SetFlag(result == 0, Zero)
	cpu.SetFlag(utils.GetBit(*op.Value, 6) != 0, Overflow)
	cpu.SetFlag(utils.GetBit(*op.Value, 7) != 0, Negative)
}

// Break
func Brk(cpu *CPU, _ Operand) {
	cpu.Interrupt(BRK)
}

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

// No operation
func Nop(_ *CPU, _ Operand) {}

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

// Load A
func Lda(cpu *CPU, op Operand) {
	cpu.a = *op.Value

	cpu.SetFlag(cpu.a == 0, Zero)
	cpu.SetFlag(cpu.a&0x80 != 0, Negative)
}

// Store A
func Sta(cpu *CPU, op Operand) {
	ram := ram.GetRam()
	ram.Write(op.Addr, cpu.a)
}

// Load X
func Ldx(cpu *CPU, op Operand) {
	cpu.irx = *op.Value

	cpu.SetFlag(cpu.irx == 0, Zero)
	cpu.SetFlag(cpu.irx&0x80 != 0, Negative)
}

// Store X
func Stx(cpu *CPU, op Operand) {
	ram := ram.GetRam()
	ram.Write(op.Addr, cpu.irx)
}

// Load Y
func Ldy(cpu *CPU, op Operand) {
	cpu.iry = *op.Value

	cpu.SetFlag(cpu.iry == 0, Zero)
	cpu.SetFlag(cpu.iry&0x80 != 0, Negative)
}

// Store Y
func Sty(cpu *CPU, op Operand) {
	ram := ram.GetRam()
	ram.Write(op.Addr, cpu.iry)
}

// Transfer A to X
func Tax(cpu *CPU, _ Operand) {
	cpu.irx = cpu.a

	cpu.SetFlag(cpu.irx == 0, Zero)
	cpu.SetFlag(cpu.irx&0x80 != 0, Negative)
}

// Transfer A to Y
func Tay(cpu *CPU, _ Operand) {
	cpu.iry = cpu.a

	cpu.SetFlag(cpu.iry == 0, Zero)
	cpu.SetFlag(cpu.iry&0x80 != 0, Negative)
}

// Transfer X to A
func Txa(cpu *CPU, _ Operand) {
	cpu.a = cpu.irx

	cpu.SetFlag(cpu.a == 0, Zero)
	cpu.SetFlag(cpu.a&0x80 != 0, Negative)
}

// Transfer Y to A
func Tya(cpu *CPU, _ Operand) {
	cpu.a = cpu.iry

	cpu.SetFlag(cpu.a == 0, Zero)
	cpu.SetFlag(cpu.a&0x80 != 0, Negative)
}

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

// Increment Memory
func Inc(cpu *CPU, op Operand) {
	ram := ram.GetRam()
	result := *ram.Read(op.Addr)
	result++

	cpu.SetFlag(result == 0, Zero)
	cpu.SetFlag(result&0x80 != 0, Negative)

	ram.Write(op.Addr, result)
}

// Decrement Memory
func Dec(cpu *CPU, op Operand) {
	ram := ram.GetRam()
	result := *ram.Read(op.Addr)
	result--

	cpu.SetFlag(result == 0, Zero)
	cpu.SetFlag(result&0x80 != 0, Negative)

	ram.Write(op.Addr, result)
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
