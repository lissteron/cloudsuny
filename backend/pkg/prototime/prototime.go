package prototime

import (
	"time"

	"github.com/gogo/protobuf/types"
)

func TimeToProto(t time.Time) *types.Timestamp {
	if t.IsZero() {
		return nil
	}
	return &types.Timestamp{
		Seconds: t.Unix(),
		Nanos:   int32(t.Nanosecond()),
	}
}

func TimeFromProto(t *types.Timestamp) time.Time {
	if t == nil {
		return time.Time{}
	}
	return time.Unix(t.Seconds, int64(t.Nanos))
}

