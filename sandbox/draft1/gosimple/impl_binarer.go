package gosimple

import "github.com/fengdotdev/golibs-traits/trait"

var _ trait.Binarer = (*GoSimple[map[string]any])(nil)

// FromBinary implements trait.Binarer.
func (g *GoSimple[T]) FromBinary([]byte) error {
	panic("unimplemented")
}

// ToBinary implements trait.Binarer.
func (g *GoSimple[T]) ToBinary() ([]byte, error) {
	panic("unimplemented")
}
