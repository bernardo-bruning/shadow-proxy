package replication

import (
	"context"
	"log"

	"github.com/bernardo-bruning/shadowproxy/domain"
)

type Log struct {
}

func NewLog() *Log {
	return &Log{}
}

func (l *Log) Emit(ctx context.Context, req *domain.Request) error {
	message, _ := req.ToJson()
	log.Printf("request: %s", message)
	return nil
}
