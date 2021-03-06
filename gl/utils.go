package gl

//#include "common.h"
import "C"

import (
	"errors"
	"reflect"
	"unsafe"
)

var (
	sliceOfByte    = reflect.TypeOf([]byte(nil))
	sliceOfInt8    = reflect.TypeOf([]int8(nil))
	sliceOfUint8   = reflect.TypeOf([]uint8(nil))
	sliceOfInt16   = reflect.TypeOf([]int16(nil))
	sliceOfUint16  = reflect.TypeOf([]uint16(nil))
	sliceOfInt32   = reflect.TypeOf([]int32(nil))
	sliceOfUint32  = reflect.TypeOf([]uint32(nil))
	sliceOfFloat32 = reflect.TypeOf([]float32(nil))
	sliceOfFloat64 = reflect.TypeOf([]float64(nil))
)

var (
	sizeOfInt8    = unsafe.Sizeof([1]int8{})
	sizeOfUint8   = unsafe.Sizeof([1]uint8{})
	sizeOfInt16   = unsafe.Sizeof([1]int16{})
	sizeOfUint16  = unsafe.Sizeof([1]uint16{})
	sizeOfInt32   = unsafe.Sizeof([1]int32{})
	sizeOfUint32  = unsafe.Sizeof([1]uint32{})
	sizeOfFloat32 = unsafe.Sizeof([1]float32{})
	sizeOfFloat64 = unsafe.Sizeof([1]float64{})
)

/*
sliceToGLData

Returns
    length of slice
    size of slice in bytes
    type of slice in GL enum
    pointer to array

Note: Does not support GL_2_BYTES, GL_3_BYTES, or GL_4_BYTES.
*/
func sliceToGLData(slice interface{}) (C.GLsizei, C.GLsizei, C.GLenum, unsafe.Pointer, error) {
	v := reflect.ValueOf(slice)

	if v.Kind() != reflect.Slice {
		return 0, 0, 0, nil, errors.New("gl: Could not convert to GL data, not a slice")
	}

	var enum C.GLenum
	var size int
	switch v.Type() {
	case sliceOfByte:
		enum = C.GL_BYTE
		size = int(sizeOfInt8)
	case sliceOfInt8:
		enum = C.GL_BYTE
		size = int(sizeOfInt8)
	case sliceOfUint8:
		enum = C.GL_UNSIGNED_BYTE
		size = int(sizeOfUint8)
	case sliceOfInt16:
		enum = C.GL_SHORT
		size = int(sizeOfInt16)
	case sliceOfUint16:
		enum = C.GL_UNSIGNED_SHORT
		size = int(sizeOfUint16)
	case sliceOfInt32:
		enum = C.GL_INT
		size = int(sizeOfInt32)
	case sliceOfUint32:
		enum = C.GL_UNSIGNED_INT
		size = int(sizeOfUint32)
	case sliceOfFloat32:
		enum = C.GL_FLOAT
		size = int(sizeOfFloat32)
	case sliceOfFloat64:
		enum = C.GL_FLOAT
		size = int(sizeOfFloat64)
	default:
		return 0, 0, 0, nil, errors.New("gl: Unrecognized type of slice")
	}

	return C.GLsizei(v.Len()), C.GLsizei(v.Len() * size), enum, unsafe.Pointer(v.Pointer()), nil
}

func glTypeSize(glType GLenum) int {
	var size int

	switch glType {
	case ByteConst:
		size = int(sizeOfInt8)
	case UnsignedByteConst:
		size = int(sizeOfUint8)
	case ShortConst:
		size = int(sizeOfInt16)
	case UnsignedShortConst:
		size = int(sizeOfUint16)
	case IntConst:
		size = int(sizeOfInt32)
	case UnsignedIntConst:
		size = int(sizeOfUint32)
	case FloatConst:
		size = int(sizeOfFloat32)
	default:
		size = 0
	}

	return size
}

func boolToGLBool(value bool) C.GLboolean {
	if value {
		return C.GL_TRUE
	}
	return C.GL_FALSE
}

func glBoolToBool(value C.GLboolean) bool {
	if value == C.GL_TRUE {
		return true
	}
	return false
}

func IsTrue(value int) bool {
	return (value == C.GL_TRUE)
}
