package repositories

type ActionHistory struct {
	Id             int     `db:"id"`
	UserId         int     `db:"user_id"`
	Username       string  `db:"username"`
	Amount         float64 `db:"amount"`
	Currentbalance float64 `db:"currentbalance"`
	ActionEng      string  `db:"action_eng"`
	ActionRus      string  `db:"action_rus"`
	Ts             string  `db:"ts"`
}
