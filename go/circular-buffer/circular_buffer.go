package circular

import (
	"errors"
)

type Buffer struct {
	data  []byte
	first int
	last  int
}

func NewBuffer(size int) *Buffer {
	arr := make([]byte, size)
	return &Buffer{data: arr, first: -1, last: -1}
}

func (buf *Buffer) ReadByte() (byte, error) {
	if buf.first == -1 {
		return 0, errors.New("Read from empty buffer.")
	}

	output := buf.data[buf.first]

	if buf.first == buf.last {
		buf.first = -1
		buf.last = -1
	} else {
		buf.first = (buf.first + 1) % len(buf.data)
	}

	return output, nil
}
func (buf *Buffer) WriteByte(c byte) error {
	if buf.first != -1 && (buf.last >= buf.first && buf.last-buf.first+1 == len(buf.data)) ||
		(buf.last < buf.first && buf.last+len(buf.data)-buf.first+1 == len(buf.data)) {
		return errors.New("Write to full buffer.")
	}

	if buf.first == -1 {
		buf.first = 0
	}
	buf.last = (buf.last + 1) % len(buf.data)
	buf.data[buf.last] = c

	return nil
}
func (buf *Buffer) Overwrite(c byte) {
	err := buf.WriteByte(c)
	if err != nil {
		buf.data[buf.first] = c
		buf.first = (buf.first + 1) % len(buf.data)
		buf.last = (buf.last + 1) % len(buf.data)
	}
}
func (buf *Buffer) Reset() { // put buffer in an empty state
	buf.first = -1
	buf.last = -1
}
