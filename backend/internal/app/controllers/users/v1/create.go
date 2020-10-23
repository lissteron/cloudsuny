// Code generated by protoc-gen-goclay, but your can (must) modify it.
// source: users.proto

package users

import (
	"context"

	desc "github.com/lissteron/cloudsuny/pkg/users/v1"
)

func (i *Implementation) Create(
	ctx context.Context,
	req *desc.CreateReq,
	) (resp *desc.CreateResp, err error) {
	if err := req.Validate(); err != nil {
		i.logger.Warnf(ctx, "validate create user request failed: %v", err)

		return nil, err
	}

	result, err := i.create.Do(ctx, req.ToDomain())
	if err != nil {
		i.logger.Errorf(ctx,"create user failed: %v", err)

		return nil, err
	}

	resp = &desc.CreateResp{
		Data: desc.UserFromDomain(result),
	}

	return resp, nil
}
