package nsqd

import (
	"testing"
	"unsafe"

	"github.com/bitly/go-nsq"
)

func BenchmarkGUIDCopy(b *testing.B) {
	source := make([]byte, 16)
	var dest nsq.MessageID
	for i := 0; i < b.N; i++ {
		copy(dest[:], source)
	}
}

func BenchmarkGUIDUnsafe(b *testing.B) {
	source := make([]byte, 16)
	var dest nsq.MessageID
	for i := 0; i < b.N; i++ {
		dest = *(*nsq.MessageID)(unsafe.Pointer(&source[0]))
	}
	_ = dest
}

func BenchmarkGUID(b *testing.B) {
	factory := &guidFactory{}
	for i := 0; i < b.N; i++ {
		guid, err := factory.NewGUID(0)
		if err != nil {
			continue
		}
		guid.Hex()
	}
}
