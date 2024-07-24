package good

var sizeA = 100

type A struct{}

func (a *A) Size() int {
	return sizeA
}

var sizeC = Size((*A)(nil)) + 10

type C struct{}

func (c *C) Size() int {
	return sizeC
}
