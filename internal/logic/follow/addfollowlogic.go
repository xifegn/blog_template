package follow

import (
	"context"

	"blog_template/internal/svc"
	"blog_template/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddFollowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddFollowLogic {
	return &AddFollowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddFollowLogic) AddFollow(req *types.AddFollowReq) (resp *types.AddFollowResp, err error) {

	return
}
