package proto

import (
	"encoding/binary"

	"github.com/charmbracelet/log"
	"github.com/google/uuid"
	"github.com/nezo32/e2ee/security"
)

const (
	PacketTypeFile = 1 << iota
	PacketTypeEncrypted
)

type packetImpl struct {
	size       uint32
	data       []byte
	packetType byte

	origin uuid.UUID
	target uuid.UUID
}

type Packet interface {
	Build() []byte

	GetData() []byte
	GetSize() uint32
	GetTarget() uuid.UUID
	GetOrigin() uuid.UUID

	IsFile() bool
	IsEncrypted() bool

	SetData(data []byte)
	SetOrigin(id uuid.UUID)
	SetTarget(id uuid.UUID)
}

type PacketParams struct {
	IsFile      bool
	IsEncrypted bool

	Data []byte

	Origin uuid.UUID
	Target uuid.UUID
}

func NewPacket(params *PacketParams) Packet {
	p := &packetImpl{
		size:       uint32(PacketHeaderSize + len(params.Data) + PacketPostfixSize),
		data:       params.Data,
		origin:     params.Origin,
		target:     params.Target,
		packetType: byte(0),
	}

	if params.IsFile {
		p.packetType |= PacketTypeFile
	}

	if params.IsEncrypted {
		p.packetType |= PacketTypeEncrypted
	}

	return p
}

func PacketFromBytes(data []byte) (Packet, error) {
	if len(data) < PacketHeaderSize+PacketPostfixSize+1 {
		return nil, ErrorPacketZeroLength
	}

	postfixPosition := len(data) - PacketPostfixSize

	hashSum := binary.LittleEndian.Uint32(data[postfixPosition:])

	if security.HashSum(data[:postfixPosition]) != uint32(hashSum) {
		return nil, ErrorPacketInvalidHashSum
	}

	p := &packetImpl{
		data:       data[PacketHeaderSize : len(data)-PacketPostfixSize],
		size:       binary.LittleEndian.Uint32(data[:4]),
		packetType: data[4],
		origin:     uuid.UUID(data[5 : 5+16]),
		target:     uuid.UUID(data[5+16 : 5+16+16]),
	}

	return p, nil
}

func (p *packetImpl) Build() []byte {
	// allocate
	buf := make([]byte, 0, p.size)

	// header
	buf = binary.LittleEndian.AppendUint32(buf, p.size)
	buf = append(buf, byte(p.packetType))
	buf = append(buf, p.origin[:]...)
	buf = append(buf, p.target[:]...)

	// data
	buf = append(buf, p.data...)

	// postfix
	hs := security.HashSum(buf)
	buf = binary.LittleEndian.AppendUint32(buf, hs)

	log.Info("Hashing buffer", "hashsum", hs)

	return buf
}

func (p *packetImpl) GetData() []byte {
	return p.data
}

func (p *packetImpl) GetSize() uint32 {
	return p.size
}

func (p *packetImpl) GetTarget() uuid.UUID {
	return p.target
}

func (p *packetImpl) GetOrigin() uuid.UUID {
	return p.origin
}

func (p *packetImpl) IsFile() bool {
	return p.packetType&PacketTypeFile != 0
}

func (p *packetImpl) IsEncrypted() bool {
	return p.packetType&PacketTypeEncrypted != 0
}

func (p *packetImpl) SetData(data []byte) {
	p.data = data
	p.size = uint32(PacketHeaderSize + len(data) + PacketPostfixSize)
}

func (p *packetImpl) SetOrigin(id uuid.UUID) {
	p.origin = id
}

func (p *packetImpl) SetTarget(id uuid.UUID) {
	p.target = id
}
