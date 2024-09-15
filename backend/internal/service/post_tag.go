package service

import (
	"github.com/yaken-org/hakushi/internal/database"
	"github.com/yaken-org/hakushi/internal/model"
)

func FindPostTag(post_id, tag_id int64) (*model.PostTag, error) {
	db := database.New()
	postTag := new(model.PostTag)
	if err := model.QueryRow(db.DB, postTag, "SELECT * FROM post_tag WHERE post_id = ? AND tag_id = ?", post_id, tag_id); err != nil {
		return nil, err
	}
	return postTag, nil
}

func FindPostTagsByPostID(postID int64) ([]*model.PostTag, error) {
	db := database.New()

	res, err := db.Query("SELECT * FROM post_tag WHERE post_id = ?", postID)
	if err != nil {
		return nil, err
	}

	postTags := make([]*model.PostTag, 0)
	for res.Next() {
		postTag := new(model.PostTag)
		if err := postTag.FromRow(res); err != nil {
			return nil, err
		}
		postTags = append(postTags, postTag)
	}

	return postTags, nil
}

func FindPostTagsByTagID(tagID int64) ([]*model.PostTag, error) {
	db := database.New()

	res, err := db.Query("SELECT * FROM post_tag WHERE tag_id = ?", tagID)
	if err != nil {
		return nil, err
	}

	postTags := make([]*model.PostTag, 0)
	for res.Next() {
		postTag := new(model.PostTag)
		if err := postTag.FromRow(res); err != nil {
			return nil, err
		}
		postTags = append(postTags, postTag)
	}

	return postTags, nil
}

func CreatePostTag(postID, tagID int64) (*model.PostTag, error) {
	db := database.New()
	_, err := db.Exec(`
		INSERT INTO post_tag (post_id, tag_id) VALUES(?, ?)
	`, postID, tagID)
	if err != nil {
		return nil, err
	}

	return FindPostTag(postID, tagID)
}

func DeletePostTag(postID, tagID int64) error {
	db := database.New()
	_, err := db.Exec("DELETE FROM post_tag WHERE post_id = ? AND tag_id = ?", postID, tagID)
	if err != nil {
		return err
	}
	return nil
}

func DeletePostTagsByPostID(postID int64) error {
	db := database.New()
	_, err := db.Exec("DELETE FROM post_tag WHERE post_id = ?", postID)
	if err != nil {
		return err
	}
	return nil
}

func DeletePostTagsByTagID(tagID int64) error {
	db := database.New()
	_, err := db.Exec("DELETE FROM post_tag WHERE tag_id = ?", tagID)
	if err != nil {
		return err
	}
	return nil
}
