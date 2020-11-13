// Code generated by protoc-gen-goclay, but your can (must) modify it.
// source: badges.proto

package badges

import (
	"context"

	desc "github.com/lissteron/cloudsuny/pkg/badges/v1"
	types "github.com/lissteron/cloudsuny/pkg/types/v1"
)

func (i *Implementation) Create(
	ctx context.Context,
	req *desc.CreateBadgesReq,
) (resp *desc.CreateBadgesResp, err error) {
	if err := req.Validate(); err != nil {
		i.logger.Warnf(ctx,"validate create badge request failed: %v", err)

		return nil, err
	}

	result, err := i.create.Do(ctx, req.ToDomain())
	if err != nil {
		i.logger.Errorf(ctx, "create badge failed: %v", err)

		return nil, err
	}

	resp = &desc.CreateBadgesResp{
		Data: types.BadgeFromDomain(result),
	}

	return resp, nil
}
