package context

import (
	"context"
	"time"
)

// Read README.md for the instructions

// TraceKey is a custom type to prevent context key collisions
type TraceKey string

// CDNAPI is just a stub for testing.
type CDNAPI struct {
	LogCtx      context.Context
	DNSCtx      context.Context
	DownloadCtx context.Context
	VerifyCtx   context.Context
}

// These are already implemented for you by the test suite.
func (api *CDNAPI) LogActivity(ctx context.Context)    { api.LogCtx = ctx }
func (api *CDNAPI) DNSLookup(ctx context.Context)      { api.DNSCtx = ctx }
func (api *CDNAPI) DownloadChunks(ctx context.Context) { api.DownloadCtx = ctx }
func (api *CDNAPI) VerifyChecksum(ctx context.Context) { api.VerifyCtx = ctx }

// TODO: Implement RunSmartDownloader
func RunSmartDownloader(api *CDNAPI) {
	// 1. Create root context
	root := context.Background()

	// 2. Attach trace ID for activity logging
	logCtx := context.WithValue(
		root,
		TraceKey("trace_id"),
		"smart-dl-123",
	)
	api.LogActivity(logCtx)

	// 3. DNS lookup with 100ms timeout
	dnsCtx, dnsCancel := context.WithTimeout(
		root,
		100*time.Millisecond,
	)
	defer dnsCancel()

	api.DNSLookup(dnsCtx)

	// 4. Create and immediately cancel download context
	downloadCtx, downloadCancel := context.WithCancel(root)
	downloadCancel()

	api.DownloadChunks(downloadCtx)

	// 5. Use TODO context for legacy checksum verification
	api.VerifyChecksum(context.TODO())
}
