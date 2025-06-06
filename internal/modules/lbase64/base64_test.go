package lbase64

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var stringTestData map[string]string = map[string]string{
	"hello world": "aGVsbG8gd29ybGQ=",
}

var urlTestData map[string]string = map[string]string{
	"https://www.example.com/bla?key1=val1&key2=val2#boo": "aHR0cHM6Ly93d3cuZXhhbXBsZS5jb20vYmxhP2tleTE9dmFsMSZrZXkyPXZhbDIjYm9v",
}

func TestEncodeString(t *testing.T) {
	for key, val := range stringTestData {
		assert.Equal(t, val, doEncodeString(key))
	}
}

func TestDecodeString(t *testing.T) {
	for key, val := range stringTestData {
		v, err := doDecodeString(val)
		assert.Nil(t, err)
		assert.Equal(t, key, string(v))
	}
}

func TestEncodeUrl(t *testing.T) {
	for key, val := range urlTestData {
		assert.Equal(t, val, doEncodeUrl(key))
	}
}

func TestDecodeUrl(t *testing.T) {
	for key, val := range urlTestData {
		v, err := doDecodeUrl(val)
		assert.Nil(t, err)
		assert.Equal(t, key, string(v))
	}
}
