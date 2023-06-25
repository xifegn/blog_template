package attachment

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AttachmentModel = (*customAttachmentModel)(nil)

type (
	// AttachmentModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAttachmentModel.
	AttachmentModel interface {
		attachmentModel
	}

	customAttachmentModel struct {
		*defaultAttachmentModel
	}
)

// NewAttachmentModel returns a model for the database table.
func NewAttachmentModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) AttachmentModel {
	return &customAttachmentModel{
		defaultAttachmentModel: newAttachmentModel(conn, c, opts...),
	}
}
