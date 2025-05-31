package gosimple

//var _ types.SimpleDB[map[string]any] = (*GoSimple)(nil)

type GoSimple[T any] struct {
	index map[string]string
	data  map[string]T
}
