package blog_model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ BlogsModel = (*customBlogsModel)(nil)

type (
	// BlogsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBlogsModel.
	BlogsModel interface {
		blogsModel
		withSession(session sqlx.Session) BlogsModel
	}

	customBlogsModel struct {
		*defaultBlogsModel
	}
)

// NewBlogsModel returns a model for the database table.
func NewBlogsModel(conn sqlx.SqlConn) BlogsModel {
	return &customBlogsModel{
		defaultBlogsModel: newBlogsModel(conn),
	}
}

func (m *customBlogsModel) withSession(session sqlx.Session) BlogsModel {
	return NewBlogsModel(sqlx.NewSqlConnFromSession(session))
}
