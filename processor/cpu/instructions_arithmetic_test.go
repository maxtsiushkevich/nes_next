package cpu

import (
	memory "NES_NEXT/processor/memory"
	"testing"
)

func makeOperand(v byte) Operand {
	value := v
	return Operand{
		Value: &value,
	}
}

func TestAdc(t *testing.T) {
	tests := []struct {
		name             string
		a                byte
		value            byte
		carryIn          bool
		expectedA        byte
		expectedCarry    byte
		expectedZero     byte
		expectedNeg      byte
		expectedOverflow byte
	}{
		{
			name:      "ADC: simple add",
			a:         1,
			value:     2,
			expectedA: 3,
		},
		{
			name:          "ADC: carry set on overflow",
			a:             0xFF,
			value:         0x01,
			expectedA:     0x00,
			expectedCarry: 1,
			expectedZero:  1,
		},
		{
			name:      "ADC: carry input respected",
			a:         0x01,
			value:     0x01,
			carryIn:   true,
			expectedA: 0x03,
		},
		{
			name:             "ADC: signed overflow positive to negative",
			a:                0x7F,
			value:            0x01,
			expectedA:        0x80,
			expectedNeg:      1,
			expectedOverflow: 1,
		},
		{
			name:             "ADC: signed overflow negative to positive",
			a:                0x80,
			value:            0x80,
			expectedA:        0x00,
			expectedCarry:    1,
			expectedZero:     1,
			expectedOverflow: 1,
		},
		{
			name:        "ADC: negative result no overflow",
			a:           0xF0,
			value:       0x01,
			expectedA:   0xF1,
			expectedNeg: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ram := memory.NewRAM()
			cpu := NewCPU(ram)

			cpu.a = tt.a

			if tt.carryIn {
				cpu.SetFlag(true, Carry)
			}

			Adc(cpu, makeOperand(tt.value))

			if cpu.a != tt.expectedA {
				t.Fatalf("A = %02X, want %02X", cpu.a, tt.expectedA)
			}

			assertFlag(t, cpu, Carry, tt.expectedCarry)
			assertFlag(t, cpu, Zero, tt.expectedZero)
			assertFlag(t, cpu, Negative, tt.expectedNeg)
			assertFlag(t, cpu, Overflow, tt.expectedOverflow)
		})
	}
}

func TestSbc(t *testing.T) {
	tests := []struct {
		name             string
		a                byte
		value            byte
		carryIn          bool
		expectedA        byte
		expectedCarry    byte
		expectedZero     byte
		expectedNeg      byte
		expectedOverflow byte
	}{
		{
			name:          "SBC: simple subtract",
			a:             5,
			value:         3,
			carryIn:       true,
			expectedA:     2,
			expectedCarry: 1,
		},
		{
			name:        "SBC: borrow required",
			a:           3,
			value:       5,
			carryIn:     true,
			expectedA:   0xFE,
			expectedNeg: 1,
		},
		{
			name:          "SBC: subtract to zero",
			a:             5,
			value:         5,
			carryIn:       true,
			expectedA:     0,
			expectedCarry: 1,
			expectedZero:  1,
		},
		{
			name:             "SBC: signed overflow",
			a:                0x80,
			value:            0x01,
			carryIn:          true,
			expectedA:        0x7F,
			expectedCarry:    1,
			expectedOverflow: 1,
		},
		{
			name:          "SBC: without carry subtract extra one",
			a:             5,
			value:         1,
			carryIn:       false,
			expectedA:     3,
			expectedCarry: 1,
		},
		{
			name:        "SBC: underflow wrap around",
			a:           0x00,
			value:       0x01,
			carryIn:     true,
			expectedA:   0xFF,
			expectedNeg: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ram := memory.NewRAM()
			cpu := NewCPU(ram)

			cpu.a = tt.a

			if tt.carryIn {
				cpu.SetFlag(true, Carry)
			}

			Sbc(cpu, makeOperand(tt.value))

			if cpu.a != tt.expectedA {
				t.Fatalf("A = %02X, want %02X", cpu.a, tt.expectedA)
			}

			assertFlag(t, cpu, Carry, tt.expectedCarry)
			assertFlag(t, cpu, Zero, tt.expectedZero)
			assertFlag(t, cpu, Negative, tt.expectedNeg)
			assertFlag(t, cpu, Overflow, tt.expectedOverflow)
		})
	}
}

func TestInc(t *testing.T) {
	// mem := ram.GetRam()

	tests := []struct {
		name         string
		initial      byte
		expected     byte
		expectedZero byte
		expectedNeg  byte
	}{
		{
			name:     "INC: normal increment",
			initial:  1,
			expected: 2,
		},
		{
			name:         "INC: overflow to zero",
			initial:      0xFF,
			expected:     0x00,
			expectedZero: 1,
		},
		{
			name:        "INC: becomes negative",
			initial:     0x7F,
			expected:    0x80,
			expectedNeg: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ram := memory.NewRAM()
			cpu := NewCPU(ram)

			addr := uint16(0x200)
			cpu.Mem.Write(addr, tt.initial)

			Inc(cpu, Operand{Addr: addr})

			got := *cpu.Mem.Read(addr)
			if got != tt.expected {
				t.Fatalf("memory = %02X, want %02X", got, tt.expected)
			}

			assertFlag(t, cpu, Zero, tt.expectedZero)
			assertFlag(t, cpu, Negative, tt.expectedNeg)
		})
	}
}

func TestDec(t *testing.T) {

	tests := []struct {
		name         string
		initial      byte
		expected     byte
		expectedZero byte
		expectedNeg  byte
	}{
		{
			name:     "DEC: normal decrement",
			initial:  2,
			expected: 1,
		},
		{
			name:         "DEC: decrement to zero",
			initial:      1,
			expected:     0,
			expectedZero: 1,
		},
		{
			name:        "DEC: underflow wrap",
			initial:     0x00,
			expected:    0xFF,
			expectedNeg: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ram := memory.NewRAM()
			cpu := NewCPU(ram)

			addr := uint16(0x200)
			cpu.Mem.Write(addr, tt.initial)

			Dec(cpu, Operand{Addr: addr})

			got := *cpu.Mem.Read(addr)
			if got != tt.expected {
				t.Fatalf("memory = %02X, want %02X", got, tt.expected)
			}

			assertFlag(t, cpu, Zero, tt.expectedZero)
			assertFlag(t, cpu, Negative, tt.expectedNeg)
		})
	}
}

func TestInxDex(t *testing.T) {
	t.Run("INX: wrap to zero", func(t *testing.T) {
		ram := memory.NewRAM()
		cpu := NewCPU(ram)
		cpu.irx = 0xFF

		Inx(cpu, Operand{})

		if cpu.irx != 0 {
			t.Fatalf("X = %02X, want 00", cpu.irx)
		}

		assertFlag(t, cpu, Zero, 1)
	})

	t.Run("INX: negative", func(t *testing.T) {
		ram := memory.NewRAM()
		cpu := NewCPU(ram)
		cpu.irx = 0x7F

		Inx(cpu, Operand{})

		if cpu.irx != 0x80 {
			t.Fatalf("X = %02X, want 80", cpu.irx)
		}

		assertFlag(t, cpu, Negative, 1)
	})

	t.Run("DEX: underflow", func(t *testing.T) {
		ram := memory.NewRAM()
		cpu := NewCPU(ram)
		cpu.irx = 0

		Dex(cpu, Operand{})

		if cpu.irx != 0xFF {
			t.Fatalf("X = %02X, want FF", cpu.irx)
		}

		assertFlag(t, cpu, Negative, 1)
	})
}

func TestInyDey(t *testing.T) {
	t.Run("INY: wrap to zero", func(t *testing.T) {
		ram := memory.NewRAM()
		cpu := NewCPU(ram)
		cpu.iry = 0xFF

		Iny(cpu, Operand{})

		if cpu.iry != 0 {
			t.Fatalf("Y = %02X, want 00", cpu.iry)
		}

		assertFlag(t, cpu, Zero, 1)
	})

	t.Run("DEY: underflow", func(t *testing.T) {
		ram := memory.NewRAM()
		cpu := NewCPU(ram)
		cpu.iry = 0

		Dey(cpu, Operand{})

		if cpu.iry != 0xFF {
			t.Fatalf("Y = %02X, want FF", cpu.iry)
		}

		assertFlag(t, cpu, Negative, 1)
	})
}
