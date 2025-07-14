package database

import (
	"context"
	"fmt"

	"github.com/umardev500/go-attendance/internal/ent"
)

type contextKeyTx struct{}

type TransactionManager struct {
	client *ent.Client
}

func NewTransactionManager(client *ent.Client) *TransactionManager {
	return &TransactionManager{
		client: client,
	}
}

func (tm *TransactionManager) WithTx(ctx context.Context, fn func(ctx context.Context) error) error {
	tx, err := tm.client.Tx(ctx)
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		}
	}()

	ctx = context.WithValue(ctx, contextKeyTx{}, tx)

	if err := fn(ctx); err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("transaction rollback error: %v, original error: %w", rbErr, err)
		}
		return err
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("committing transaction: %w", err)
	}

	return nil
}

func (tm *TransactionManager) FromContext(ctx context.Context) *ent.Client {
	if tx, ok := ctx.Value(contextKeyTx{}).(*ent.Tx); ok {
		return tx.Client()
	}
	return tm.client
}
