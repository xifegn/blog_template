package like

import (
	"context"

	"blog_template/internal/svc"
	"blog_template/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddLikeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddLikeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLikeLogic {
	return &AddLikeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddLikeLogic) AddLike(req *types.AddLikeReq) (resp *types.AddLikeResp, err error) {
	// todo: add your logic here and delete this line

	return
}
