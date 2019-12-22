package main

import (
	"bytes"
	"io"
)

type newlineRemovingReader struct {
	reader io.Reader
}

func newNewlineRemovingReader(reader io.Reader) *newlineRemovingReader {
	return &newlineRemovingReader{reader: reader}
}

func isNotNewline(r byte) bool {
	if r == '\n' || r == '\r' {
		return false
	}
	return true
}

func (a *newlineRemovingReader) Read(p []byte) (int, error) {
	n, err := a.reader.Read(p)
	if err != nil {
		return n, err
	}
	read := 0
	buf := new(bytes.Buffer)

	for i := 0; i < n; i++ {
		if isNotNewline(p[i]) {
			buf.WriteByte(p[i])
			read++
		}
	}

	copy(p, buf.Bytes())
	return read, nil
}
