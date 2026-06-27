package cpu

import (
	ram "NES_NEXT/processor/ram"
	"testing"
)

func TestBranchInstructions(t *testing.T) {
	testCases := []struct {
		name      string
		exec      func(cpu *CPU, op Operand)
		flag      CpuFlag
		flagValue byte
		offset    byte
		startPC   uint16
		expected  uint16
	}{
		{
			name:      "BCC taken when carry clear",
			exec:      Bcc,
			flag:      Carry,
			flagValue: 0,
			offset:    0x10,
			startPC:   0x1000,
			expected:  0x1010,
		},
		{
			name:      "BCC not taken when carry set",
			exec:      Bcc,
			flag:      Carry,
			flagValue: 1,
			offset:    0x10,
			startPC:   0x1000,
			expected:  0x1000,
		},
		{
			name:      "BEQ taken when zero set",
			exec:      Beq,
			flag:      Zero,
			flagValue: 1,
			offset:    0x05,
			startPC:   0x2000,
			expected:  0x2005,
		},
		{
			name:      "BNE taken when zero clear",
			exec:      Bne,
			flag:      Zero,
			flagValue: 0,
			offset:    0x05,
			startPC:   0x2000,
			expected:  0x2005,
		},
		{
			name:      "BMI taken when negative set",
			exec:      Bmi,
			flag:      Negative,
			flagValue: 1,
			offset:    0x01,
			startPC:   0x3000,
			expected:  0x3001,
		},
		{
			name:      "BPL not taken when negative set",
			exec:      Bpl,
			flag:      Negative,
			flagValue: 1,
			offset:    0x01,
			startPC:   0x3000,
			expected:  0x3000,
		},
		{
			name:      "BVC taken when overflow clear",
			exec:      Bvc,
			flag:      Overflow,
			flagValue: 0,
			offset:    0x02,
			startPC:   0x4000,
			expected:  0x4002,
		},
		{
			name:      "BVS taken when overflow set",
			exec:      Bvs,
			flag:      Overflow,
			flagValue: 1,
			offset:    0x02,
			startPC:   0x4000,
			expected:  0x4002,
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			ram := ram.NewRAM()
			cpu := NewCPU(ram)

			cpu.pc = tc.startPC

			cpu.SetFlag(tc.flagValue == 1, tc.flag)

			value := tc.offset

			tc.exec(cpu, Operand{
				Value: &value,
			})

			if cpu.pc != tc.expected {
				t.Fatalf(
					"PC = %04X, want %04X",
					cpu.pc,
					tc.expected,
				)
			}
		})
	}
}

func TestBranchWrapAround(t *testing.T) {
	ram := ram.NewRAM()
	cpu := NewCPU(ram)

	cpu.pc = 0xFFF0

	cpu.SetFlag(true, Zero)

	offset := byte(0x20) // +32 => 0xFFF0 + 0x20 = 0x0010 (wrap)

	Beq(cpu, Operand{
		Value: &offset,
	})

	expected := uint16(0x0010)

	if cpu.pc != expected {
		t.Fatalf(
			"PC wrap failed: got %04X, want %04X",
			cpu.pc,
			expected,
		)
	}
}
