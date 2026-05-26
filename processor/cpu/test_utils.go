package cpu

import (
	"testing"
)

func (f CpuFlag) String() string {
	switch f {
	case Negative:
		return "Negative"
	case Overflow:
		return "Overflow"
	case Break:
		return "Break"
	case Decimal:
		return "Decimal"
	case InterruptDisable:
		return "InterruptDisable"
	case Zero:
		return "Zero"
	case Carry:
		return "Carry"
	default:
		return "UnknownFlag"
	}
}

func assertFlag(t *testing.T, cpu *CPU, flag CpuFlag, expected byte) {
	t.Helper()

	got := cpu.GetFlag(flag)
	if got != expected {
		t.Fatalf("flag %s = %d, want %d", flag, got, expected)
	}
}
