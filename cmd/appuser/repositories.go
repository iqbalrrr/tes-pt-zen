package appuser

import (
	"database/sql"
	"tesptzen/utilitys"
	"time"
)

type IRepository interface {
	// RegisterNewUser(appuser ApplicationUser, roles []string) (string, error)
	Insert(applicationuser ApplicationUser) (ApplicationUser, error)
	Update(appuser ApplicationUser) (ApplicationUser, error)
	Delete(appuser ApplicationUser) (ApplicationUser, error)
	// SyncUserToRole(applicationUser ApplicationUser, roles []string) (ApplicationUser, error)
	// FindByName(username string) (ApplicationUser, error)
	FindByFieldName(fieldName string, value string) ([]ApplicationUser, error)
	// FindAll() ([]ApplicationUser, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *repository {
	return &repository{db}
}

func (r *repository) Insert(applicationuser ApplicationUser) (ApplicationUser, error) {
	applicationuser.CreatedBy = "admin"
	applicationuser.CreatedOn = time.Now()
	querys := "insert into applicationuser(username,email,passwordhash,phonenumber,lockout,status,createdby,createdon,lastupdatedby,lastupdatedon,deletedby,deletedon,rowstatus) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13)"
	_, err := r.db.Exec(querys,
		applicationuser.Username,
		applicationuser.Email,
		applicationuser.PasswordHash,
		applicationuser.PhoneNumber,
		applicationuser.Lockout,
		applicationuser.Status,
		applicationuser.CreatedBy,
		applicationuser.CreatedOn,
		applicationuser.LastUpdatedBy,
		applicationuser.LastUpdatedOn,
		applicationuser.DeletedBy,
		applicationuser.DeletedOn,
		applicationuser.RowStatus,
	)

	if err != nil {
		utilitys.LogError(err)
		return applicationuser, err
	}

	return applicationuser, nil
}

func (r *repository) Update(appuser ApplicationUser) (ApplicationUser, error) {
	appuser.LastUpdatedBy = "admin"
	appuser.LastUpdatedOn = time.Now()
	querys := "update applicationuser set email = $1,phonenumber=$2, passwordhash = $3,LastUpdatedBy = $4, LastUpdatedOn = $5 where Id = $6"
	_, err := r.db.Exec(querys,
		appuser.Email,
		appuser.PhoneNumber,
		appuser.PasswordHash,
		appuser.LastUpdatedBy,
		appuser.LastUpdatedOn,
		appuser.Id,
	)

	if err != nil {
		return appuser, err
	}

	return appuser, nil
}

func (r *repository) Delete(appuser ApplicationUser) (ApplicationUser, error) {
	appuser.DeletedBy = "admin"
	appuser.DeletedOn = time.Now()
	appuser.RowStatus = 0
	querys := "update ApplicationUser set RowStatus = ?,DeletedBy=?, DeletedOn = ? where Id = ?"
	_, err := r.db.Exec(querys,
		appuser.RowStatus,
		appuser.DeletedBy,
		appuser.DeletedOn,
		appuser.Id,
	)
	if err != nil {
		return appuser, err
	}
	return appuser, nil
}

func (r *repository) FindByFieldName(fieldName string, value string) ([]ApplicationUser, error) {
	var applicationuser ApplicationUser
	var applicationusers []ApplicationUser
	rows, err := r.db.Query("select * from applicationuser where username = $1", value)
	if err != nil {
		utilitys.LogError(err)
		return applicationusers, err
	}
	defer rows.Close()
	for rows.Next() {

		err2 := rows.Scan(&applicationuser.Id,
			&applicationuser.Username,
			&applicationuser.Email,
			&applicationuser.PasswordHash,
			&applicationuser.PhoneNumber,
			&applicationuser.Lockout,
			&applicationuser.Status,
			&applicationuser.CreatedBy,
			&applicationuser.CreatedOn,
			&applicationuser.LastUpdatedBy,
			&applicationuser.LastUpdatedOn,
			&applicationuser.DeletedBy,
			&applicationuser.DeletedOn,
			&applicationuser.RowStatus,
		)
		if err2 != nil {
			return applicationusers, err2
		}
		applicationusers = append(applicationusers, applicationuser)

	}

	return applicationusers, nil
}
