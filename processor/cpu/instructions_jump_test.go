package cpu

import (
	ram "NES_NEXT/processor/ram"
	"testing"
)

func TestJumpAndBreak(t *testing.T) {
	testCases := []struct {
		name     string
		exec     func(cpu *CPU, op Operand)
		addr     uint16
		startPC  uint16
		expected uint16
	}{
		{
			name:     "JMP sets PC",
			exec:     Jmp,
			addr:     0x8000,
			startPC:  0x1234,
			expected: 0x8000,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ram := ram.NewRAM()
			cpu := NewCPU(ram)

			cpu.pc = tc.startPC

			tc.exec(cpu, Operand{
				Addr: tc.addr,
			})

			if cpu.pc != tc.expected {
				t.Fatalf("PC=%04X want %04X", cpu.pc, tc.expected)
			}
		})
	}
}

func TestJSR(t *testing.T) {
	ram := ram.NewRAM()
	cpu := NewCPU(ram)

	cpu.pc = 0x1234

	target := uint16(0x8000)

	Jsr(cpu, Operand{
		Addr: target,
	})

	// PC should jump
	if cpu.pc != target {
		t.Fatalf("PC=%04X want %04X", cpu.pc, target)
	}

	// stack should contain return address (pc - 1)
	// pop order: lo then hi (because LIFO reverse)
	lo := cpu.pop()
	hi := cpu.pop()

	ret := uint16(lo) | uint16(hi)<<8

	if ret != 0x1233 {
		t.Fatalf("return addr=%04X want %04X", ret, 0x1233)
	}
}

func TestRTS(t *testing.T) {
	ram := ram.NewRAM()
	cpu := NewCPU(ram)

	// simulate stack: return address = 0x1234
	cpu.push(0x12)
	cpu.push(0x34)

	Rts(cpu, Operand{})

	// RTS adds +1
	if cpu.pc != 0x1235 {
		t.Fatalf("PC=%04X want %04X", cpu.pc, 0x1235)
	}
}

func TestRTI(t *testing.T) {
	ram := ram.NewRAM()
	cpu := NewCPU(ram)

	// simulate status register
	cpu.ps = 0x00

	// push PC = 0x2000
	cpu.push(0x20)
	cpu.push(0x00)

	// push PS
	cpu.push(0xFF)

	cpu.pc = 0x1111

	Rti(cpu, Operand{})

	// PC restored
	if cpu.pc != 0x2000 {
		t.Fatalf("PC=%04X want %04X", cpu.pc, 0x2000)
	}

	// bit5 must be set, bit4 cleared
	if cpu.GetFlag(5) != 1 {
		t.Fatalf("bit5 not set")
	}
}
