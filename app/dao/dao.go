package dao

import "github.com/google/wire"

var WireDaoSet = wire.NewSet(
	UserDaoSet,
)
