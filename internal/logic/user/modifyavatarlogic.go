package user

import (
	"blog_template/models/user"
	"context"
	"google.golang.org/grpc/status"

	"blog_template/internal/svc"
	"blog_template/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ModifyAvatarLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewModifyAvatarLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ModifyAvatarLogic {
	return &ModifyAvatarLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}
func (l *ModifyAvatarLogic) ModifyAvatar(req *types.ModifyAvatarReq) (resp *types.ModifyAvatarResp, err error) {
	res, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, req.Username)
	if err != nil {
		return nil, status.Error(500, "")
	}

	newUser := user.Users{
		Id:        res.Id,
		Username:  res.Username,
		Password:  res.Password,
		Email:     res.Email,
		Avatar:    req.Avatar,
		Signature: res.Signature,
		Name:      res.Name,
	}
	err = l.svcCtx.UserModel.Update(l.ctx, &newUser)
	if err != nil {
		return nil, status.Error(500, "update failed")
	}

	return &types.ModifyAvatarResp{Code: 200}, nil
}
