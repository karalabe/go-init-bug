package good

type Sizer interface {
	Size() int
}

func Size(obj Sizer) int {
	return obj.Size()
}
