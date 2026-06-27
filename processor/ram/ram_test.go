package memory

import "testing"

func TestRAM_ReadWrite(t *testing.T) {
	mem := NewRAM()

	addr := uint16(0x1234)
	value := byte(0xAB)

	mem.Write(addr, value)

	got := *mem.Read(addr)

	if got != value {
		t.Fatalf("expected %#x, got %#x", value, got)
	}
}

func TestRAM_Overwrite(t *testing.T) {
	mem := NewRAM()

	addr := uint16(0x2000)

	mem.Write(addr, 0x11)
	mem.Write(addr, 0x22)

	if *mem.Read(addr) != 0x22 {
		t.Fatalf("expected 0x22")
	}
}

func TestRAM_WriteBlock(t *testing.T) {
	mem := NewRAM()

	data := []byte{1, 2, 3, 4, 5}

	mem.WriteBlock(0x1000, data)

	for i, v := range data {
		if *mem.Read(uint16(0x1000 + uint16(i))) != v {
			t.Fatalf("mismatch at %d", i)
		}
	}
}

func TestRAM_WriteBlock_Bounds(t *testing.T) {
	mem := NewRAM()

	data := []byte{1, 2, 3}

	mem.WriteBlock(0xFFFE, data)

	if *mem.Read(0xFFFE) != 1 {
		t.Fatalf("expected 1 at 0xFFFE")
	}
}
