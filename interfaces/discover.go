package interfaces

type sizer interface {
	Size() int64
}

func fits(capacity int64, v sizer) bool {
	return capacity > v.Size()
}

func isEmailable(v sizer) bool {
	return 1<<20 > v.Size()
}

// sample implementation
type file struct {
	info interface {
		Size() int64
	}
}

func (f *file) size() int64 {
	return f.info.Size()
}

// multiple sizers type implementing size function
type sizers []sizer

func (s sizers) size() int64 {
	var total int64
	for _, sizer := range s {
		total += sizer.Size()
	}

	return total
}

// function implementation
type sizeFunc func() int64

func (s sizeFunc) Size() int64 {
	return s()
}

// concrete type implementation
type size int64

func (s size) Size() int64 {
	return int64(s)
}
