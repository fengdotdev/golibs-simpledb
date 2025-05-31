package gosimple

import "github.com/fengdotdev/golibs-traits/trait/query"

var _ query.QueryableWithCTX = (*Queryable)(nil)

type GoSimple struct {
	data map[string]any
}

func NewSimpleDB(data map[string]any) *GoSimple {
	return &GoSimple{data: data}
}
