package service

import (
	"github.com/yaken-org/hakushi/internal/database"
	"github.com/yaken-org/hakushi/internal/model"
)

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
