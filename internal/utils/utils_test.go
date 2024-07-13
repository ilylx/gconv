package utils

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"regexp"
	"testing"
)

var (
	replaceCharReg, _ = regexp.Compile(`[\-\.\_\s]+`)
)

func TestReadCloser_Read(t *testing.T) {
	var (
		n    int
		b    = make([]byte, 3)
		body = NewReadCloser([]byte{1, 2, 3, 4}, false)
	)

	n, _ = body.Read(b)
	assert.Equal(t, b[:n], []byte{1, 2, 3})

	n, _ = body.Read(b)
	assert.Equal(t, b[:n], []byte{4})

	n, _ = body.Read(b)
	assert.Equal(t, b[:n], []byte{})

	n, _ = body.Read(b)
	assert.Equal(t, b[:n], []byte{})
}

func TestReadCloser_ReadAll(t *testing.T) {
	var (
		r    []byte
		body = NewReadCloser([]byte{1, 2, 3, 4}, false)
	)
	r, _ = ioutil.ReadAll(body)
	assert.Equal(t, r, []byte{1, 2, 3, 4})
	r, _ = ioutil.ReadAll(body)
	assert.Equal(t, r, []byte{})
}

func TestReadCloser_ReadAndReadAll(t *testing.T) {
	var (
		n    int
		r    []byte
		b    = make([]byte, 3)
		body = NewReadCloser([]byte{1, 2, 3, 4}, true)
	)
	n, _ = body.Read(b)
	assert.Equal(t, b[:n], []byte{1, 2, 3})
	n, _ = body.Read(b)
	assert.Equal(t, b[:n], []byte{4})

	n, _ = body.Read(b)
	assert.Equal(t, b[:n], []byte{1, 2, 3})
	n, _ = body.Read(b)
	assert.Equal(t, b[:n], []byte{4})

	r, _ = ioutil.ReadAll(body)
	assert.Equal(t, r, []byte{1, 2, 3, 4})
	r, _ = ioutil.ReadAll(body)
	assert.Equal(t, r, []byte{1, 2, 3, 4})
}

func Test_RemoveSymbols(t *testing.T) {
	assert.Equal(t, RemoveSymbols(`-a-b._a c1!@#$%^&*()_+:";'.,'01`), `abac101`)
}

func Benchmark_RemoveSymbols(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RemoveSymbols(`-a-b._a c1!@#$%^&*()_+:";'.,'01`)
	}
}

func Benchmark_RegularReplaceChars(b *testing.B) {
	for i := 0; i < b.N; i++ {
		replaceCharReg.ReplaceAllString(`-a-b._a c1!@#$%^&*()_+:";'.,'01`, "")
	}
}
