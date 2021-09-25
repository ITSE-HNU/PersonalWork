package service

import "github.com/google/wire"

// WireServiceSet 控制器 DI
var WireServiceSet = wire.NewSet(
	LoginServiceSet,
	PaperServiceSet,
)
