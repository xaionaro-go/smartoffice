package smartoffice

import "bytes"

type BufferCloser struct {
	*bytes.Buffer
}

func (*BufferCloser) Close() error {
	return nil
}
