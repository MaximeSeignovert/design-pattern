package pkg

import (
	"bytes"
	"compress/gzip"
	"io"
)

// Compress compresse les données en utilisant gzip
func Compress(data []byte) ([]byte, error) {
	var buf bytes.Buffer
	zw := gzip.NewWriter(&buf)

	_, err := zw.Write(data)
	if err != nil {
		return nil, err
	}

	if err := zw.Close(); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// Decompress décompresse les données gzip
func Decompress(data []byte) ([]byte, error) {
	zr, err := gzip.NewReader(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	defer zr.Close()

	return io.ReadAll(zr)
}
