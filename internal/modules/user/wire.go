package user

import "github.com/google/wire"

var Set = wire.NewSet(
	NewHandler,
	NewService,
	NewRepository,
)
