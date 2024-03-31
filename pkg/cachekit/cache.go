package cachekit

import (
	"context"
	"time"
)

type Cache interface {
	GetByFunc(context context.Context, prefix, key string, container interface{}, getter OneTimeGetterFunc) error
	Get(context context.Context, prefix, key string, container interface{}) error
	Set(context context.Context, prefix string, key string, value interface{}, duration time.Duration) error
}

// OneTimeGetterFunc should be provided as a parameter in GetByFunc()
type OneTimeGetterFunc func() (interface{}, error)
