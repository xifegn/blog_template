package user

import (
	"context"

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
	// todo: add your logic here and delete this line

	return
}
