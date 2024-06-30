# ioslice

[![](https://godoc.org/github.com/ateliersjp/ioslice?status.svg)](http://godoc.org/github.com/ateliersjp/ioslice)

This package slices an underlying input source into two Readers; a SliceReader reads from the underlying input source but drops the head bytes or lines while another Reader reads the dropped portion.
