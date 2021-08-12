package model

import "time"

// Actor ...
type Actor struct {
	ID          int64       `json:"id"`
	Username    string      `json:"username"`
	RoleLabel   string      `json:"role_label"`
	LastLoginAt interface{} `json:"last_login_at"`
}

// UserDetail ...
type UserDetail struct {
	ID                int64     `json:"id"`
	Username          string    `json:"username"`
	Email             string    `json:"email,omitempty"`
	LastLoginAt       int64     `json:"last_login_at,omitempty"`
	Role              int64     `json:"role"`
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
	VillageID         int64     `json:"kel_id,omitempty"`
	Village           string    `json:"kelurahan,omitempty"`
	DistrictID        int64     `json:"kec_id,omitempty"`
	District          string    `json:"kecamatan,omitempty"`
	RegencyID         int64     `json:"kabkota_id,omitempty"`
	Regency           string    `json:"kabkota,omitempty"`
	Latitude          string    `json:"lat,omitempty"`
	Longitude         string    `json:"long,omitempty"`
	PhotoUrl          string    `json:"photo_url,omitempty"`
	Facebook          string    `json:"facebook,omitempty"`
	Twitter           string    `json:"twitter,omitempty"`
	Instagram         string    `json:"instagram,omitempty"`
	PasswordUpdatedAt int64     `json:"password_updated_at,omitempty"`
	ProfileUpdatedAt  int64     `json:"profile_updated_at,omitempty"`
	LastAccessAt      int64     `json:"last_access_at,omitempty"`
}
