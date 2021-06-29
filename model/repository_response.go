package model

import (
	"database/sql"
	"time"
)

type UserDetail struct {
	ID                int64     `json:"id"`
	Username          string    `json:"username"`
	Email             string    `json:"email,omitempty"`
	LastLoginAt       time.Time `json:"last_login_at,omitempty"`
	RoleLabel         string    `json:"role_label"`
	Name              string    `json:"name"`
	Phone             string    `json:"phone,omitempty"`
	Address           string    `json:"address,omitempty"`
	JobTypeID         int64     `json:"job_type_id,omitempty"`
	JobType           string    `json:"job_type,omitempty"`
	EducationLevelID  int64     `json:"education_level_id,omitempty"`
	EducationLevel    string    `json:"education_level,omitempty"`
	BirthDate         time.Time `json:"birth_date,omitempty"`
	RT                string    `json:"rt,omitempty"`
	RW                string    `json:"rw,omitempty"`
	KelurahanID       int64     `json:"kel_id,omitempty"`
	Kelurahan         string    `json:"kelurahan,omitempty"`
	KecamatanID       int64     `json:"kec_id,omitempty"`
	Kecamatan         string    `json:"kecamatan,omitempty"`
	KabKotaID         int64     `json:"kabkota_id,omitempty"`
	KabKota           string    `json:"kabkota,omitempty"`
	Latitude          string    `json:"lat,omitempty"`
	Longitude         string    `json:"long,omitempty"`
	PhotoUrl          string    `json:"photo_url,omitempty"`
	Facebook          string    `json:"facebook,omitempty"`
	Twitter           string    `json:"twitter,omitempty"`
	Instagram         string    `json:"instagram,omitempty"`
	PasswordUpdatedAt time.Time `json:"password_updated_at,omitempty"`
	ProfileUpdatedAt  time.Time `json:"profile_updated_at,omitempty"`
	LastAccessAt      time.Time `json:"last_access_at,omitempty"`
}

type User struct {
	ID                int64          `db:"id"`
	Username          string         `db:"username"`
	Email             sql.NullString `db:"email"`
	LastLoginAt       sql.NullTime   `db:"last_login_at"`
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
	PasswordUpdatedAt sql.NullTime   `db:"password_updated_at"`
	ProfileUpdatedAt  sql.NullTime   `db:"profile_updated_at"`
	LastAccessAt      sql.NullTime   `db:"last_access_at"`
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
