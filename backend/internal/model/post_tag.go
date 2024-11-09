package model

type PostTag struct {
	PostID int64 `json:"post_id" gorm:"primaryKey"`
	TagID  int64 `json:"tag_id" gorm:"primaryKey"`
}

func (p *PostTag) FromRow(row Scannable) error {
	return row.Scan(
		&p.PostID,
		&p.TagID,
	)
}

func (p *PostTag) TableName() string {
	return "post_tag"
}
