package cast

type Struct struct {
	Foo string
	Bar string
	Baz string
}

func (s Struct) String() string {
	return s.Foo + s.Bar + s.Baz
}

type StructLike struct {
	Foo string
	Bar string
	Baz string
}

func (s StructLike) String() string {
	return s.Baz + s.Bar + s.Foo
}
