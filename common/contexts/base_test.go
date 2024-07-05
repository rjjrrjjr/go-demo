package contexts

import (
	"context"
	"testing"

	"git.code.oa.com/trpc-go/trpc-go/log"
	"github.com/google/uuid"
)

func TestLog(t *testing.T) {

	ctx := context.WithValue(context.Background(), "traceId", uuid.New().String())

	log.DebugContextf(ctx, "jinrruan: %+v", "CAO")
}
