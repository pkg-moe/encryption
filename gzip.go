package encryption

import (
	"bytes"
	"compress/gzip"
	"io"
)

func Gzip(b []byte) ([]byte, error) {
	var bw bytes.Buffer
	r := gzip.NewWriter(&bw)

	if _, err := r.Write(b); err != nil {
		return nil, err
	}

	if err := r.Flush(); err != nil {
		return nil, err
	}

	if err := r.Close(); err != nil {
		return nil, err
	}

	return bw.Bytes(), nil
}

func UnGzip(b []byte) ([]byte, error) {
	r, err := gzip.NewReader(bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	return io.ReadAll(r)
}
