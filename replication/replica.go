package replication

import (
	"context"
	"shadowproxy/domain"
)

type Replica interface {
	Emit(ctx context.Context, req *domain.Request) error
}
