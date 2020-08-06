//+build wireinject

package main

import (
	"github.com/google/wire"

	"github.com/feiquan123/go-program-framework/src/app/ops"
	"github.com/feiquan123/go-program-framework/src/pkg/app"
	"github.com/feiquan123/go-program-framework/src/pkg/config"
	"github.com/feiquan123/go-program-framework/src/pkg/log"
)

var providerSet = wire.NewSet(
	log.ProviderSet,
	config.ProviderSet,
	ops.ProviderSet,
)

func NewAPP(config string) (*app.Application, error) {
	panic(wire.Build(providerSet))
}
