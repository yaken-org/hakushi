package service

import (
	"github.com/yaken-org/hakushi/internal/database"
	"github.com/yaken-org/hakushi/internal/model"
)

func FindAllPosts() ([]*model.Post, error) {
	db := database.New()

	res, err := db.Query("SELECT * FROM post")
	if err != nil {
		return nil, err
	}

	posts := make([]*model.Post, 0)
	for res.Next() {
		post := new(model.Post)
		if err := post.FromRow(res); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func FindPostsOrderByLikes() ([]*model.Post, error) {
	db := database.New()

	res, err := db.Query(`
		SELECT *
		FROM post
		WHERE likes > 0
		ORDER BY likes DESC
	`)
	if err != nil {
		return nil, err
	}

	posts := make([]*model.Post, 0)
	for res.Next() {
		post := new(model.Post)
		if err := post.FromRow(res); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func FindPostByID(id int64) (*model.Post, error) {
	db := database.New()

	post := new(model.Post)
	if err := model.QueryRow(db.DB, post, "SELECT * FROM post WHERE id = ?", id); err != nil {
		return nil, err
	}
	return post, nil
}

func FindAPIPostByID(id int64) (*model.APIPost, error) {
	post, err := FindPostByID(id)
	if err != nil {
		return nil, err
	}

	annotations, err := FindAnnotationsByPostID(id)
	if err != nil {
		return nil, err
	}

	tags, err := FindPostRelatedTags(post.ID)
	if err != nil {
		return nil, err
	}

	return post.ToAPIPost(annotations, tags), nil
}

func FindPostByNameRough(name string) ([]*model.Post, error) {
	db := database.New()

	name = "%" + name + "%"
	res, err := db.Query(`SELECT * FROM post WHERE title LIKE ?`, name)
	if err != nil {
		return nil, err
	}

	posts := make([]*model.Post, 0)
	for res.Next() {
		post := new(model.Post)
		if err := post.FromRow(res); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func FindPostsByUserAccountID(userAccountID int64) ([]*model.Post, error) {
	db := database.New()

	res, err := db.Query("SELECT * FROM post WHERE user_account_id = ?", userAccountID)
	if err != nil {
		return nil, err
	}

	posts := make([]*model.Post, 0)
	for res.Next() {
		post := new(model.Post)
		if err := post.FromRow(res); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func CreatePost(userAccount model.UserAccount, imageId int64, title string, content string) (*model.Post, error) {
	db := database.New()

	res, err := db.Exec(`
		INSERT INTO post (user_account_id, image_id, title, content) VALUES (?, ?, ?, ?)
	`, userAccount.ID, imageId, title, content)
	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	return FindPostByID(id)
}

func DeletePost(id int64) error {
	db := database.New()

	_, err := db.Exec("DELETE FROM post WHERE id = ?", id)
	if err != nil {
		return err
	}

	return nil
}

func UpdatePost(id int64, title string, content string) (*model.Post, error) {
	db := database.New()

	_, err := db.Exec("UPDATE post SET title = ?, content = ? WHERE id = ?", title, content, id)
	if err != nil {
		return nil, err
	}

	return FindPostByID(id)
}

func FindPostRelatedTags(postID int64) ([]*model.Tag, error) {
	db := database.New()

	res, err := db.Query(`
		SELECT * FROM tag
		WHERE id IN (
			SELECT tag_id FROM post_tag WHERE post_id = ?
		)
	`, postID)
	if err != nil {
		return nil, err
	}

	tags := make([]*model.Tag, 0)
	for res.Next() {
		tag := new(model.Tag)
		if err := tag.FromRow(res); err != nil {
			return nil, err
		}
		tags = append(tags, tag)
	}

	return tags, nil
}

func FindPostRelatedTagsByPostIDs(postIDs []int64) (map[int64][]*model.Tag, error) {
	// N+1 が起きているが、無視する
	postIdToTags := make(map[int64][]*model.Tag)
	for _, postID := range postIDs {
		tags, err := FindPostRelatedTags(postID)
		if err != nil {
			return nil, err
		}
		postIdToTags[postID] = tags
	}

	return postIdToTags, nil
}

func FindPostsByTag(tag *model.Tag) ([]*model.Post, error) {
	db := database.New()

	res, err := db.Query(`
		SELECT * FROM post
		WHERE id IN (
			SELECT post_id FROM post_tag WHERE tag_id = ?
		)
	`, tag.ID)
	if err != nil {
		return nil, err
	}

	posts := make([]*model.Post, 0)
	for res.Next() {
		post := new(model.Post)
		if err := post.FromRow(res); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func IncreamentPostLikeCount(postID int64) (int, error) {
	db := database.New()

	_, err := db.Exec("UPDATE post SET likes = likes + 1 WHERE id = ?", postID)
	if err != nil {
		return 0, err
	}

	post, err := FindPostByID(postID)
	if err != nil {
		return 0, err
	}

	return post.Likes, nil
}
