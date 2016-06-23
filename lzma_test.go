package lzma

import (
	"log"
	"testing"
)

func TestLzmaCompress(t *testing.T) {
	str := "1234567890"
	r1, err := Compress([]byte(str))
	if err != nil {
		t.Error(err.Error(), r1)
	}
	log.Println(r1, len(r1))
	r2, err := Uncompress(r1)
	if err != nil {
		t.Error(err.Error())
	}
	if string(r2) != str {
		t.Error("lzma test failed")
	}
	log.Println(string(r2))
}
