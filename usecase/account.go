package usecase

import (
	"errors"
	"payment-monitoring/models"
	"payment-monitoring/repository"

	gtr_validate "github.com/fikrigatrh/validate"
)

type AccountUsecaseStruct struct {
	account repository.AccountRepoInterface
}

func NewAccountUsecaseImpl(account repository.AccountRepoInterface) AccountUsecaseInterface {
	return &AccountUsecaseStruct{account}
}

func (au AccountUsecaseStruct) AddAccountUsecase(accOfficer models.OfficerAccount) (models.OfficerAccount, models.WorkUnitAccount, error) {
	check := gtr_validate.Num(accOfficer.Username)
	if check {
		return models.OfficerAccount{}, models.WorkUnitAccount{}, errors.New("username must be number")
	}

	if accOfficer.Role != models.Admin {
		if accOfficer.Role != models.GeneralSupport {
			if accOfficer.Role != models.Accounting {
				if accOfficer.Role != models.Customer {
					return models.OfficerAccount{}, models.WorkUnitAccount{}, errors.New("role not available")
				}
			}
		}
	}

	if accOfficer.Role == 4 {
		count, err := au.account.GetAllAccountWorkUnit(accOfficer.Username)
		if err != nil {
			return models.OfficerAccount{}, models.WorkUnitAccount{}, err
		} else if count > 0 {
			return models.OfficerAccount{}, models.WorkUnitAccount{}, errors.New("username has been registered")
		}
	} else {
		count, err := au.account.GetAllAccountOfficer(accOfficer.Username)
		if err != nil {
			return models.OfficerAccount{}, models.WorkUnitAccount{}, err
		} else if count > 0 {
			return models.OfficerAccount{}, models.WorkUnitAccount{}, errors.New("username has been registered")
		}
	}

	accDataOfficer, accDataCustomer, err := au.account.AddAccountRepo(accOfficer)
	if err != nil {
		return models.OfficerAccount{}, models.WorkUnitAccount{}, err
	}

	return accDataOfficer, accDataCustomer, nil
}

func (au AccountUsecaseStruct) GetAllOfficerAccountUsecase() ([]models.OfficerAccount, error) {
	accData, err := au.account.GetAllOfficerAccountRepo()
	if err != nil {
		return nil, err
	}

	return accData, nil
}

func (au AccountUsecaseStruct) GetAllWorkUnitAccountUsecase() ([]models.WorkUnitAccount, error) {
	accData, err := au.account.GetAllWorkUnitAccountRepo()
	if err != nil {
		return nil, err
	}

	return accData, nil
}

func (au AccountUsecaseStruct) GetAccountUsecase(accOfficer models.OfficerAccount, accCustomer models.WorkUnitAccount, id string) (models.OfficerAccount, models.WorkUnitAccount) {
	accDataOfficer := au.account.GetAccountOfficerRepo(accOfficer, id)
	accDataCustomer := au.account.GetAccountWorkUnitRepo(accCustomer, id)

	return accDataOfficer, accDataCustomer
}

func (au AccountUsecaseStruct) UpdateAccountUsecase(accOfficer models.OfficerAccount, accCustomer models.WorkUnitAccount, id string) (models.OfficerAccount, models.WorkUnitAccount, error) {
	accDataOfficer := au.account.GetAccountOfficerRepo(accOfficer, id)
	accDataCustomer := au.account.GetAccountWorkUnitRepo(accCustomer, id)

	accDataOfficer, accDataCustomer, err := au.account.UpdateAccountRepo(accOfficer, id)

	if err != nil {
		return models.OfficerAccount{}, models.WorkUnitAccount{}, err
	}
	return accDataOfficer, accDataCustomer, nil
}

func (au AccountUsecaseStruct) DeleteAccountUsecase(accOfficer models.OfficerAccount, accCustomer models.WorkUnitAccount, id string) (models.OfficerAccount, models.WorkUnitAccount) {
	accDataOfficer := au.account.DeleteAccountOfficerRepo(accOfficer, id)
	accDataCustomer := au.account.DeleteAccountWorkUnitRepo(accCustomer, id)

	return accDataOfficer, accDataCustomer
}

func (au AccountUsecaseStruct) AddUserIdUsecase(userID models.User) models.User {
	roleData := au.account.AddUserIdRepo(userID)

	return roleData
}
