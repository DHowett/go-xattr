package xattr

/*
#include <inttypes.h>
#include <sys/types.h>
#include <sys/xattr.h>
#include <stdlib.h>
#cgo darwin CFLAGS: -DCGO_DARWIN=1
*/
import "C"
import "unsafe"

const (
	XattrNoFollow        = C.XATTR_NOFOLLOW
	XattrShowCompression = C.XATTR_SHOWCOMPRESSION
	XattrReplace         = C.XATTR_REPLACE
	XattrCreate          = C.XATTR_CREATE
)

func Setxattr(path, name string, data []byte, offset uint32, options int) (err error) {
	if data == nil || len(data) <= 0 {
		return nil
	}

	cpath, cattrname := C.CString(path), C.CString(name)
	defer func() {
		C.free(unsafe.Pointer(cpath))
		C.free(unsafe.Pointer(cattrname))
	}()

	var ptr unsafe.Pointer
	if len(data) > 0 {
		ptr = unsafe.Pointer(&data[0])
	}
	ret, err := C.setxattr(cpath, cattrname, ptr, C.size_t(len(data)), C.u_int32_t(offset), C.int(options))
	if ret == -1 {
		return
	}

	return
}

func Getxattr(path, name string, offset uint32, options int) ([]byte, error) {
	cpath, cattrname := C.CString(path), C.CString(name)
	defer func() {
		C.free(unsafe.Pointer(cpath))
		C.free(unsafe.Pointer(cattrname))
	}()

	var ret C.ssize_t
	var err error

	ret, err = C.getxattr(cpath, cattrname, nil, C.size_t(0), C.u_int32_t(offset), C.int(options))
	if ret == -1 {
		return nil, err
	}

	if ret == 0 {
		return []byte{}, nil
	}

	data := make([]byte, int(ret))

	ret, err = C.getxattr(cpath, cattrname, unsafe.Pointer(&data[0]), C.size_t(len(data)), C.u_int32_t(offset), C.int(options))
	if ret == -1 {
		return nil, err
	}

	return data, nil
}
