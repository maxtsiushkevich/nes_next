package cpu

// Arithmetic Shift Left
func Asl(cpu *CPU, op Operand) {
	if op.Value == nil {
		// accumulator mode
		cpu.SetFlag(cpu.a&0x80 != 0, Carry)

		cpu.a <<= 1

		cpu.SetFlag(cpu.a == 0, Zero)
		cpu.SetFlag(cpu.a&0x80 != 0, Negative)
		return
	}

	// memory mode
	val := *cpu.Mem.Read(op.Addr)

	cpu.SetFlag(val&0x80 != 0, Carry)

	val <<= 1
	cpu.Mem.Write(op.Addr, val)

	cpu.SetFlag(val == 0, Zero)
	cpu.SetFlag(val&0x80 != 0, Negative)
}

// Rotate Left
func Rol(cpu *CPU, op Operand) {
	if op.Value == nil {
		// accumulator mode
		oldCarry := cpu.GetFlag(Carry)

		cpu.SetFlag(cpu.a&0x80 != 0, Carry)

		cpu.a = (cpu.a << 1) | oldCarry

		cpu.SetFlag(cpu.a == 0, Zero)
		cpu.SetFlag(cpu.a&0x80 != 0, Negative)
		return
	}

	// memory mode
	val := *cpu.Mem.Read(op.Addr)

	oldCarry := cpu.GetFlag(Carry)

	cpu.SetFlag(val&0x80 != 0, Carry)

	val = (val << 1) | oldCarry
	cpu.Mem.Write(op.Addr, val)

	cpu.SetFlag(val == 0, Zero)
	cpu.SetFlag(val&0x80 != 0, Negative)
}

// Rotate Right
func Ror(cpu *CPU, op Operand) {
	if op.Value == nil {
		// accumulator mode
		oldCarry := cpu.GetFlag(Carry)

		cpu.SetFlag(cpu.a&0x01 != 0, Carry)

		cpu.a = (cpu.a >> 1) | (oldCarry << 7)

		cpu.SetFlag(cpu.a == 0, Zero)
		cpu.SetFlag(cpu.a&0x80 != 0, Negative)
		return
	}

	// memory mode
	val := *cpu.Mem.Read(op.Addr)

	oldCarry := cpu.GetFlag(Carry)

	cpu.SetFlag(val&0x01 != 0, Carry)

	val = (val >> 1) | (oldCarry << 7)
	cpu.Mem.Write(op.Addr, val)

	cpu.SetFlag(val == 0, Zero)
	cpu.SetFlag(val&0x80 != 0, Negative)
}

// Logical Shift Right
func Lsr(cpu *CPU, op Operand) {
	if op.Value == nil {
		// accumulator mode
		cpu.SetFlag(cpu.a&0x01 != 0, Carry)

		cpu.a >>= 1

		cpu.SetFlag(cpu.a == 0, Zero)
		cpu.SetFlag(false, Negative)
		return
	}

	// memory mode
	val := *cpu.Mem.Read(op.Addr)

	cpu.SetFlag(val&0x01 != 0, Carry)

	val >>= 1
	cpu.Mem.Write(op.Addr, val)

	cpu.SetFlag(val == 0, Zero)
	cpu.SetFlag(false, Negative)
}
