// Generated by tron. You must modify it.
package users

import (
	context "context"

	"github.com/lissteron/simplerr"

	usersV1 "github.com/lissteron/cloudsuny/pkg/users/v1"
)

func (i *Implementation) Delete(
	ctx context.Context,
	req *usersV1.DeleteReq,
) (resp *usersV1.DeleteResp, err error) {
	if err := req.Validate(); err != nil {
		return nil, simplerr.Wrap(err, "invalid request")
	}

	if err := i.service.Delete(ctx, req.UserId); err != nil {
		return nil, simplerr.Wrap(err, "can't delete user")
	}

	return &usersV1.DeleteResp{}, nil
}
