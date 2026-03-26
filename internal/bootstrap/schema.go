package bootstrap

import (
	"__MODULE__/internal/store"
	"context"
	"fmt"
	"time"
)

func CheckSchemaVersion(store *store.Store, required int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	current, err := store.CurrentSchemaVersion(ctx)
	if err != nil {
		return fmt.Errorf("read schema version: %w", err)
	}

	if current != required {
		return fmt.Errorf(
			"schema version mismatch: current=%d required=%d",
			current,
			required,
		)
	}

	return nil
}
