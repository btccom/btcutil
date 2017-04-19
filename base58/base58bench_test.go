// Copyright (c) 2013-2014 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package base58_test

import (
	"bytes"
	"testing"

	"github.com/btcsuite/btcutil/base58"
)

var result string
var resultb []byte

func BenchmarkBase58FastEncode(b *testing.B) {
	b.StopTimer()
	data := bytes.Repeat([]byte{0xff}, 5000)
	b.SetBytes(int64(len(data)))
	b.StartTimer()

	var s string
	for i := 0; i < b.N; i++ {
		s = base58.FastEncode(data)
	}

	result = s
}

func BenchmarkBase58SlowEncode(b *testing.B) {
	b.StopTimer()
	data := bytes.Repeat([]byte{0xff}, 5000)
	b.SetBytes(int64(len(data)))
	b.StartTimer()

	var s string
	for i := 0; i < b.N; i++ {
		s = base58.SlowEncode(data)
	}

	result = s
}
func BenchmarkBase58FastEncodeSmall(b *testing.B) {
	b.StopTimer()
	data := bytes.Repeat([]byte{0xff}, 30)
	b.SetBytes(int64(len(data)))
	b.StartTimer()

	var s string
	for i := 0; i < b.N; i++ {
		s = base58.FastEncode(data)
	}

	result = s
}

func BenchmarkBase58SlowEncodeSmall(b *testing.B) {
	b.StopTimer()
	data := bytes.Repeat([]byte{0xff}, 30)
	b.SetBytes(int64(len(data)))
	b.StartTimer()

	var s string
	for i := 0; i < b.N; i++ {
		s = base58.SlowEncode(data)
	}

	result = s
}

func BenchmarkBase58Decode(b *testing.B) {
	b.StopTimer()
	data := bytes.Repeat([]byte{0xff}, 5000)
	encoded := base58.Encode(data)
	b.SetBytes(int64(len(encoded)))
	b.StartTimer()

	var bs []byte
	for i := 0; i < b.N; i++ {
		bs = base58.Decode(encoded)
	}

	resultb = bs
}
