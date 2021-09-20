package model

import "github.com/google/wire"

var WireModelSet = wire.NewSet(
	LoginModelSet,
)
