package like

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ LikedModel = (*customLikedModel)(nil)

type (
	// LikedModel is an interface to be customized, add more methods here,
	// and implement the added methods in customLikedModel.
	LikedModel interface {
		likedModel
	}

	customLikedModel struct {
		*defaultLikedModel
	}
)

// NewLikedModel returns a model for the database table.
func NewLikedModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) LikedModel {
	return &customLikedModel{
		defaultLikedModel: newLikedModel(conn, c, opts...),
	}
}
