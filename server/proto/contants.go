package proto

const (
	PacketHeaderSize  = 4 + 1 + 16 + 16 // size + type + origin + target
	PacketPostfixSize = 4               // hashsum
	PacketBufferSize  = 128
	PacketMaxMessage  = 4 * 1024 * 1024 // 4 MB
)