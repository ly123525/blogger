package db

import (
	"blogger/model"
	"testing"
	"time"
)

func init() {
	dns := "root:@tcp(localhost:3306)/blogger?parseTime=true"
	err := Init(dns)
	if err != nil {
		panic(err)
	}

}

func TestInsertArticle(t *testing.T) {
	article := &model.ArticleDetail{}
	article.ArticleInfo.CategoryId = 1
	article.ArticleInfo.CommentCount = 0
	article.ArticleInfo.Summary = "this a test ak dkdkdkddkddkd"
	article.ArticleInfo.CreateTime = time.Now()
	article.Content = `使用mysql的时间字段遇到如下两个问题
	1.使用go-sql-driver来连接mysql数据库，获取的时区默认是UTC +0的，与本地的东八区是有区别，在业务处理中会出现问题
	2.获取mysql中的日期，是string类型，需要在代码中用time.Parse进行转化`
	article.ArticleInfo.Title = "GOLANG 连接Mysql的时区问题"
	article.ArticleInfo.Username = "少林之巅"
	article.ArticleInfo.ViewCount = 1
	article.Category.CategoryId = 1
	articleId, err := InsertArticle(article)
	if err != nil {
		t.Errorf("insert article failed, err:%v\n", err)
		return
	}

	t.Logf("insert article succ, articleId:%d\n", articleId)
}

func TestGetArticleList(t *testing.T) {
	articleList, err := GetArticleList(1, 15)
	if err != nil {
		t.Errorf("get article failed, err:%v\n", err)
		return
	}
	t.Logf("fet article succ, len:%d\n", len(articleList))
}
