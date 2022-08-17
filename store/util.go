package store

func copyBytes(value []byte) []byte {
	copied := make([]byte, len(value))
	copy(copied, value)
	return copied
}
