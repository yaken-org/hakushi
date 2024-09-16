package server

import "github.com/yaken-org/hakushi/internal/server/handler"

func (s *Server) configureRoute() {
	e := s.Engine

	e.GET("/health", handler.Health)

	api := e.Group("/api")
	api.POST("/account", handler.CreateUserAccount)     // アカウント作成
	api.GET("/account/:id", handler.GetUserAcocunt)     // アカウントの詳細取得
	api.GET("/account/:id/posts", handler.GetUserPosts) // アカウントの投稿一覧取得

	api.GET("/post", handler.GetAllPosts)              // 投稿一覧取得
	api.POST("/post", handler.CreatePost)              // 投稿作成
	api.GET("/post/:id", handler.GetPost)              // 投稿の詳細取得
	api.GET("/post/:id/tags", handler.GetPostTags)     // 投稿のタグ一覧取得
	api.POST("/post/:id/like", handler.SendLikeToPost) // 投稿にタグを追加

	api.GET("/tag", handler.GetAllTags)               // タグ一覧取得
	api.POST("/tag", handler.CreateTag)               // タグ作成
	api.GET("/tag/:id/posts", handler.GetTaggedPosts) // タグがついた投稿一覧取得

	api.GET("/ranking", handler.GetRanking) // ランキング取得

	api.GET("/search", handler.Search) // 検索
}
