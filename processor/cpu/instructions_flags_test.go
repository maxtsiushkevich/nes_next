package cpu

import (
	ram "NES_NEXT/processor/ram"
	"testing"
)

func TestFlagInstructions(t *testing.T) {
	testCases := []struct {
		name     string
		exec     func(cpu *CPU, op Operand)
		flag     CpuFlag
		expected byte
	}{
		{"CLC clears carry", Clc, Carry, 0},
		{"SEC sets carry", Sec, Carry, 1},
		{"CLI clears interrupt disable", Cli, InterruptDisable, 0},
		{"SEI sets interrupt disable", Sei, InterruptDisable, 1},
		{"CLV clears overflow", Clv, Overflow, 0},
		{"CLD clears decimal", Cld, Decimal, 0},
		{"SED sets decimal", Sed, Decimal, 1},
	}

	for _, tc := range testCases {

		t.Run(tc.name, func(t *testing.T) {
			ram := ram.NewRAM()
			cpu := NewCPU(ram)

			cpu.SetFlag(tc.expected == 0, tc.flag)

			tc.exec(cpu, Operand{})

			assertFlag(t, cpu, tc.flag, tc.expected)
		})
	}
}
