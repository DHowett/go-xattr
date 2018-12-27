package xattr

const (
	XattrNoFollow        = 0x00 // Stubbed; Exists on Darwin
	XattrShowCompression = 0x00 // Stubbed; Exists on Darwin
	XattrReplace         = 0x02
	XattrCreate          = 0x01
)

func Setxattr(path, name string, data []byte, offset uint32, options int) (err error) {
	return nil
}

func Getxattr(path, name string, offset uint32, options int) ([]byte, error) {
	return []byte{}, nil
}
