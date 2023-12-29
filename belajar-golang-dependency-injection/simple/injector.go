//go:build wireinject
// +build wireinject

package simple

import (
	"github.com/google/wire"
)

func InitializedService() *SimpleServices {
	wire.Build(
		NewSimpleRepository, NewSimpleService,
	)
	return nil
}
