package cpu

import (
	ram "NES_NEXT/processor/ram"
	"testing"
)

func TestTransferInstructions(t *testing.T) {
	testCases := []struct {
		name string
		exec func(cpu *CPU, op Operand)

		a byte
		x byte
		y byte

		expectedA byte
		expectedX byte
		expectedY byte

		checkA bool
		checkX bool
		checkY bool
	}{
		{
			name:      "TAX transfers A to X",
			exec:      Tax,
			a:         0x42,
			expectedX: 0x42,
			checkX:    true,
		},
		{
			name:      "TAX sets zero flag",
			exec:      Tax,
			a:         0x00,
			expectedX: 0x00,
			checkX:    true,
		},

		{
			name:      "TAY transfers A to Y",
			exec:      Tay,
			a:         0x55,
			expectedY: 0x55,
			checkY:    true,
		},
		{
			name:      "TXA transfers X to A",
			exec:      Txa,
			x:         0x99,
			expectedA: 0x99,
			checkA:    true,
		},
		{
			name:      "TXA zero flag",
			exec:      Txa,
			x:         0x00,
			expectedA: 0x00,
			checkA:    true,
		},
		{
			name:      "TYA transfers Y to A",
			exec:      Tya,
			y:         0x80,
			expectedA: 0x80,
			checkA:    true,
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			ram := ram.NewRAM()
			cpu := NewCPU(ram)

			cpu.a = tc.a
			cpu.irx = tc.x
			cpu.iry = tc.y

			tc.exec(cpu, Operand{})

			if tc.checkA && cpu.a != tc.expectedA {
				t.Fatalf("A=%02X want=%02X", cpu.a, tc.expectedA)
			}

			if tc.checkX && cpu.irx != tc.expectedX {
				t.Fatalf("X=%02X want=%02X", cpu.irx, tc.expectedX)
			}

			if tc.checkY && cpu.iry != tc.expectedY {
				t.Fatalf("Y=%02X want=%02X", cpu.iry, tc.expectedY)
			}
		})
	}
}
