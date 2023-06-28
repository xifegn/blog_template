package summary

import (
	"blog_template/internal/svc"
	"blog_template/internal/types"
	"blog_template/models/summary"
	"context"
	"google.golang.org/grpc/status"
	"time"

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

	res, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, req.Username)
	if err != nil {
		return nil, status.Error(500, "user not found")
	}
	newSummary := summary.Summary{
		UserId:   res.Id,
		Date:     time.Time{},
		Title:    req.Title,
		Content:  req.Content,
		IsShared: false,
	}
	_, err = l.svcCtx.SummaryModel.Insert(l.ctx, &newSummary)
	if err != nil {
		return nil, status.Error(500, "insert failed")
	}
	return &types.CreateSummaryResp{Code: 200}, nil
}
