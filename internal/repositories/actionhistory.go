package repositories

import (
	"context"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
)

func InsertHistory(ctx context.Context, db *sqlx.DB, userId, actionId int, amount float64) (int64, error) {
	var id int64
	query := fmt.Sprintf("INSERT INTO action_history (user_id, action, amount) VALUES ('%d','%d', '%v') RETURNING  id", userId, actionId, amount)
	err := db.GetContext(ctx, &id, query)
	if err != nil {
		return 0, errors.New(fmt.Sprintf("Ошибка запроса: %s \n Ошибка: %v", query, err))
	}
	return id, nil
}

func GetHistory(ctx context.Context, db *sqlx.DB, userId int) ([]ActionHistory, error) {
	var history []ActionHistory
	var postQuery string
	if userId > 0 {
		postQuery = fmt.Sprintf(" AND user_id = %d ", userId)
	}
	query := fmt.Sprintf("SELECT h.id, u.id as user_id, u.name as username, h.amount, u.balance as currentbalance, d.eng as action_eng, d.rus as action_rus, h.ts FROM users u,action_history h  LEFT JOIN dict_action d ON h.action=d.id WHERE u.id = h.user_id %s ORDER BY id DESC;", postQuery)
	err := db.SelectContext(ctx, &history, query)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Ошибка запроса: %s \n Ошибка: %v", query, err))
	}
	return history, nil
}
