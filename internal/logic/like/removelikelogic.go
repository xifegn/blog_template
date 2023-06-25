package like

import (
	"context"

	"blog_template/internal/svc"
	"blog_template/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveLikeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveLikeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveLikeLogic {
	return &RemoveLikeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveLikeLogic) RemoveLike(req *types.RemoveLikedReq) (resp *types.RemoveLikedResp, err error) {
	// todo: add your logic here and delete this line

	return
}
