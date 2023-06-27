package user

import (
	"context"
	"encoding/json"
	"google.golang.org/grpc/status"

	"blog_template/internal/svc"
	"blog_template/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo() (resp *types.UserInfoResp, err error) {
	uid, _ := l.ctx.Value("uid").(json.Number).Int64()
	res, err := l.svcCtx.UserModel.FindOne(l.ctx, uid)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}
	return &types.UserInfoResp{
		Id:        res.Id,
		Name:      res.Name,
		Username:  res.Username,
		Email:     res.Email,
		Avator:    res.Avatar,
		Signature: res.Signature,
	}, nil
}
