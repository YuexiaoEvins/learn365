package lib

import (
	"context"
	"time"
)

type Caller interface {
	BuildReq() RawReq
	Call(ctx context.Context, req RawReq, timeout time.Duration) (RawResp, error)
	CheckResp()
}
