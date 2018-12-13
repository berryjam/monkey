package monkey

// Assembles a jump to a function value
func jmpToFunctionValue(to uintptr) []byte {
	return []byte{
		0x48, 0xC7, 0xC2,
		byte(to),
		byte(to >> 8),
		byte(to >> 16),
		byte(to >> 24), // MOV rdx, to
		0xFF, 0xE2,     // JMP rdx
	}
}
