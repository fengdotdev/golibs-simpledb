package types

import "github.com/fengdotdev/golibs-traits/trait"





type SimpleDB[T any] interface {
	trait.CRUD[string, T]
	trait.Saver
	trait.MapSerializer
	trait.Binarer
}

// All implements trait.C
