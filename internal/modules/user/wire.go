package user

import "github.com/google/wire"

var UserSet = wire.NewSet(
	NewHandler,
	NewService,
	NewRepository,
)
