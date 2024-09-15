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

	return post.ToAPIPost(annotations), nil
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
