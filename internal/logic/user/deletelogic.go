package user

import (
	"context"

	"blog_template/internal/svc"
	"blog_template/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteLogic {
	return &DeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteLogic) Delete(req *types.DelReq) (*types.DelResp, error) {
	err := l.svcCtx.UserModel.Delete(l.ctx, req.Id)
	if err == nil {
		return &types.DelResp{
			Code: 0,
		}, nil
	}
	return &types.DelResp{
		Code: 400,
	}, nil
}
