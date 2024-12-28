// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3

package types

type Article struct {
	Id        int64  `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	UserId    int64  `json:"user_id"`
	ViewCount int64  `json:"view_count"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Author    User   `json:"author"`
}

type ArticleListReq struct {
	Page     int64 `form:"page,default=1"`
	PageSize int64 `form:"page_size,default=10"`
}

type ArticleListResp struct {
	Total    int64     `json:"total"`
	Articles []Article `json:"articles"`
}

type ArticleReq struct {
	Id int64 `path:"id"`
}

type ArticleResp struct {
	Article Article `json:"article"`
}

type Comment struct {
	Id        int64  `json:"id"`
	ArticleId int64  `json:"article_id"`
	UserId    int64  `json:"user_id"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
	Author    User   `json:"author"`
}

type CommentListReq struct {
	ArticleId int64 `path:"article_id"`
	Page      int64 `form:"page,default=1"`
	PageSize  int64 `form:"page_size,default=10"`
}

type CommentListResp struct {
	Total    int64     `json:"total"`
	Comments []Comment `json:"comments"`
}

type CreateArticleReq struct {
	Title   string `json:"title" validate:"required,min=1,max=255"`
	Content string `json:"content" validate:"required"`
}

type CreateArticleResp struct {
	Id int64 `json:"id"`
}

type CreateCommentReq struct {
	ArticleId int64  `path:"article_id"`
	Content   string `json:"content" validate:"required"`
}

type DeleteArticleReq struct {
	Id int64 `path:"id"`
}

type DeleteCommentReq struct {
	Id int64 `path:"id"`
}

type LoginReply struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}

type LoginReq struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type RegisterReq struct {
	Username string `json:"username" validate:"required,min=4,max=32"`
	Password string `json:"password" validate:"required,min=6,max=32"`
}

type UpdateArticleReq struct {
	Id      int64  `path:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type User struct {
	Id        int64  `json:"id"`
	Username  string `json:"username"`
	CreatedAt string `json:"created_at"`
}
