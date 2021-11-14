package pool

import (
	"bytes"
	"strings"
	"sync"
)

var (
	poolBuffer = &sync.Pool{
		New: func() interface{} {
			return new(bytes.Buffer)
		},
	}
	poolStrBuilder = &sync.Pool{
		New: func() interface{} {
			return new(strings.Builder)
		},
	}
)

// GetBuffer returns the bytes.Buffer instance.
func GetBuffer() *bytes.Buffer {
	instance := poolBuffer.Get().(*bytes.Buffer)
	instance.Reset()
	return instance
}

// GetStrBuilder returns the strings.Builder instance.
func GetStrBuilder() *strings.Builder {
	instance := poolStrBuilder.Get().(*strings.Builder)
	instance.Reset()
	return instance
}

// Put puts the instance back to the pool.
func Put(val interface{}) {
	switch instance := val.(type) {
	case *bytes.Buffer:
		poolBuffer.Put(instance)
	case *strings.Builder:
		poolStrBuilder.Put(instance)
	}
}
