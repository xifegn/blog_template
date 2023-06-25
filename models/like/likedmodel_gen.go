// Code generated by goctl. DO NOT EDIT.

package like

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	likedFieldNames          = builder.RawFieldNames(&Liked{}, true)
	likedRows                = strings.Join(likedFieldNames, ",")
	likedRowsExpectAutoSet   = strings.Join(stringx.Remove(likedFieldNames, "id"), ",")
	likedRowsWithPlaceHolder = builder.PostgreSqlJoin(stringx.Remove(likedFieldNames, "id"))

	cachePublicLikedIdPrefix = "cache:public:liked:id:"
)

type (
	likedModel interface {
		Insert(ctx context.Context, data *Liked) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Liked, error)
		Update(ctx context.Context, data *Liked) error
		Delete(ctx context.Context, id int64) error
	}

	defaultLikedModel struct {
		sqlc.CachedConn
		table string
	}

	Liked struct {
		Id        int64 `db:"id"`
		SummaryId int64 `db:"summary_id"`
		UserId    int64 `db:"user_id"`
	}
)

func newLikedModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultLikedModel {
	return &defaultLikedModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      `"public"."liked"`,
	}
}

func (m *defaultLikedModel) withSession(session sqlx.Session) *defaultLikedModel {
	return &defaultLikedModel{
		CachedConn: m.CachedConn.WithSession(session),
		table:      `"public"."liked"`,
	}
}

func (m *defaultLikedModel) Delete(ctx context.Context, id int64) error {
	publicLikedIdKey := fmt.Sprintf("%s%v", cachePublicLikedIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where id = $1", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, publicLikedIdKey)
	return err
}

func (m *defaultLikedModel) FindOne(ctx context.Context, id int64) (*Liked, error) {
	publicLikedIdKey := fmt.Sprintf("%s%v", cachePublicLikedIdPrefix, id)
	var resp Liked
	err := m.QueryRowCtx(ctx, &resp, publicLikedIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where id = $1 limit 1", likedRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultLikedModel) Insert(ctx context.Context, data *Liked) (sql.Result, error) {
	publicLikedIdKey := fmt.Sprintf("%s%v", cachePublicLikedIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values ($1, $2)", m.table, likedRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.SummaryId, data.UserId)
	}, publicLikedIdKey)
	return ret, err
}

func (m *defaultLikedModel) Update(ctx context.Context, data *Liked) error {
	publicLikedIdKey := fmt.Sprintf("%s%v", cachePublicLikedIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where id = $1", m.table, likedRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.Id, data.SummaryId, data.UserId)
	}, publicLikedIdKey)
	return err
}

func (m *defaultLikedModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cachePublicLikedIdPrefix, primary)
}

func (m *defaultLikedModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where id = $1 limit 1", likedRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultLikedModel) tableName() string {
	return m.table
}
