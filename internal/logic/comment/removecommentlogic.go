package comment

import (
	"context"

	"blog_template/internal/svc"
	"blog_template/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveCommentLogic {
	return &RemoveCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveCommentLogic) RemoveComment(req *types.RemoveCommentReq) (resp *types.RemoveCommentResp, err error) {
	// todo: add your logic here and delete this line

	return
}
