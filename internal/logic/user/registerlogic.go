package user

import (
	"blog_template/common/cryptx"
	"blog_template/internal/svc"
	"blog_template/internal/types"
	"blog_template/models/user"
	"context"
	"google.golang.org/grpc/status"

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

func (l *RegisterLogic) Register(req *types.RegisterReq) (*types.RegisterResp, error) {
	_, err := l.svcCtx.UserModel.FindOneByMobile(l.ctx, req.Mobile)
	if err == nil {
		return nil, status.Error(100, "user has been registered")
	}
	if err == user.ErrNotFound {
		newUser := user.Users{
			Name:     req.Name,
			Mobile:   req.Mobile,
			Password: cryptx.PasswordEncrypt(l.svcCtx.Config.Salt, req.Password),
		}
		res, err := l.svcCtx.UserModel.Insert(l.ctx, &newUser)
		if err != nil {
			return nil, status.Error(500, err.Error())
		}
		newUser.Id, err = res.LastInsertId()
		resp := user.Users{
			Name:   req.Name,
			Mobile: req.Mobile,
		}
		return &types.RegisterResp{
			Name:   resp.Name,
			Mobile: resp.Mobile,
		}, nil
	}
	return &types.RegisterResp{}, nil
}
