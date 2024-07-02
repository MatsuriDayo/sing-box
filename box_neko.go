package box

import (
	"io"
	"reflect"
	"unsafe"
)

// SetLogWritter used in nekoray, change log writer without providing platform interface
func (s *Box) SetLogWritter(w io.Writer) {
	writer_ := reflect.Indirect(reflect.ValueOf(s.logFactory)).FieldByName("writer")
	writer_ = reflect.NewAt(writer_.Type(), unsafe.Pointer(writer_.UnsafeAddr())).Elem()
	writer_.Set(reflect.ValueOf(w))
}
