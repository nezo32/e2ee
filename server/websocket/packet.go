package websocket

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"time"

	"github.com/nezo32/e2ee/proto"
	"golang.org/x/net/websocket"
)

func WritePacket(conn *websocket.Conn, packet proto.Packet) error {
	data := packet.Build()
	totalWritten := 0
	for totalWritten < len(data) {
		n, err := conn.Write(data[totalWritten:])
		if err != nil {
			return fmt.Errorf("websocket error: write error after %d bytes: %w", totalWritten, err)
		}

		totalWritten += n

		if totalWritten < len(data) {
			time.Sleep(10 * time.Millisecond)
		}
	}
	return nil
}

func ReadPacket(r io.Reader) (proto.Packet, error) {
	buf := make([]byte, proto.PacketBufferSize)
	var completeMessage bytes.Buffer

	firstPacket := true
	totalRead := 0
	targetRead := uint32(0)

	for {
		n, err := r.Read(buf)
		if err != nil {
			if err == io.EOF && totalRead > 0 {
				return nil, io.EOF
			}
			return nil, err
		}

		totalRead += n
		completeMessage.Write(buf[:n])

		if firstPacket {
			if len(buf) < proto.PacketHeaderSize {
				return nil, proto.ErrorPacketNoHeader
			}
			firstPacket = false
			targetRead = binary.LittleEndian.Uint32(buf[:4])
		}

		if uint32(totalRead) >= targetRead {
			break
		}

		if totalRead > proto.PacketMaxMessage {
			return nil, proto.ErrorPacketTooLarge
		}
	}

	return proto.PacketFromBytes(completeMessage.Bytes())
}