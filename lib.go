package stream

import (
	"github.com/gogo/protobuf/proto"
	"io"
	"encoding/binary"
	"sync"
)

type Message interface {
	proto.Marshaler
	proto.Unmarshaler
	proto.Sizer
}

type LockableWriter interface {
	sync.Locker
	io.Writer
}

func Write(w io.Writer, m Message) (err error) {
	messageBytes, err := m.Marshal()

	if err == nil {
		size := m.Size()
		sizeBytes := make([]byte, 8)
		binary.BigEndian.PutUint64(sizeBytes, uint64(size))

		w.Write(sizeBytes)
		w.Write(messageBytes)
	}

	return
}

func WriteWithLock(w LockableWriter, m Message) (err error) {
	messageBytes, err := m.Marshal()

	if err == nil {
		size := m.Size()
		sizeBytes := make([]byte, 8)
		binary.BigEndian.PutUint64(sizeBytes, uint64(size))

		w.Lock()
		defer w.Unlock()
		w.Write(sizeBytes)
		w.Write(messageBytes)
	}

	return
}

func Read(r io.Reader, m Message) (err error) {
	sizeBytes  := make([]byte, 8)

	if _, err = io.ReadFull(r, sizeBytes); err == nil {
		size := binary.BigEndian.Uint64(sizeBytes)
		messageBytes := make([]byte, size)

		if _, err = io.ReadFull(r, messageBytes); err == nil {
			m.Unmarshal(messageBytes)
		}
	}

	return
}