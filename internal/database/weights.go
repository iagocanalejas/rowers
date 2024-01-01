package database

type UserWeight struct {
	UserId       int64   `db:"user_id" json:"user_id"`
	Weight       float64 `db:"weight" json:"weight" form:"weight"`
	CreationDate string  `db:"creation_date" json:"creation_date" form:"creation_date"`
}
