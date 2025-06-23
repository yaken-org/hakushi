package helper

import (
	"database/sql"
	"database/sql/driver"
	"regexp"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

type AnyTime struct{}

func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

func SetupMockDB() (sqlmock.Sqlmock, *sql.DB, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}
	return mock, db, nil
}

func ExpectFindPostByID(mock sqlmock.Sqlmock, postID int64, userAccountID, imageID int64, title, content string, likes int) {
	rows := sqlmock.NewRows([]string{
		"id", "user_account_id", "image_id", "title", "content", "likes", "created_at", "updated_at",
	}).AddRow(postID, userAccountID, imageID, title, content, likes, time.Now(), time.Now())

	mock.ExpectQuery(regexp.QuoteMeta("SELECT id, user_account_id, image_id, title, content, likes, created_at, updated_at FROM post WHERE id = ?")).
		WithArgs(postID).
		WillReturnRows(rows)
}

func ExpectFindAllPosts(mock sqlmock.Sqlmock, posts [][]interface{}) {
	rows := sqlmock.NewRows([]string{
		"id", "user_account_id", "image_id", "title", "content", "likes", "created_at", "updated_at",
	})

	for _, post := range posts {
		rows.AddRow(post[0], post[1], post[2], post[3], post[4], post[5], post[6], post[7])
	}

	mock.ExpectQuery(regexp.QuoteMeta("SELECT id, user_account_id, image_id, title, content, likes, created_at, updated_at FROM post ORDER BY created_at DESC")).
		WillReturnRows(rows)
}

func ExpectCreatePost(mock sqlmock.Sqlmock, userAccountID, imageID int64, title, content string, insertID int64) {
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO post (user_account_id, image_id, title, content, likes, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?)")).
		WithArgs(userAccountID, imageID, title, content, 0, AnyTime{}, AnyTime{}).
		WillReturnResult(sqlmock.NewResult(insertID, 1))
}

func ExpectUpdatePost(mock sqlmock.Sqlmock, postID int64, title, content string, likes int) {
	mock.ExpectExec(regexp.QuoteMeta("UPDATE post SET title = ?, content = ?, likes = ?, updated_at = ? WHERE id = ?")).
		WithArgs(title, content, likes, AnyTime{}, postID).
		WillReturnResult(sqlmock.NewResult(0, 1))
}

func ExpectDeletePost(mock sqlmock.Sqlmock, postID int64) {
	mock.ExpectExec(regexp.QuoteMeta("DELETE FROM post WHERE id = ?")).
		WithArgs(postID).
		WillReturnResult(sqlmock.NewResult(0, 1))
}