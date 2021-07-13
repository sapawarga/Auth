package model

import (
	"database/sql"
)

type User struct {
	ID                int64          `db:"id"`
	Username          string         `db:"username"`
	Email             sql.NullString `db:"email"`
	LastLoginAt       sql.NullInt64  `db:"last_login_at"`
	Role              int64          `db:"role"`
	Name              string         `db:"name"`
	Phone             sql.NullString `db:"phone"`
	Address           sql.NullString `db:"address"`
	JobTypeID         sql.NullInt64  `db:"job_type_id"`
	EducationLevelID  sql.NullInt64  `db:"education_level_id"`
	BirthDate         sql.NullTime   `db:"birth_date"`
	RT                sql.NullString `db:"rt"`
	RW                sql.NullString `db:"rw"`
	KelurahanID       sql.NullInt64  `db:"kel_id"`
	KecamatanID       sql.NullInt64  `db:"kec_id"`
	KabKotaID         sql.NullInt64  `db:"kabkota_id"`
	Latitude          sql.NullString `db:"lat"`
	Longitude         sql.NullString `db:"lon"`
	PhotoUrl          sql.NullString `db:"photo_url"`
	Facebook          sql.NullString `db:"facebook"`
	Twitter           sql.NullString `db:"twitter"`
	Instagram         sql.NullString `db:"instagram"`
	PasswordUpdatedAt sql.NullInt64  `db:"password_updated_at"`
	ProfileUpdatedAt  sql.NullInt64  `db:"profile_updated_at"`
	LastAccessAt      sql.NullInt64  `db:"last_access_at"`
}

type Location struct {
	ID   int64  `db:"id"`
	Name string `db:"name"`
}

type Job struct {
	ID    int64  `db:"id"`
	Title string `db:"title"`
}

var EducationLevel = map[int64]string{
	1:  "Tidak Ada",
	2:  "SD/MI",
	3:  "SMP/MTS",
	4:  "SMA/SMK",
	5:  "D1",
	6:  "D2",
	7:  "D3",
	8:  "S1",
	9:  "S2",
	10: "S3",
	11: "Lainnya",
}

var RoleLabel = map[int64]string{
	10:  "user",
	49:  "trainer",
	50:  "staffRW",
	60:  "staffKel",
	70:  "staffKec",
	80:  "staffKabKota",
	88:  "staffOPD",
	89:  "staffSaberhoax",
	90:  "staffProv",
	91:  "pimpinan",
	99:  "admin",
	100: "service_account",
}
