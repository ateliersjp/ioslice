package ioslice

import (
	"bufio"
	"io"
)

// SliceReader reads from the underlying input source but skips the dropped bytes.
type SliceReader struct {
	*bufio.Reader
	pr      io.Reader
	pw      io.WriteCloser
}

// Slice is equivalent to SliceReader.Take followed by SliceReader.Drop.
func (s *SliceReader) Slice(n int64) io.Reader {
	go s.Drop(n)
	return s.pr
}

// NewReader returns a SliceReader that reads from r but skips the dropped bytes.
func NewReader(r io.Reader) (s *SliceReader) {
	s = &SliceReader{}
	s.Reader = bufio.NewReader(r)
	s.pr, s.pw = io.Pipe()
	return
}
