package common

import (
	"go.uber.org/zap"
	"sync"
)

var (
	Logger *zap.Logger
	Once   sync.Once
)
