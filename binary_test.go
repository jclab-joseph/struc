package struc

import (
	"bytes"
	"encoding/hex"
	"testing"
)

func TestBinaryFallback(t *testing.T) {
	var buf bytes.Buffer
	v := []byte("abcdefg")
	if err := Pack(&buf, &v); err != nil {
		t.Fatal(err)
	}
	println(hex.EncodeToString(buf.Bytes()))
}
