package gosimple

import "github.com/fengdotdev/golibs-traits/trait"

var _ trait.Saver = (*GoSimple[map[string]any])(nil)

// Load implements trait.Saver.
func (g *GoSimple[T]) Load() error {
	panic("unimplemented")
}

// LoadFrom implements trait.Saver.
func (g *GoSimple[T]) LoadFrom(path string) error {
	panic("unimplemented")
}

// Save implements trait.Saver.
func (g *GoSimple[T]) Save() error {
	panic("unimplemented")
}

// SaveTo implements trait.Saver.
func (g *GoSimple[T]) SaveTo(path string) error {
	panic("unimplemented")
}
