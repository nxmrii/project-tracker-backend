package task

import (
	"context"
)

type Cache interface {
	Delete(
		ctx context.Context,
		key string,
	) error
}
