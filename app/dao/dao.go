// Package dao 数据库操作层
package dao

import "github.com/google/wire"

// WireDaoSet 控制器 DI
var WireDaoSet = wire.NewSet(
	UserDaoSet,
	TitleDaoSet,
)
