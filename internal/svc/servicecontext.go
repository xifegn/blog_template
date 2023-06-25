package svc

import (
	"blog_template/internal/config"
	"blog_template/models/attachment"
	"blog_template/models/comment"
	"blog_template/models/follow"
	"blog_template/models/image"
	"blog_template/models/like"
	"blog_template/models/summary"
	"blog_template/models/user"
	"github.com/zeromicro/go-zero/core/stores/postgres"
)

type ServiceContext struct {
	Config          config.Config
	UserModel       user.UsersModel
	SummaryModel    summary.SummaryModel
	LikeModel       like.LikedModel
	ImageModel      image.ImageModel
	FollowModel     follow.FollowModel
	CommentModel    comment.CommentModel
	AttachmentModel attachment.AttachmentModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := postgres.New(c.Postgres.DataSource)
	return &ServiceContext{
		Config:          c,
		UserModel:       user.NewUsersModel(conn, c.CacheRedis),
		SummaryModel:    summary.NewSummaryModel(conn, c.CacheRedis),
		LikeModel:       like.NewLikedModel(conn, c.CacheRedis),
		ImageModel:      image.NewImageModel(conn, c.CacheRedis),
		FollowModel:     follow.NewFollowModel(conn, c.CacheRedis),
		CommentModel:    comment.NewCommentModel(conn, c.CacheRedis),
		AttachmentModel: attachment.NewAttachmentModel(conn, c.CacheRedis),
	}
}
