package dao

import "learn365/redis/2023-07-09/article/protocol"

type PostArticleInfo struct {
}

type ArticleDao interface {
	PostArticle(request protocol.PostArticleRequest)
}
