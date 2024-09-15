package service

import (
	"slices"

	"github.com/yaken-org/hakushi/internal/database"
	"github.com/yaken-org/hakushi/internal/model"
)

func FindTagByID(id int64) (*model.Tag, error) {
	db := database.New()
	tag := new(model.Tag)
	if err := model.QueryRow(db.DB, tag, "SELECT * FROM tag WHERE id = ?", id); err != nil {
		return nil, err
	}
	return tag, nil
}

func FindTagByName(name string) (*model.Tag, error) {
	db := database.New()
	tag := new(model.Tag)
	if err := model.QueryRow(db.DB, tag, "SELECT * FROM tag WHERE name = ?", name); err != nil {
		return nil, err
	}
	return tag, nil
}

func FindAllTags() ([]*model.Tag, error) {
	db := database.New()
	rows, err := db.Query("SELECT * FROM tag")
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
		INSERT INTO tag (name) VALUES(?)
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

func CreateTagAndPostTagMapping(postID int64, tags []*model.Tag) ([]*model.Tag, error) {
	// 既存タグを取得
	allTags, err := FindAllTags()
	if err != nil {
		return nil, err
	}
	allTagNames := make([]string, len(allTags))
	for i, tag := range allTags {
		allTagNames[i] = tag.Name
	}
	nameToTag := make(map[string]*model.Tag)
	for _, tag := range allTags {
		nameToTag[tag.Name] = tag
	}

	relatedTags, err := FindPostRelatedTags(postID)
	if err != nil {
		return nil, err
	}
	relatedTagNames := make([]string, len(relatedTags))
	for i, tag := range relatedTags {
		relatedTagNames[i] = tag.Name
	}

	includeTags := make([]*model.Tag, 0)
	for _, tag := range tags {
		t := nameToTag[tag.Name]

		// タグが存在しない場合は作成
		if !slices.Contains(allTagNames, tag.Name) {
			t, err = CreateTag(tag.Name)
			if err != nil {
				return nil, err
			}
		}

		// タグとポストのマッピングを作成
		if !slices.Contains(relatedTagNames, t.Name) {
			_, err := CreatePostTag(postID, t.ID)
			if err != nil {
				return nil, err
			}
		}
		includeTags = append(includeTags, t)
	}
	return includeTags, nil
}

func DeleteTag(id int64) error {
	db := database.New()
	_, err := db.Exec("DELETE FROM tag WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}

func UpdateTag(id int64, name string) (*model.Tag, error) {
	db := database.New()
	_, err := db.Exec("UPDATE tag SET name = ? WHERE id = ?", name, id)
	if err != nil {
		return nil, err
	}
	return FindTagByID(id)
}
