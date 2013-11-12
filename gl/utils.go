package gl

//#include "common.h"
import "C"

import (
	"errors"
	"reflect"
	"unsafe"
)

var sliceOfUint32 = reflect.TypeOf([]uint32(nil))

func sliceToGLData(slice interface{}) (C.GLsizei, C.GLenum, unsafe.Pointer, error) {
	v := reflect.ValueOf(slice)

	if v.Kind() != reflect.Slice {
		return 0, 0, nil, errors.New("gl: Could not convert to GL data, not a slice")
	}

	var enum C.GLenum
	switch v.Type() {
	case sliceOfUint32:
		enum = C.GL_UNSIGNED_INT
	default:
		return 0, 0, nil, errors.New("gl: Unrecognized type of slice")
	}

	return C.GLsizei(v.Len()), enum, unsafe.Pointer(v.Pointer()), nil
}
