package summary

import (
	"context"

	"blog_template/internal/svc"
	"blog_template/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateSummaryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateSummaryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateSummaryLogic {
	return &CreateSummaryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateSummaryLogic) CreateSummary(req *types.CreateSummaryReq) (resp *types.CreateSummaryResp, err error) {
	// todo: add your logic here and delete this line

	return
}
