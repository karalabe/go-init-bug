package bad

var sizeB = Size((*C)(nil)) + Size((*C)(nil))

type B struct{}

func (b *B) Size() int {
	return sizeB
}
