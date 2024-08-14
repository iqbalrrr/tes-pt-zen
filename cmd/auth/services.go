package auth

import (
	"errors"
	"os"
	"sync"
	"tesptzen/cmd/appuser"
	"tesptzen/utilitys"
	"tesptzen/utilitys/token"

	"golang.org/x/crypto/bcrypt"
)

type IAuthServices interface {
	RegisterUser(input appuser.RegisterUserInputDto) (appuser.ApplicationUser, error)
	VerifyPassword(password, hashedPassword string) error
	LoginCheck(username string, password string) (string, error)
	ResetPassword(input appuser.ResetPasswordDto) (string, error)

	// FindApplicationuserByName(username string) (ApplicationUser, error)
	// ChangePassword(input ChangePasswordDto) (string, error)
	// DeleteUser(input DeletePasswordDto) error
	// UpdateUserRole(input UpdateUserDto) error
	// GetAllUser() ([]App_applicationusers, error)
}

type authServices struct {
	repository appuser.IRepository
}

func NewService(repository appuser.IRepository) *authServices {
	return &authServices{repository}
}

func (s *authServices) RegisterUser(input appuser.RegisterUserInputDto) (appuser.ApplicationUser, error) {
	var wg sync.WaitGroup
	appuserz := appuser.NewApplicationuser()
	//automapper.MapLoose(input, &Appuser)
	appuserz.Username = input.Username
	appuserz.Email = input.Email
	appuserz.PhoneNumber = input.PhoneNumber

	errResultChan := make(chan error)
	firstResultChan := make(chan []byte)
	wg.Add(1)
	go func() {
		defer wg.Done()
		firstResult, errfirstchan := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
		firstResultChan <- firstResult
		errResultChan <- errfirstchan
	}()

	//passwordhash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	passwordhash := <-firstResultChan
	err := <-errResultChan

	if err != nil {
		utilitys.LogError(err)
		return appuserz, err
	}
	appuserz.PasswordHash = string(passwordhash)

	SecondResultChan := make(chan []appuser.ApplicationUser)
	wg.Add(1)
	go func() {
		defer wg.Done()
		secResult, errfirstchan := s.repository.FindByFieldName("Username", appuserz.Username)
		SecondResultChan <- secResult
		errResultChan <- errfirstchan
	}()
	ads := <-SecondResultChan
	errx := <-errResultChan
	//ads, errx := s.repository.FindByFieldName("Username", appuserz.Username)
	if errx != nil {
		utilitys.LogError(errx)
		return appuserz, errx
	}
	if len(ads) != 0 {
		return appuserz, errors.New("User Sudah Tersedia")
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		_, errfirstchan := s.repository.Insert(appuserz)
		errResultChan <- errfirstchan
	}()
	//_, err2 := s.repository.Insert(appuserz)
	err2 := <-errResultChan
	if err2 != nil {
		utilitys.LogError(err2)
		return appuserz, err2
	}
	go func() {
		wg.Wait()
		close(firstResultChan)
		close(SecondResultChan)
		close(errResultChan)
	}()
	return appuserz, nil
}

func (u *authServices) VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (u *authServices) LoginCheck(username string, password string) (string, error) {
	var err error

	Appuser := appuser.NewApplicationuser()
	Appuser.Username = username
	ads, errx := u.repository.FindByFieldName("Username", Appuser.Username)
	if errx != nil {
		utilitys.LogError(errx)
		return "", errx
	}
	if len(ads) == 0 {
		return "User Not Found", err
	}
	appuserIsExist := ads[0]

	err = u.VerifyPassword(password, appuserIsExist.PasswordHash)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		utilitys.LogError(err)
		return "Wrong Password", err
	}

	token, err2 := token.GenerateToken(uint(appuserIsExist.Id))
	if err2 != nil {
		utilitys.LogError(err2)
		return "Error Generate Token", err2
	}
	return token, nil
}

func (u *authServices) ResetPassword(input appuser.ResetPasswordDto) (string, error) {
	var err error
	Appuser := appuser.NewApplicationuser()
	Appuser.Username = input.Username
	ads, errx := u.repository.FindByFieldName("Username", Appuser.Username)
	if errx != nil {
		return "", errx
	}
	if len(ads) == 0 {
		return "User Not Found", err
	}
	appuserIsExist := ads[0]
	passwordhash, err := bcrypt.GenerateFromPassword([]byte(os.Getenv("DEFAULT_PASSWORD")), bcrypt.MinCost)
	if err != nil {
		return "Error", err
	}
	appuserIsExist.PasswordHash = string(passwordhash)
	updatedpassworduser, err := u.repository.Update(appuserIsExist)
	if err != nil {
		return "Error", err
	}
	return updatedpassworduser.Username, nil
}
