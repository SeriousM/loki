package dslog

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/weaveworks/common/middleware"

	"github.com/grafana/dskit/tenant"
)

// WithTraceID returns a Logger that has information about the traceID in
// its details.
func WithTraceID(traceID string, l log.Logger) log.Logger {
	// See note in WithContext.
	return log.With(l, "traceID", traceID)
}

// WithContext returns a Logger that has information about the current user in
// its details.
//
// e.g.
//   log := util.WithContext(ctx)
//   log.Errorf("Could not chunk chunks: %v", err)
func WithContext(ctx context.Context, l log.Logger) log.Logger {
	// Weaveworks uses "orgs" and "orgID" to represent Cortex users,
	// even though the code-base generally uses `userID` to refer to the same thing.
	userID, err := tenant.ID(ctx)
	if err == nil {
		l = WithUserID(userID, l)
	}

	traceID, ok := middleware.ExtractTraceID(ctx)
	if !ok {
		return l
	}

	return WithTraceID(traceID, l)
}
