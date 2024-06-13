package ioslice

import (
	"bufio"
	"io"
)

type SliceReader struct {
	*bufio.Reader
	pr      io.Reader
	pw      io.WriteCloser
}

func (s *SliceReader) Slice(n int64) io.Reader {
	go s.Drop(n)
	return s.pr
}

func NewReader(r io.Reader) (s *SliceReader) {
	s = &SliceReader{}
	s.Reader = bufio.NewReader(r)
	s.pr, s.pw = io.Pipe()
	return
}
