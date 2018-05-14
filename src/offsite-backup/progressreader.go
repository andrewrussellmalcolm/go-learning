package main

import (
	"io"
)

// ProgressReader
type ProgressReader struct {
	r     io.Reader
	f     func(chunk, total int64)
	total int64
}

// NewProgressReader makes a new Reader that counts the bytes
// read through it.
func NewProgressReader(r io.Reader, f func(chunk, total int64)) *ProgressReader {
	return &ProgressReader{r: r, f: f}
}

func (pr *ProgressReader) Read(p []byte) (n int, err error) {
	n, err = pr.r.Read(p)
	pr.total += int64(n)
	pr.f(int64(n), pr.total)
	return
}
