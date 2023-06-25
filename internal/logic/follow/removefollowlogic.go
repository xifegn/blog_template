package follow

import (
	"context"

	"blog_template/internal/svc"
	"blog_template/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveFollowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveFollowLogic {
	return &RemoveFollowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveFollowLogic) RemoveFollow(req *types.RemoveFollowReq) (resp *types.RemoveFollowResp, err error) {
	// todo: add your logic here and delete this line

	return
}
