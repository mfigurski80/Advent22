package utils

/**
 * BitVec is a bitmap of 64 bits. Maps uint below 64 to a bool
 */
type BitVec uint64

func (b *BitVec) Set(i uint8) {
	*b |= 1 << i
}

func (b *BitVec) Unset(i uint8) {
	*b &= ^(1 << i)
}

func (b *BitVec) Get(i uint8) bool {
	return (*b & (1 << i)) != 0
}

func (b *BitVec) Toggle(i uint8) {
	*b ^= 1 << i
}
