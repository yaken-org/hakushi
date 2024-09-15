package service

import (
	"github.com/yaken-org/hakushi/internal/database"
	"github.com/yaken-org/hakushi/internal/model"
)

func FindAnnotationByID(id int64) (*model.Annotation, error) {
	db := database.New()

	annotation := new(model.Annotation)
	if err := model.QueryRow(db.DB, annotation, "SELECT * FROM annotation WHERE id = ?", id); err != nil {
		return nil, err
	}

	return annotation, nil
}

func FindAnnotationsByPostID(postID int64) ([]*model.Annotation, error) {
	db := database.New()

	rows, err := db.Query("SELECT * FROM annotation WHERE post_id = ?", postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	annotations := []*model.Annotation{}
	for rows.Next() {
		annotation := new(model.Annotation)
		if err := annotation.FromRow(rows); err != nil {
			return nil, err
		}
		annotations = append(annotations, annotation)
	}

	return annotations, nil
}

func FindAnnotationsByPostIDs(postIDs []int64) (map[int64][]*model.Annotation, error) {
	db := database.New()

	query := "SELECT * FROM annotation WHERE post_id IN ("
	for i := range postIDs {
		if i > 0 {
			query += ", "
		}
		query += "?"
	}
	query += ")"

	args := make([]interface{}, len(postIDs))
	for i, id := range postIDs {
		args[i] = id
	}

	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	annotations := []*model.Annotation{}
	for rows.Next() {
		annotation := new(model.Annotation)
		if err := annotation.FromRow(rows); err != nil {
			return nil, err
		}
		annotations = append(annotations, annotation)
	}

	// PostID to Annotations Map
	postIDToAnnotations := map[int64][]*model.Annotation{}
	for _, annotation := range annotations {
		postIDToAnnotations[annotation.PostID] = append(postIDToAnnotations[annotation.PostID], annotation)
	}

	return postIDToAnnotations, nil
}

func CreateAnnotation(post *model.Post, productID int64, DisplayName string, x int64, y int64) (*model.Annotation, error) {
	db := database.New()

	res, err := db.Exec(`
		INSERT INTO annotation (post_id, product_id, display_name, x, y) VALUES (?, ?, ?, ?, ?)
	`, post.ID, productID, DisplayName, x, y)
	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	return FindAnnotationByID(id)
}
