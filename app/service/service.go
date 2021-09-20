package service

import "github.com/google/wire"

var WireServiceSet = wire.NewSet(
	LoginModelSet,
)
