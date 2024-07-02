package box

import (
	"io"
	"reflect"
	"unsafe"

	"github.com/sagernet/sing-box/log"
)

func (s *Box) SetLogWritter(w io.Writer) {
	writer_ := reflect.Indirect(reflect.ValueOf(s.logFactory)).FieldByName("writer")
	writer_ = reflect.NewAt(writer_.Type(), unsafe.Pointer(writer_.UnsafeAddr())).Elem()
	writer_.Set(reflect.ValueOf(w))
}

func (s *Box) GetLogPlatformFormatter() *log.Formatter {
	platformFormatter_ := reflect.Indirect(reflect.ValueOf(s.logFactory)).FieldByName("platformFormatter")
	platformFormatter_ = reflect.NewAt(platformFormatter_.Type(), unsafe.Pointer(platformFormatter_.UnsafeAddr()))
	platformFormatter := platformFormatter_.Interface().(*log.Formatter)
	return platformFormatter
}
