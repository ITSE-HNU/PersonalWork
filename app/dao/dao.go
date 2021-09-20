package dao

import "github.com/google/wire"

// WireDaoSet 控制器 DI
var WireDaoSet = wire.NewSet(
	UserDaoSet,
)
