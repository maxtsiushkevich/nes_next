package cpu

import (
	ram "NES_NEXT/processor/ram"
	"testing"
)

func TestCompare(t *testing.T) {
	testCases := []struct {
		name             string
		register         byte
		value            byte
		expectedCarry    byte
		expectedZero     byte
		expectedNegative byte
	}{
		{
			name:             "Compare: Register >= value",
			register:         10,
			value:            5,
			expectedCarry:    1,
			expectedZero:     0,
			expectedNegative: 0,
		},
		{
			name:             "Compare: Register == value",
			register:         25,
			value:            25,
			expectedCarry:    1,
			expectedZero:     1,
			expectedNegative: 0,
		},
		{
			name:             "Compare: Register < value",
			register:         10,
			value:            25,
			expectedCarry:    0,
			expectedZero:     0,
			expectedNegative: 1,
		},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			ram := ram.NewRAM()
			cpu := NewCPU(ram)

			compare(cpu, test.register, test.value)

			assertFlag(t, cpu, Carry, test.expectedCarry)
			assertFlag(t, cpu, Zero, test.expectedZero)
			assertFlag(t, cpu, Negative, test.expectedNegative)

		})
	}
}

func TestCompareUseCorrectRegister(t *testing.T) {
	testCases := []struct {
		name string
		exec func(cpu *CPU, op Operand)
		a    byte
		x    byte
		y    byte
		want byte
	}{
		{
			name: "CMP: accumulator",
			exec: Cmp,
			a:    10,
			x:    50,
			y:    100,
			want: 10,
		},
		{
			name: "CPX: X register",
			exec: Cpx,
			a:    10,
			x:    50,
			y:    100,
			want: 50,
		},
		{
			name: "CPY: Y register",
			exec: Cpy,
			a:    10,
			x:    50,
			y:    100,
			want: 100,
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

			value := tc.want

			tc.exec(cpu, Operand{
				Value: &value,
			})

			assertFlag(t, cpu, Zero, 1)
		})
	}
}
