//+build ignore
// Code above this line should be ignored as it's not part of the example

// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch10/acme/internal/config"
	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch10/acme/internal/modules/exchange"
	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch10/acme/internal/modules/get"
	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch10/acme/internal/modules/list"
	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch10/acme/internal/modules/register"
	"github.com/PacktPublishing/Hands-On-Dependency-Injection-in-Go/ch10/acme/internal/rest"
)

// Injectors from wire.go:

func initializeServer() (*rest.Server, error) {
	configConfig, err := config.Load()
	if err != nil {
		return nil, err
	}
	getter := get.NewGetter(configConfig)
	lister := list.NewLister(configConfig)
	converter := exchange.NewConverter(configConfig)
	registerer := register.NewRegisterer(configConfig, converter)
	server := rest.New(configConfig, getter, lister, registerer)
	return server, nil
}
