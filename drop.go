package ioslice

import (
	"io"
)

func (s *SliceReader) DropN(n int64) (int64, error) {
	return io.CopyN(s.pw, s, n)
}

func (s *SliceReader) DropUntil(delim byte) (data []byte, err error) {
	if data, err = s.ReadBytes(delim); err == nil || err == io.EOF {
		_, err = s.pw.Write(data)
	}
	return
}

func (s *SliceReader) DropClose() error {
	return s.pw.Close()
}

func (s *SliceReader) Drop(n int64) (int64, error) {
	defer s.DropClose()
	return s.DropN(n)
}

func (s *SliceReader) Take() io.Reader {
	return s.pr
}
