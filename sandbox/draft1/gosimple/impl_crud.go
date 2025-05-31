package gosimple

import "github.com/fengdotdev/golibs-traits/trait"

var _ trait.CRUD[string, map[string]any] = (*GoSimple[map[string]any])(nil)

// All implements trait.CRUD.
func (g *GoSimple[T]) All() map[string]T {
	panic("unimplemented")
}

// Clean implements trait.CRUD.
func (g *GoSimple[T]) Clean() {
	panic("unimplemented")
}

// Create implements trait.CRUD.
func (g *GoSimple[T]) Create(id string, item T) error {
	panic("unimplemented")
}

// Delete implements trait.CRUD.
func (g *GoSimple[T]) Delete(id string) error {
	panic("unimplemented")
}

// Exists implements trait.CRUD.
func (g *GoSimple[T]) Exists(id string) (bool, error) {
	panic("unimplemented")
}

// Iterate implements trait.CRUD.
func (g *GoSimple[T]) Iterate(fn func(string, T) (stop bool, err error)) error {
	panic("unimplemented")
}

// Keys implements trait.CRUD.
func (g *GoSimple[T]) Keys() []string {
	panic("unimplemented")
}

// Len implements trait.CRUD.
func (g *GoSimple[T]) Len() int {
	panic("unimplemented")
}

// Populate implements trait.CRUD.
func (g *GoSimple[T]) Populate(map[string]T) {
	panic("unimplemented")
}

// Read implements trait.CRUD.
func (g *GoSimple[T]) Read(id string) (T, error) {
	panic("unimplemented")
}

// Update implements trait.CRUD.
func (g *GoSimple[T]) Update(id string, item T) error {
	panic("unimplemented")
}

// Values implements trait.CRUD.
func (g *GoSimple[T]) Values() []T {
	panic("unimplemented")
}
