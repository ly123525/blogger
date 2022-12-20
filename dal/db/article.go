package db

import (
	_ "github.com/go-sql-driver/mysql"

	"blogger/model"
)

func InsertArticle(article *model.ArticleDetail) (articleId int64, err error) {
	sqlstr := "insert into article(content, summary, title, username, category_id, view_count, comment_count)values(?, ?, ?, ?, ?, ?, ?)"
	result, err := DB.Exec(sqlstr, article.Content, article.Summary,
		article.Title, article.Username, article.ArticleInfo.CategoryId,
		article.ArticleInfo.ViewCount, article.ArticleInfo.CommentCount)
	if err != nil {
		return
	}

	articleId, err = result.LastInsertId()
	return
}
