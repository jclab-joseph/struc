package struc

import (
	"io"
	"reflect"
)

type Packer interface {
	Pack(buf []byte, val reflect.Value, options *Options) (int, error)
	Unpack(r io.Reader, val reflect.Value, options *Options) error
	Sizeof(val reflect.Value, options *Options) int
	OffsetOf(val reflect.Value, options *Options, index int) int
	String() string
}
