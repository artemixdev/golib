package middleware

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

const TxKey Key = "tx"

func Tx(db *sql.DB) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		tx, err := db.Begin()
		if err != nil {
			return fmt.Errorf("tx middleware: begin: %w", err)
		}
		defer tx.Rollback() //nolint:errcheck

		newCtx := context.WithValue(ctx.UserContext(), TxKey, tx)
		ctx.SetUserContext(newCtx)

		if err := ctx.Next(); err != nil {
			return err
		}

		if err := tx.Commit(); err != nil {
			return fmt.Errorf("tx middleware: commit: %w", err)
		}

		return nil
	}
}
