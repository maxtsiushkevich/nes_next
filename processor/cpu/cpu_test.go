package processor

import "testing"

func TestSetFlag(t *testing.T) {
	cpu := GetCPU()

	cpu.SetFlag(true, Carry)

	if cpu.GetFlag(Carry) != 1 {
		t.Errorf("expected Carry flag = 1")
	}

	cpu.SetFlag(false, Carry)

	if cpu.GetFlag(Carry) != 0 {
		t.Errorf("expected Carry flag = 0")
	}
}

func TestGetFlag(t *testing.T) {
	cpu := GetCPU()

	if cpu.GetFlag(Zero) != 0 {
		t.Errorf("expected Zero flag = 0")
	}

	cpu.SetFlag(true, Negative)

	if cpu.GetFlag(Negative) != 1 {
		t.Errorf("expected Negative flag = 1")
	}
}
