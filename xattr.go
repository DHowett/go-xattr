package xattr

/*
#include <inttypes.h>
#include <sys/types.h>
#include <sys/xattr.h>
#include <stdlib.h>
#cgo linux CFLAGS: -DCGO_LINUX=1
#cgo darwin CFLAGS: -DCGO_DARWIN=1
#if CGO_LINUX == 1
#define XATTR_SHOWCOMPRESSION 0
#define XATTR_NOFOLLOW 0
static int _cgo_setxattr_thunk(const char *path, const char *attr, void *val, size_t size, uint32_t offset, int opts) {
	return setxattr(path, attr, val, size, opts);
}
static ssize_t _cgo_getxattr_thunk(const char *path, const char *attr, void *val, size_t size, uint32_t offset, int opts) {
	return getxattr(path, attr, val, size);
}
#elif CGO_DARWIN == 1
static int _cgo_setxattr_thunk(const char *path, const char *attr, void *val, size_t size, uint32_t offset, int opts) {
	return setxattr(path, attr, val, size, offset, opts);
}
static ssize_t _cgo_getxattr_thunk(const char *path, const char *attr, void *val, size_t size, uint32_t offset, int opts) {
	return getxattr(path, attr, val, size, offset, opts);
}
#endif
*/
import "C"
import "unsafe"

const (
	XattrNoFollow        = C.XATTR_NOFOLLOW
	XattrReplace         = C.XATTR_REPLACE
	XattrCreate          = C.XATTR_CREATE
	XattrShowCompression = C.XATTR_SHOWCOMPRESSION
)

func Setxattr(path, name string, data []byte, offset uint32, options int) error {
	cpath, cattrname := C.CString(path), C.CString(name)
	defer func() {
		C.free(unsafe.Pointer(cpath))
		C.free(unsafe.Pointer(cattrname))
	}()

	ret, err := C._cgo_setxattr_thunk(cpath, cattrname, unsafe.Pointer(&data[0]), C.size_t(len(data)), C.uint32_t(offset), C.int(options))
	if ret == -1 {
		return err
	}

	return nil
}

func Getxattr(path, name string, offset uint32, options int) ([]byte, error) {
	cpath, cattrname := C.CString(path), C.CString(name)
	defer func() {
		C.free(unsafe.Pointer(cpath))
		C.free(unsafe.Pointer(cattrname))
	}()

	var ret C.ssize_t
	var err error

	ret, err = C._cgo_getxattr_thunk(cpath, cattrname, nil, C.size_t(0), C.uint32_t(offset), C.int(options))
	if ret == -1 {
		return nil, err
	}

	data := make([]byte, int(ret))

	ret, err = C._cgo_getxattr_thunk(cpath, cattrname, unsafe.Pointer(&data[0]), C.size_t(len(data)), C.uint32_t(offset), C.int(options))
	if ret == -1 {
		return nil, err
	}

	return data, nil
}
