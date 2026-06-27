package cpu

import (
	ram "NES_NEXT/processor/ram"
	"testing"
)

func TestLogicalInstructions(t *testing.T) {
	testCases := []struct {
		name             string
		exec             func(cpu *CPU, op Operand)
		a                byte
		value            byte
		expectedA        byte
		expectedZero     byte
		expectedNegative byte
	}{
		{
			name:             "ORA sets bits",
			exec:             Ora,
			a:                0b00001111,
			value:            0b11110000,
			expectedA:        0b11111111,
			expectedZero:     0,
			expectedNegative: 1,
		},
		{
			name:             "ORA produces zero",
			exec:             Ora,
			a:                0,
			value:            0,
			expectedA:        0,
			expectedZero:     1,
			expectedNegative: 0,
		},
		{
			name:             "EOR xor result",
			exec:             Eor,
			a:                0b11110000,
			value:            0b11110000,
			expectedA:        0,
			expectedZero:     1,
			expectedNegative: 0,
		},
		{
			name:             "EOR negative result",
			exec:             Eor,
			a:                0b10000000,
			value:            0,
			expectedA:        0b10000000,
			expectedZero:     0,
			expectedNegative: 1,
		},
		{
			name:             "AND keeps matching bits",
			exec:             And,
			a:                0b11110000,
			value:            0b10101010,
			expectedA:        0b10100000,
			expectedZero:     0,
			expectedNegative: 1,
		},
		{
			name:             "AND produces zero",
			exec:             And,
			a:                0b11110000,
			value:            0b00001111,
			expectedA:        0,
			expectedZero:     1,
			expectedNegative: 0,
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			ram := ram.NewRAM()
			cpu := NewCPU(ram)

			cpu.a = tc.a

			value := tc.value

			tc.exec(cpu, Operand{
				Value: &value,
			})

			if cpu.a != tc.expectedA {
				t.Fatalf(
					"A = %08b, want %08b",
					cpu.a,
					tc.expectedA,
				)
			}

			assertFlag(t, cpu, Zero, tc.expectedZero)
			assertFlag(t, cpu, Negative, tc.expectedNegative)
		})
	}
}

func TestBit(t *testing.T) {
	testCases := []struct {
		name             string
		a                byte
		value            byte
		expectedZero     byte
		expectedOverflow byte
		expectedNegative byte
	}{
		{
			name:             "BIT overlap sets zero false",
			a:                0b00001111,
			value:            0b00000001,
			expectedZero:     0,
			expectedOverflow: 0,
			expectedNegative: 0,
		},
		{
			name:             "BIT no overlap sets zero true",
			a:                0b00001111,
			value:            0b11110000,
			expectedZero:     1,
			expectedOverflow: 1,
			expectedNegative: 1,
		},
		{
			name:             "BIT sets overflow from bit6",
			a:                0xFF,
			value:            0b01000000,
			expectedZero:     0,
			expectedOverflow: 1,
			expectedNegative: 0,
		},
		{
			name:             "BIT sets negative from bit7",
			a:                0xFF,
			value:            0b10000000,
			expectedZero:     0,
			expectedOverflow: 0,
			expectedNegative: 1,
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			ram := ram.NewRAM()
			cpu := NewCPU(ram)

			cpu.a = tc.a
			initialA := cpu.a

			value := tc.value

			Bit(cpu, Operand{
				Value: &value,
			})

			if cpu.a != initialA {
				t.Fatalf(
					"BIT modified A: got %08b, want %08b",
					cpu.a,
					initialA,
				)
			}

			assertFlag(t, cpu, Zero, tc.expectedZero)
			assertFlag(t, cpu, Overflow, tc.expectedOverflow)
			assertFlag(t, cpu, Negative, tc.expectedNegative)
		})
	}
}
