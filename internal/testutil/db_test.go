package testutil_test

import (
	"__MODULE__/internal/testutil"
	"context"
	"testing"
)

func TestWithRollbackTx_StartsTransaction(t *testing.T) {
	tdb := testutil.OpenTestDB(t)

	testutil.WithRollbackTx(t, tdb, func(ctx context.Context) {
		if ctx == nil {
			t.Fatal("expected non-nil context")
		}
	})
}
