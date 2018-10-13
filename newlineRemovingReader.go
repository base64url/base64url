package main

import (
	"io"
)

type newlineRemovingReader struct {
	reader io.Reader
}

func NewNewlineRemovingReader(reader io.Reader) *newlineRemovingReader {
	return &newlineRemovingReader{reader: reader}
}

func isNewline(r byte) byte {
	if (r == '\n' || r == '\r') {
		return 0
	}
	return r
}

func (a *newlineRemovingReader) Read(p []byte) (int, error) {
	n, err := a.reader.Read(p)
	if err != nil {
		return n, err
	}
	read := 0
	buf := make([]byte, n)
	for i := 0; i < n; i++ {
		if char := isNewline(p[i]); char != 0 {
			buf[i] = char
			read++
		}
	}

	copy(p, buf)
	return read - 1, nil
}