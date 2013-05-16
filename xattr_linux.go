package xattr

import "syscall"

const (
	XattrNoFollow        = 0x00 // Stubbed; Exists on Darwin
	XattrShowCompression = 0x00 // Stubbed; Exists on Darwin
	XattrReplace         = 0x02
	XattrCreate          = 0x01
)


func Setxattr(path, name string, data []byte, offset uint32, options int) (err error) {
	err = syscall.Setxattr(path, name, data, options)
	return
}

func Getxattr(path, name string, offset uint32, options int) ([]byte, error) {
	var err error

	sz, err := syscall.Getxattr(path, name, nil)
	if err != nil {
		return nil, err
	}

	data := make([]byte, int(sz))

	sz, err = syscall.Getxattr(path, name, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
