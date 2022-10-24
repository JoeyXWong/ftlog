package mytypes

type (
	MyInt    int
	MyString string
)

func (i MyInt) Twice() MyInt {
	return i * 2
}

func (s MyString) StringLen() int {
	return len(s)
}
