package utils

func GetBit(value byte, position byte) byte {
	if value&(1<<position) != 0 {
		return 1
	}
	return 0
}
