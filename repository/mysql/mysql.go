package mysql

import (
	"bytes"
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/sapawarga/auth/model"
)

type Auth struct {
	conn *sqlx.DB
}

func NewAuth(conn *sqlx.DB) *Auth {
	return &Auth{
		conn: conn,
	}
}

func (r *Auth) GetActorCurrentLoginByUsername(ctx context.Context, username string) (*model.Actor, error) {
	user, err := r.getUserByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	return &model.Actor{
		ID:        user.ID,
		Username:  user.Username,
		RoleLabel: model.RoleLabel[user.Role],
	}, nil
}

func (r *Auth) GetActorDetailByUsername(ctx context.Context, username string) (*model.UserDetail, error) {
	var err error
	user, err := r.getUserByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	mapLocations, err := r.getDetailLocationOfUser(ctx, user)
	if err != nil {
		return nil, err
	}

	job := &model.Job{}
	if user.JobTypeID.Valid {
		job, err = r.getJobTypeByID(ctx, user.JobTypeID.Int64)
		if err != nil {
			return nil, err
		}
	}

	return &model.UserDetail{
		ID: user.ID, Username: user.Username, Name: user.Name,
		Email: user.Email.String, Phone: user.Phone.String, Address: user.Address.String,
		LastLoginAt: user.LastLoginAt.Time,
		RoleLabel:   model.RoleLabel[user.Role],
		RT:          user.RT.String, RW: user.RW.String,
		KelurahanID: mapLocations["kelurahan"].ID,
		Kelurahan:   mapLocations["kelurahan"].Name,
		KecamatanID: mapLocations["kecamatan"].ID,
		Kecamatan:   mapLocations["kecamatan"].Name,
		KabKotaID:   mapLocations["kabkota"].ID,
		KabKota:     mapLocations["kabkota"].Name,
		Latitude:    user.Latitude.String,
		Longitude:   user.Longitude.String,
		BirthDate:   user.BirthDate.Time,
		JobTypeID:   job.ID, JobType: job.Title,
		EducationLevelID:  user.EducationLevelID.Int64,
		EducationLevel:    model.EducationLevel[user.EducationLevelID.Int64],
		PhotoUrl:          user.PhotoUrl.String,
		Facebook:          user.Facebook.String,
		Twitter:           user.Twitter.String,
		Instagram:         user.Instagram.String,
		PasswordUpdatedAt: user.PasswordUpdatedAt.Time,
		ProfileUpdatedAt:  user.ProfileUpdatedAt.Time,
		LastAccessAt:      user.LastAccessAt.Time,
	}, nil
}

func (r *Auth) getDetailLocationOfUser(ctx context.Context, user *model.User) (map[string]*model.Location, error) {
	var mapLocation = make(map[string]*model.Location)
	if user.KelurahanID.Valid {
		kelurahan, err := r.getLocationByID(ctx, user.KelurahanID.Int64)
		if err != nil {
			return nil, err
		}
		mapLocation["kelurahan"] = kelurahan
	}

	if user.KecamatanID.Valid {
		kecamatan, err := r.getLocationByID(ctx, user.KecamatanID.Int64)
		if err != nil {
			return nil, err
		}
		mapLocation["kecamatan"] = kecamatan
	}

	if user.KabKotaID.Valid {
		kabKota, err := r.getLocationByID(ctx, user.KabKotaID.Int64)
		if err != nil {
			return nil, err
		}
		mapLocation["kabkota"] = kabKota
	}

	return mapLocation, nil
}

func (r *Auth) getUserByUsername(ctx context.Context, username string) (*model.User, error) {
	var query bytes.Buffer
	var err error
	var result = &model.User{}

	query.WriteString(` SELECT id, username, email,	 FROM_UNIXTIME(last_login_at) as last_login_at, role, name, phone, address, job_type_id, education_level_id, 
		birth_date,	rt, rw, kel_id, kec_id, kabkota_id, lat, lon, photo_url, facebook, twitter, instagram, FROM_UNIXTIME(password_updated_at) as password_updated_at, 
		FROM_UNIXTIME(profile_updated_at) as profile_updated_at, FROM_UNIXTIME(last_access_at) as last_access_at
		FROM user
		WHERE username = ? AND status = ?`)

	if ctx != nil {
		err = r.conn.GetContext(ctx, result, query.String(), username, 10)
	} else {
		err = r.conn.Get(result, query.String(), username, 10)
	}

	if err != sql.ErrNoRows || err != nil {
		return nil, err
	}

	return result, nil
}

func (r *Auth) getLocationByID(ctx context.Context, id int64) (*model.Location, error) {
	var query bytes.Buffer
	var err error
	var result = &model.Location{}

	query.WriteString(` SELECT id, name FROM areas WHERE id = ? and status = 1`)

	if ctx != nil {
		err = r.conn.GetContext(ctx, result, query.String(), id)
	} else {
		err = r.conn.Get(result, query.String(), id)
	}

	if err != sql.ErrNoRows || err != nil {
		return nil, err
	}

	return result, nil
}

func (r *Auth) getJobTypeByID(ctx context.Context, id int64) (*model.Job, error) {
	var query bytes.Buffer
	var err error
	var result = &model.Job{}

	query.WriteString(`SELECT id, title FROM job_types WHERE id = ? AND status = 10`)

	if ctx != nil {
		err = r.conn.GetContext(ctx, result, query.String(), id)
	} else {
		err = r.conn.Get(result, query.String(), id)
	}

	if err != sql.ErrNoRows || err != nil {
		return nil, err
	}

	return result, nil
}
