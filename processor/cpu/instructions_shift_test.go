package cpu

import (
	ram "NES_NEXT/processor/ram"
	"testing"
)

func TestShiftInstructions(t *testing.T) {
	testCases := []struct {
		name             string
		exec             func(cpu *CPU, op Operand)
		input            byte
		carryIn          byte
		expectedValue    byte
		expectedCarry    byte
		expectedZero     byte
		expectedNegative byte
	}{
		{
			name:             "ASL shifts left, sets carry from bit7",
			exec:             Asl,
			input:            0b10000100,
			expectedValue:    0b00001000,
			expectedCarry:    1,
			expectedZero:     0,
			expectedNegative: 0,
		},
		{
			name:             "ASL results zero",
			exec:             Asl,
			input:            0,
			expectedValue:    0,
			expectedCarry:    0,
			expectedZero:     1,
			expectedNegative: 0,
		},
		{
			name:             "LSR shifts right, carry from bit0",
			exec:             Lsr,
			input:            0b00000011,
			expectedValue:    0b00000001,
			expectedCarry:    1,
			expectedZero:     0,
			expectedNegative: 0,
		},
		{
			name:             "LSR clears negative always",
			exec:             Lsr,
			input:            0b10000001,
			expectedValue:    0b01000000,
			expectedCarry:    1,
			expectedZero:     0,
			expectedNegative: 0,
		},
		{
			name:             "ROL rotates left with carry in",
			exec:             Rol,
			input:            0b01000000,
			carryIn:          1,
			expectedValue:    0b10000001,
			expectedCarry:    0,
			expectedZero:     0,
			expectedNegative: 1,
		},
		{
			name:             "ROL zero result",
			exec:             Rol,
			input:            0,
			carryIn:          0,
			expectedValue:    0,
			expectedCarry:    0,
			expectedZero:     1,
			expectedNegative: 0,
		},
		{
			name:             "ROR rotates right with carry in",
			exec:             Ror,
			input:            0b00000001,
			carryIn:          1,
			expectedValue:    0b10000000,
			expectedCarry:    1,
			expectedZero:     0,
			expectedNegative: 1,
		},
		{
			name:             "ROR zero result",
			exec:             Ror,
			input:            0,
			carryIn:          0,
			expectedValue:    0,
			expectedCarry:    0,
			expectedZero:     1,
			expectedNegative: 0,
		},
	}

	for _, tc := range testCases {

		t.Run(tc.name, func(t *testing.T) {
			ram := ram.NewRAM()
			cpu := NewCPU(ram)

			value := tc.input
			valueAddr := uint16(0x55)
			cpu.Mem.Write(valueAddr, value)

			cpu.SetFlag(tc.carryIn == 1, Carry)

			tc.exec(cpu, Operand{
				Value: &value,
				Addr:  valueAddr,
			})

			value = *cpu.Mem.Read(valueAddr)

			if value != tc.expectedValue {
				t.Fatalf("Value=%08b want=%08b", value, tc.expectedValue)
			}

			assertFlag(t, cpu, Carry, tc.expectedCarry)
			assertFlag(t, cpu, Zero, tc.expectedZero)
			assertFlag(t, cpu, Negative, tc.expectedNegative)
		})
	}
}
