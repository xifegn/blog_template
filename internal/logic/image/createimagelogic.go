package image

import (
	"context"

	"blog_template/internal/svc"
	"blog_template/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateImageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateImageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateImageLogic {
	return &CreateImageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateImageLogic) CreateImage(req *types.CreateImageReq) (resp *types.CreateImageResp, err error) {
	// todo: add your logic here and delete this line

	return
}
