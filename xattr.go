package xattr

// #include <inttypes.h>
// #include <sys/types.h>
// #include <sys/xattr.h>
// #include <stdlib.h>
// #cgo linux CFLAGS: -DXATTR_SHOWCOMPRESSION=0
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

	ret, err := C.setxattr(cpath, cattrname, unsafe.Pointer(&data[0]), C.size_t(len(data)), C.u_int32_t(offset), C.int(options))
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

	ret, err = C.getxattr(cpath, cattrname, nil, C.size_t(0), C.u_int32_t(offset), C.int(options))
	if ret == -1 {
		return nil, err
	}

	data := make([]byte, int(ret))

	ret, err = C.getxattr(cpath, cattrname, unsafe.Pointer(&data[0]), C.size_t(len(data)), C.u_int32_t(offset), C.int(options))
	if ret == -1 {
		return nil, err
	}

	return data, nil
}
