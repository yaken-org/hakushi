package model

type PostTag struct {
	PostID int64 `json:"post_id"`
	TagID  int64 `json:"tag_id"`
}

func (p *PostTag) FromRow(row Scannable) error {
	return row.Scan(
		&p.PostID,
		&p.TagID,
	)
}
