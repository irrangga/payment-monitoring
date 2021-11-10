package usecase

import (
	"payment-monitoring/auth"
	"payment-monitoring/models"
	"payment-monitoring/repository"
	"strconv"
)

type LoginUsecaseStruct struct {
	login repository.LoginRepoInterface
}

func NewLoginUsecaseImpl(login repository.LoginRepoInterface) LoginUsecaseInterface {
	return &LoginUsecaseStruct{login}
}

func (lu LoginUsecaseStruct) GetDataWorkUnit(username string, password string) (models.WorkUnitAccount, error) {
	res, err := lu.login.GetDataWorkUnit(username, password)
	if err != nil {
		return models.WorkUnitAccount{}, err
	}
	return res, nil
}

func (lu LoginUsecaseStruct) CreateAuth(user models.OfficerAccount) (*models.Auth, error) {
	var CustomerID int
	var idOffice int

	if user.LoginAs == "4" {
		user.LoginAs = ""
	}

	if user.LoginAs == "" {
		customer, err := lu.login.GetDataWorkUnit(user.Username, user.Password)
		if err != nil {
			return nil, err
		}
		CustomerID = int(customer.ID)
		user.LoginAs = "customer"
		user.Name = customer.Name
		user.Role = 4
	} else {
		account, err := lu.login.GetDataOfficerAccount(user.Username, user.Password, user.LoginAs)
		if err != nil {
			return nil, err
		}
		idOffice = int(account.ID)
		user.LoginAs = "officer account"
		user.Name = account.Name
		user.Role = account.Role
	}

	roleData := strconv.Itoa(user.Role)
	dataAuth, err := lu.login.CreateAuth(user.Username, roleData, user.LoginAs, user.Name)
	if err != nil {
		return nil, err
	}

	dataAuth.CustomerID = CustomerID
	dataAuth.OfficerID = idOffice
	return dataAuth, nil
}

func (lu LoginUsecaseStruct) SignIn(authD models.Auth) (string, error) {
	token, err := auth.CreateToken(authD)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (lu LoginUsecaseStruct) DeleteAuth(uuid string) error {
	return lu.login.DeleteAuth(uuid)
}
