package lzma

import (
	"testing"
)

func TestLzmaCompress(t *testing.T) {
	str := "1234567890"
	r1, err := Compress([]byte(str), 9)
	if err != nil {
		t.Error(err.Error(), r1)
	}
	r2, err := Uncompress(r1)
	if err != nil {
		t.Error(err.Error())
	}
	if string(r2) != str {
		t.Error("lzma test failed")
	}
}

func BenchmarkLzma(b *testing.B) {
	for i := 0; i < b.N; i++ {
		str := "1234567890"
		r1, _ := Compress([]byte(str), 5)
		r2, _ := Uncompress(r1)
		if string(r2) != str {
			b.Fail()
		}
	}
}
