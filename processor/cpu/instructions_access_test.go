package cpu

import (
	ram "NES_NEXT/processor/ram"
	"testing"
)

func TestLoadOps(t *testing.T) {
	tests := []struct {
		name             string
		loadVal          byte
		actualVal        func(*CPU) byte
		testInst         func(cpu *CPU, operand Operand)
		expectedZero     byte
		expectedNegative byte
	}{
		{
			name:      "LDX: normal load",
			loadVal:   0x55,
			actualVal: func(cpu *CPU) byte { return cpu.irx },
			testInst:  Ldx,
		},
		{
			name:         "LDX: load zero",
			loadVal:      0x00,
			actualVal:    func(cpu *CPU) byte { return cpu.irx },
			testInst:     Ldx,
			expectedZero: 1,
		},
		{
			name:             "LDX: load negative",
			loadVal:          0x80,
			actualVal:        func(cpu *CPU) byte { return cpu.irx },
			testInst:         Ldx,
			expectedNegative: 1,
		},
		{
			name:      "LDY: normal load",
			loadVal:   0x55,
			actualVal: func(cpu *CPU) byte { return cpu.iry },
			testInst:  Ldy,
		},
		{
			name:         "LDY: load zero",
			loadVal:      0x00,
			actualVal:    func(cpu *CPU) byte { return cpu.iry },
			testInst:     Ldy,
			expectedZero: 1,
		},
		{
			name:             "LDY: load negative",
			loadVal:          0x80,
			actualVal:        func(cpu *CPU) byte { return cpu.iry },
			testInst:         Ldy,
			expectedNegative: 1,
		},
		{
			name:      "LDA: normal load",
			loadVal:   0x55,
			actualVal: func(cpu *CPU) byte { return cpu.a },
			testInst:  Lda,
		},
		{
			name:         "LDA: load zero",
			loadVal:      0x00,
			actualVal:    func(cpu *CPU) byte { return cpu.a },
			testInst:     Lda,
			expectedZero: 1,
		},
		{
			name:             "LDA: load negative",
			loadVal:          0x80,
			actualVal:        func(cpu *CPU) byte { return cpu.a },
			testInst:         Lda,
			expectedNegative: 1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ram := ram.NewRAM()
			cpu := NewCPU(ram)

			loadValAddr := uint16(0x50)

			op := Operand{
				Value: &test.loadVal,
				Addr:  loadValAddr,
			}

			cpu.Mem.Write(loadValAddr, test.loadVal)

			test.testInst(cpu, op)

			got := test.actualVal(cpu)
			want := test.loadVal

			if got != want {
				t.Fatalf(
					"%s: register mismatch: got=%02X want=%02X",
					test.name,
					got,
					want,
				)
			}

			assertFlag(t, cpu, Zero, test.expectedZero)
			assertFlag(t, cpu, Negative, test.expectedNegative)

		})
	}
}

func TestStoreOps(t *testing.T) {
	tests := []struct {
		name      string
		storeVal  byte
		storeAddr uint16
		testInst  func(cpu *CPU, operand Operand)
	}{
		{
			name:      "STA",
			storeVal:  0x55,
			storeAddr: 0x0099,
			testInst:  Sta,
		},
		{
			name:      "STX",
			storeVal:  0x81,
			storeAddr: 0x00F4,
			testInst:  Stx,
		},
		{
			name:      "STY",
			storeVal:  0x71,
			storeAddr: 0x009A,
			testInst:  Sty,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ram := ram.NewRAM()
			cpu := NewCPU(ram)

			cpu.a = 0x55
			cpu.irx = 0x81
			cpu.iry = 0x71

			originalPS := byte(0xFF)
			cpu.ps = originalPS

			op := Operand{
				Addr: test.storeAddr,
			}
			test.testInst(cpu, op)

			storedValue := cpu.Mem.Read(test.storeAddr)
			if *storedValue != test.storeVal {
				t.Fatalf("%s: memory mismatch at %04X: got=%02X want=%02X",
					test.name, test.storeAddr, *storedValue, test.storeVal)
			}

			if cpu.ps != originalPS {
				t.Fatalf("%s: status register modified", test.name)
			}
		})
	}
}
