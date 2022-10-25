package mytypes

import "strings"

type (
	MyInt     int
	MyString  string
	MyBuilder struct {
		Contents strings.Builder
	}
)

func (i MyInt) Twice() MyInt {
	return i * 2
}

func (s MyString) StringLen() int {
	return len(s)
}

func (b MyBuilder) Len() int {
	return 0
}

func (b MyBuilder) Hello() string {
	return "Hello, Gophers!"
}

func (b MyBuilder) ToUpper() string {
	return strings.ToUpper(b.Contents.String())
}

func (i *MyInt) Double() {
	*i *= 2
}
