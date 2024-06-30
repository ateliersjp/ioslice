package ioslice

import (
	"io"
)

// DropN drops n bytes in SliceReader. The dropped data can be taken with SliceReader.Take.
func (s *SliceReader) DropN(n int64) (int64, error) {
	return io.CopyN(s.pw, s, n)
}

// DropUntil drops until the first occurrence of delim in SliceReader. The dropped data can be taken with SliceReader.Take.
func (s *SliceReader) DropUntil(delim byte) (data []byte, err error) {
	if data, err = s.ReadBytes(delim); err == nil || err == io.EOF {
		_, err = s.pw.Write(data)
	}
	return
}

// DropClose must be called after all SliceReader.Drop* operations done.
func (s *SliceReader) DropClose() error {
	return s.pw.Close()
}

// Drop is equivalent to SliceReader.DropN followed by SliceReader.DropClose.
func (s *SliceReader) Drop(n int64) (int64, error) {
	defer s.DropClose()
	return s.DropN(n)
}

// Take must be called beforehand to allocate another Reader that reads the dropped data.
func (s *SliceReader) Take() io.Reader {
	return s.pr
}
