package gosimple

func NewSimpleDB[T any]() *GoSimple[T] {
	return &GoSimple[T]{
		index: make(map[string]string),
		data:  make(map[string]T)}
}
