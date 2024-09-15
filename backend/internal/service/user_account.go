package service

import (
	"log/slog"
	"strconv"

	"github.com/yaken-org/hakushi/internal/database"
	"github.com/yaken-org/hakushi/internal/model"
)

func FindUserAccountById(id int64) (*model.UserAccount, error) {
	db := database.New()

	slog.Info("CALL FindUserAccountById")
	user := new(model.UserAccount)
	if err := model.QueryRow(db.DB, user, "SELECT * FROM user_account WHERE id = ?", id); err != nil {
		slog.Error(err.Error())
		return nil, err
	}

	return user, nil
}

func FindUserAccountByName(name string) (*model.UserAccount, error) {
	db := database.New()

	user := new(model.UserAccount)
	if err := model.QueryRow(db.DB, user, "SELECT * FROM user_account WHERE name = ?", name); err != nil {
		return nil, err
	}

	return user, nil
}

func FindUserAccountBySub(sub string) (*model.UserAccount, error) {
	db := database.New()
	subnumber, err := strconv.ParseInt(sub, 10, 64)
	if err != nil {
		return nil, err
	}
	user := new(model.UserAccount)
	if err := model.QueryRow(db.DB, user, "SELECT * FROM user_account WHERE sub = ?", subnumber); err != nil {
		return nil, err
	}

	return user, nil
}

func CreateUserAccount(name string, display_name string, icon_url string, sub string) (*model.UserAccount, error) {
	db := database.New()

	res, err := db.Exec(`
		INSERT INTO user_account (name, display_name, icon_url, sub) VALUES(?, ?, ?, ?)
	`, name, display_name, icon_url, sub)
	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	return FindUserAccountById(id)
}
