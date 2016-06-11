package db

type User struct {
	Uid               uint32
	Uname             string `db:"username"`
	Password          string
	Coin, Score, Icon uint32
	VipLevel          uint32 `db:"vip_level"`
	MaxDayWin         uint32 `db:"max_day_win"`
	MaxSingleWin      uint32 `db:"max_single_win"`
	MaxSingleLose     uint32 `db:"max_single_lose"`
	BestCards         string `db:"best_cards"`
}

func GetUserByUid(id int) (user User, err error) {
	uid := uint32(id)
	err = db.Get(&user, "SELECT * FROM user WHERE uid=?", uid)
	return
}

func GetUserByUidPassword(id int, password string) (user User, status int, err error) {
	user = User{}
	user, err = GetUserByUid(id)
	if err != nil {
		return
	}
	if user.Password == password {
		status = 1
	} else {
		status = -1
	}
	return
}
