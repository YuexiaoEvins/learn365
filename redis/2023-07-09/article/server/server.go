package server

import (
	"context"
	"learn365/redis/2023-07-09/article/protocol"
)

type ArticleServer struct {
	protocol.UnimplementedArticleServiceServer
}

func (a *ArticleServer) PostArticle(ctx context.Context, request *protocol.PostArticleRequest) (*protocol.PostArticleResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (a *ArticleServer) VoteArticle(ctx context.Context, request *protocol.VoteArticleRequest) (*protocol.VoteArticleResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (a *ArticleServer) GetArticles(ctx context.Context, request *protocol.GetArticlesRequest) (*protocol.GetArticlesResponse, error) {
	//TODO implement me
	panic("implement me")
}
