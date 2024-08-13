package appuser

import "time"

type ApplicationUser struct {
	Id            int       `default:"0"`
	Username      string    `default:""`
	Email         string    `default:""`
	PasswordHash  string    `default:""`
	PhoneNumber   string    `default:""`
	Lockout       uint8     `default:"0"`
	Status        int16     `default:"0"`
	CreatedBy     string    `default:""`
	CreatedOn     time.Time `default:"1900:01:01"`
	LastUpdatedBy string    `default:""`
	LastUpdatedOn time.Time `default:"1900:01:01"`
	DeletedBy     string    `default:""`
	DeletedOn     time.Time `default:"1900:01:01"`
	RowStatus     int16     `default:"1"`
}

func NewApplicationuser() ApplicationUser {
	applicationuser := ApplicationUser{}
	applicationuser.Id = 0
	applicationuser.RowStatus = 1
	applicationuser.CreatedBy = ""
	applicationuser.CreatedOn = time.Date(1900, time.Month(1), 1, 0, 0, 0, 0, time.UTC)
	applicationuser.LastUpdatedBy = ""
	applicationuser.LastUpdatedOn = time.Date(1900, time.Month(1), 1, 0, 0, 0, 0, time.UTC)
	applicationuser.DeletedBy = ""
	applicationuser.DeletedOn = time.Date(1900, time.Month(1), 1, 0, 0, 0, 0, time.UTC)
	applicationuser.Username = ""
	applicationuser.Email = ""
	applicationuser.PasswordHash = ""
	applicationuser.PhoneNumber = ""
	applicationuser.Lockout = 0
	applicationuser.Status = 0

	return applicationuser
}
