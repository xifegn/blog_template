package user

import (
	"blog_template/common/cryptx"
	"blog_template/models/user"
	"context"
	"google.golang.org/grpc/status"

	"blog_template/internal/svc"
	"blog_template/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	_, err = l.svcCtx.UserModel.FindOneByEmail(l.ctx, req.Email)
	if err == nil {
		return nil, status.Error(500, "user exist")
	}
	if err == user.ErrNotFound {
		newUser := user.Users{
			Name:     req.Name,
			Username: req.Username,
			Password: cryptx.PasswordEncrypt(l.svcCtx.Config.Salt, req.Password),
			Email:    req.Email,
		}
		res, err := l.svcCtx.UserModel.Insert(l.ctx, &newUser)
		if err != nil {
			return nil, err
		}
		newUser.Id, err = res.LastInsertId()
		return &types.RegisterResp{
			Id:       newUser.Id,
			Name:     newUser.Name,
			Username: newUser.Username,
			Email:    newUser.Email,
		}, nil
	}
	return &types.RegisterResp{}, nil
}
