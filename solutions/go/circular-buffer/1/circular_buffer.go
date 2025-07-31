package circular

import "errors"

type Buffer struct {
  buff []byte
  maxSize int
}

func NewBuffer(size int) *Buffer {
  return &Buffer{buff: []byte{}, maxSize: size}
}

func (b *Buffer) ReadByte() (byte, error) {
  if b.IsEmpty() {
    return 0, errors.New("Buffer is empty")
  }

  val := b.buff[0]

  b.buff = b.buff[1:]

  return val, nil
}

func (b *Buffer) WriteByte(c byte) error {
  if b.IsFull() {
    return errors.New("Buffer is full")
  }

  b.buff = append(b.buff, c)
  return nil
}

func (b *Buffer) Overwrite(c byte) {
  if b.IsFull() {
    b.buff = b.buff[1:]
  }

  b.WriteByte(c)
}

func (b *Buffer) Reset() {
  b.buff = []byte{}
}

func (b *Buffer) IsFull() bool {
  return len(b.buff) == b.maxSize
}

func (b *Buffer) IsEmpty() bool {
  return len(b.buff) == 0
}
