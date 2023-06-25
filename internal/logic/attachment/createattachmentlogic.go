package attachment

import (
	"context"

	"blog_template/internal/svc"
	"blog_template/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateAttachmentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateAttachmentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateAttachmentLogic {
	return &CreateAttachmentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateAttachmentLogic) CreateAttachment(req *types.CreateAttachmentReq) (resp *types.CreateAttachmentResp, err error) {
	// todo: add your logic here and delete this line

	return
}
