// Package model 逻辑处理层
package model

import "github.com/google/wire"

// WireModelSet 控制器 DI
var WireModelSet = wire.NewSet(
	LoginModelSet,
	PaperModelSet,
)
