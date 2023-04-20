package user

import (
	"blog_template/common/cryptx"
	"blog_template/common/jwtx"
	"blog_template/models/user"
	"context"
	"google.golang.org/grpc/status"
	"time"

	"blog_template/internal/svc"
	"blog_template/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	res, err := l.svcCtx.UserModel.FindOneByMobile(l.ctx, req.Mobile)
	if err != nil {
		if err == user.ErrNotFound {
			return nil, status.Error(100, "user not found")
		}
		return nil, status.Error(500, err.Error())
	}
	password := cryptx.PasswordEncrypt(l.svcCtx.Config.Salt, req.Password)
	if password != res.Password {
		return nil, status.Error(100, "password error")
	}
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.Auth.AccessExpire
	accessToken, err := jwtx.GetToken(l.svcCtx.Config.Auth.AccessSecret, now, accessExpire, res.Id)
	if err != nil {
		return nil, err
	}

	return &types.LoginResp{
		AccessToken:  accessToken,
		AccessExpire: now + accessExpire,
	}, nil
}
