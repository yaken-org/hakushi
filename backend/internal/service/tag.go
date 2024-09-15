package service

import (
	"github.com/yaken-org/hakushi/internal/database"
	"github.com/yaken-org/hakushi/internal/model"
)

func FindTagByID(id int64) (*model.Tag, error) {
	db := database.New()
	tag := new(model.Tag)
	if err := model.QueryRow(db.DB, tag, "SELECT * FROM tags WHERE id = ?", id); err != nil {
		return nil, err
	}
	return tag, nil
}

func FindAllTags() ([]*model.Tag, error) {
	db := database.New()
	rows, err := db.Query("SELECT * FROM tags")
	if err != nil {
		return nil, err
	}

	tags := make([]*model.Tag, 0)
	for rows.Next() {
		tag := new(model.Tag)
		if err := tag.FromRow(rows); err != nil {
			return nil, err
		}
		tags = append(tags, tag)
	}
	return tags, nil
}

func CreateTag(name string) (*model.Tag, error) {
	db := database.New()
	res, err := db.Exec(`
		INSERT INTO tags (name) VALUES(?)
	`, name)
	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	return FindTagByID(id)
}

func DeleteTag(id int64) error {
	db := database.New()
	_, err := db.Exec("DELETE FROM tags WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}

func UpdateTag(id int64, name string) (*model.Tag, error) {
	db := database.New()
	_, err := db.Exec("UPDATE tags SET name = ? WHERE id = ?", name, id)
	if err != nil {
		return nil, err
	}
	return FindTagByID(id)
}
