package cpu

import (
	ram "NES_NEXT/processor/ram"
	"testing"
)

func TestStackTransferAndFlags(t *testing.T) {
	testCases := []struct {
		name string
		exec func(cpu *CPU, op Operand)

		setup func(cpu *CPU)

		check func(t *testing.T, cpu *CPU)
	}{
		{
			name: "TXS transfers X to SP",
			exec: Txs,
			setup: func(cpu *CPU) {
				cpu.irx = 0x42
			},
			check: func(t *testing.T, cpu *CPU) {
				if cpu.sp != 0x42 {
					t.Fatalf("SP=%02X want 42", cpu.sp)
				}
			},
		},
		{
			name: "TSX transfers SP to X and sets flags",
			exec: Tsx,
			setup: func(cpu *CPU) {
				cpu.sp = 0x80
			},
			check: func(t *testing.T, cpu *CPU) {
				if cpu.irx != 0x80 {
					t.Fatalf("X=%02X want 80", cpu.irx)
				}
				assertFlag(t, cpu, Negative, 1)
			},
		},
		{
			name: "TSX zero flag",
			exec: Tsx,
			setup: func(cpu *CPU) {
				cpu.sp = 0x00
			},
			check: func(t *testing.T, cpu *CPU) {
				assertFlag(t, cpu, Zero, 1)
				assertFlag(t, cpu, Negative, 0)
			},
		},

		{
			name: "PHA pushes A to stack",
			exec: Pha,
			setup: func(cpu *CPU) {
				cpu.a = 0x99
			},
			check: func(t *testing.T, cpu *CPU) {
				val := cpu.pop()
				if val != 0x99 {
					t.Fatalf("stack=%02X want 99", val)
				}
			},
		},
		{
			name: "PLA pulls A and sets flags",
			exec: Pla,
			setup: func(cpu *CPU) {
				cpu.push(0x80)
			},
			check: func(t *testing.T, cpu *CPU) {
				if cpu.a != 0x80 {
					t.Fatalf("A=%02X want 80", cpu.a)
				}
				assertFlag(t, cpu, Negative, 1)
			},
		},
		{
			name: "PLA zero flag",
			exec: Pla,
			setup: func(cpu *CPU) {
				cpu.push(0x00)
			},
			check: func(t *testing.T, cpu *CPU) {
				assertFlag(t, cpu, Zero, 1)
			},
		},
		{
			name: "PHP pushes status with bit5 and bit4 set",
			exec: Php,
			setup: func(cpu *CPU) {
				cpu.ps = 0b00000000
			},
			check: func(t *testing.T, cpu *CPU) {
				val := cpu.pop()

				if (val & (1 << 5)) == 0 {
					t.Fatalf("bit5 not set")
				}
				if (val & (1 << 4)) == 0 {
					t.Fatalf("bit4 not set")
				}
			},
		},
		{
			name: "PLP restores status and clears bit4",
			exec: Plp,
			setup: func(cpu *CPU) {
				// simulate stack value
				cpu.push(0b11111111)
			},
			check: func(t *testing.T, cpu *CPU) {
				if cpu.GetFlag(Break) != 0 {
					t.Fatalf("break flag must be cleared")
				}
				if cpu.GetFlag(5) != 1 {
					t.Fatalf("bit5 must be set")
				}
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ram := ram.NewRAM()
			cpu := NewCPU(ram)

			if tc.setup != nil {
				tc.setup(cpu)
			}

			tc.exec(cpu, Operand{})

			if tc.check != nil {
				tc.check(t, cpu)
			}
		})
	}
}
