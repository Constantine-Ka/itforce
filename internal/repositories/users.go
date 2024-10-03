package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
)

func GetBalanceById(ctx context.Context, db *sqlx.DB, id int) (float64, error) {
	var balance float64
	query := fmt.Sprintf("SELECT users.balance FROM users WHERE id=%d", id)
	err := db.GetContext(ctx, &balance, query)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, errors.New("Gользователь не найден")
		}
		return 0, errors.New(fmt.Sprintf("Ошибка запроса: %s \n Ошибка: %v", query, err))
	}
	return balance, nil
}

func UpdateAmount(ctx context.Context, db *sqlx.DB, user_id int, amount float64) (float64, error) {
	var out float64
	query := fmt.Sprintf("UPDATE users SET balance = balance - %f WHERE id = %d AND balance >= %f RETURNING balance;", amount, user_id, amount)
	err := db.GetContext(ctx, &out, query)
	if err != nil {
		return out, errors.New(fmt.Sprintf("Ошибка запроса: %s \n Ошибка: %v", query, err))
	}

	return out, nil
}
