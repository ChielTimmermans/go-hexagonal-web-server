package internal

import "bytes"

func Concat(values ...string) string {
	var buffer bytes.Buffer
	for _, s := range values {
		buffer.WriteString(s)
	}
	return buffer.String()
}

func ConcatBytes(values ...[]byte) []byte {
	var buffer bytes.Buffer
	for _, s := range values {
		buffer.Write(s)
	}
	return buffer.Bytes()
}
