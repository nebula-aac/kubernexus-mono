//go:build wireinject
// +build wireinject

package main

import "github.com/google/wire"

func InitializeApp() NexusApp {
	wire.Build(NewEcho, NewNexusApp)
	return NexusApp{}
}
