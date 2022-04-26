package replication

import (
	"context"

	"github.com/bernardo-bruning/shadowproxy/domain"
)

type Replica interface {
	Emit(ctx context.Context, req *domain.Request) error
}
