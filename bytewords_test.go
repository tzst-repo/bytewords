package bytewords

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncodeShortBytewords(t *testing.T) {
	assert.Equal(t, "aeadzm", EncodeShortBytewords([]byte("\x00\x01\xff")))
}

func TestDecodeShortBytewords(t *testing.T) {
	assert.Equal(t, []byte("\x00\x01\xff"), DecodeShortBytewords("aeadzm"))
}
