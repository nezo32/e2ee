package proto

import "fmt"

var (
	ErrorPacketZeroLength     = fmt.Errorf("invalid packet: zero length")
	ErrorPacketTooLarge       = fmt.Errorf("invalid packet: too large")
	ErrorPacketNoHeader       = fmt.Errorf("invalid packet: no header")
	ErrorPacketInvalidHashSum = fmt.Errorf("invalid packet: hashsum mismatch")
)