package processor

func (cpu *cpu) Clc() {
	cpu.ps.carry = false
}

func (cpu *cpu) Sec() {
	cpu.ps.carry = true
}

func (cpu *cpu) Cli() {
	cpu.ps.carry = false
}

func (cpu *cpu) Sei() {
	cpu.ps.carry = true
}

func (cpu *cpu) Clv() {
	cpu.ps.carry = false
}

func (cpu *cpu) Cld() {
	cpu.ps.carry = false
}

func (cpu *cpu) Sed() {
	cpu.ps.carry = true
}

func (cpu *cpu) Adc(value byte) {
	result := cpu.ps.carry + 
}
