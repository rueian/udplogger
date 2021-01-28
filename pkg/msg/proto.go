package msg

import (
	"encoding/binary"
	"errors"
)

type Message struct {
	Check     uint16
	Group     uint16
	Code      uint32
	Size      uint32
	Identity  uint64
	Duration  uint64
	Timestamp uint64
	Remaining []byte
}

func (m *Message) FromDatagram(in []byte) error {
	if len(in) < 36 {
		return ErrMalformed
	}

	m.Check = binary.BigEndian.Uint16(in[0:2])
	m.Group = binary.BigEndian.Uint16(in[2:4])
	m.Code = binary.BigEndian.Uint32(in[4:8])
	m.Size = binary.BigEndian.Uint32(in[8:12])

	m.Identity = binary.BigEndian.Uint64(in[12:20])
	m.Duration = binary.BigEndian.Uint64(in[20:28])
	m.Timestamp = binary.BigEndian.Uint64(in[28:36])
	m.Remaining = in[36:]
	return nil
}

func (m *Message) ToDatagram(out []byte) (int, error) {
	if len(out) < 36 {
		return 0, ErrMalformed
	}

	binary.BigEndian.PutUint16(out[0:2], m.Check)
	binary.BigEndian.PutUint16(out[2:4], m.Group)
	binary.BigEndian.PutUint32(out[4:8], m.Code)
	binary.BigEndian.PutUint32(out[8:12], m.Size)

	binary.BigEndian.PutUint64(out[12:20], m.Identity)
	binary.BigEndian.PutUint64(out[20:28], m.Duration)
	binary.BigEndian.PutUint64(out[28:36], m.Timestamp)

	remain := len(out) - 36
	if left := len(m.Remaining); left < remain {
		remain = left
	}
	copy(out[36:], m.Remaining[:remain])
	return 36 + remain, nil
}

var (
	ErrMalformed = errors.New("malformed")
)
