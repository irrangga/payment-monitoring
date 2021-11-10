package repository

import (
	"payment-monitoring/models"

	"gorm.io/gorm"
)

type AccountRepoStruct struct {
	db *gorm.DB
}

func NewAccountRepoImpl(db *gorm.DB) AccountRepoInterface {
	return &AccountRepoStruct{db}
}

func (ar AccountRepoStruct) AddAccountRepo(accOfficer models.OfficerAccount) (models.OfficerAccount, models.WorkUnitAccount, error) {
	accCustomer := models.WorkUnitAccount{
		UserID:   accOfficer.UserID,
		Name:     accOfficer.Name,
		Username: accOfficer.Username,
		Password: accOfficer.Password,
	}

	if accOfficer.Role == 4 {
		accOfficer.Role = 0
	}

	if accOfficer.Role == 0 {
		err := ar.db.Create(&accCustomer).Error

		if err != nil {
			return models.OfficerAccount{}, models.WorkUnitAccount{}, err
		}
		return models.OfficerAccount{}, accCustomer, err
	} else {
		err := ar.db.Debug().Create(&accOfficer).Error
		if err != nil {
			return models.OfficerAccount{}, models.WorkUnitAccount{}, err
		}
		return accOfficer, models.WorkUnitAccount{}, err
	}
}

func (ar AccountRepoStruct) GetAllOfficerAccountRepo() ([]models.OfficerAccount, error) {
	var account []models.OfficerAccount
	err := ar.db.Debug().Find(&account).Error
	if err != nil {
		return nil, err
	}
	return account, nil
}

func (ar AccountRepoStruct) GetAllWorkUnitAccountRepo() ([]models.WorkUnitAccount, error) {
	var account []models.WorkUnitAccount
	err := ar.db.Debug().Find(&account).Error
	if err != nil {
		return nil, err
	}
	return account, nil
}

func (ar AccountRepoStruct) GetAccountOfficerRepo(accOfficer models.OfficerAccount, id string) models.OfficerAccount {
	err := ar.db.Debug().Where("user_id = ?", id).First(&accOfficer).Error
	if err != nil {
		return models.OfficerAccount{}
	}

	return accOfficer
}

func (ar AccountRepoStruct) GetAccountWorkUnitRepo(accCustomer models.WorkUnitAccount, id string) models.WorkUnitAccount {
	err := ar.db.Debug().Where("user_id = ?", id).First(&accCustomer).Error
	if err != nil {
		return models.WorkUnitAccount{}
	}

	return accCustomer
}

func (ar AccountRepoStruct) UpdateAccountRepo(accOfficer models.OfficerAccount, id string) (models.OfficerAccount, models.WorkUnitAccount, error) {
	accCustomer := models.WorkUnitAccount{
		UserID:   accOfficer.UserID,
		Name:     accOfficer.Name,
		Username: accOfficer.Username,
		Password: accOfficer.Password,
	}

	if accOfficer.Role == 0 {
		err := ar.db.Debug().Model(&accCustomer).Where("user_id = ?", id).Updates(&accCustomer).Error
		if err != nil {
			return models.OfficerAccount{}, models.WorkUnitAccount{}, err
		}
		return models.OfficerAccount{}, accCustomer, err
	} else {
		err := ar.db.Debug().Model(&accOfficer).Where("user_id = ?", id).Updates(&accOfficer).Error
		if err != nil {
			return models.OfficerAccount{}, models.WorkUnitAccount{}, err
		}
		return accOfficer, models.WorkUnitAccount{}, err
	}
}

func (ar AccountRepoStruct) DeleteAccountOfficerRepo(account models.OfficerAccount, id string) models.OfficerAccount {
	err := ar.db.Debug().Where("user_id = ?", id).Delete(&account).Error
	if err != nil {
		return models.OfficerAccount{}
	}

	return account
}

func (ar AccountRepoStruct) DeleteAccountWorkUnitRepo(account models.WorkUnitAccount, id string) models.WorkUnitAccount {
	err := ar.db.Debug().Where("user_id = ?", id).Delete(&account).Error
	if err != nil {
		return models.WorkUnitAccount{}
	}

	return account
}

func (ar AccountRepoStruct) GetAllAccountOfficer(username string) (int, error) {
	var count int64
	var tbl models.OfficerAccount
	err := ar.db.Debug().Model(tbl).Where("username = ?", username).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return int(count), nil
}

func (ar AccountRepoStruct) GetAllAccountWorkUnit(username string) (int, error) {
	var count int64
	var tbl models.WorkUnitAccount
	err := ar.db.Debug().Model(tbl).Where("username = ?", username).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return int(count), nil
}

func (ar AccountRepoStruct) AddUserIdRepo(userID models.User) models.User {
	err := ar.db.Debug().Create(&userID).Error
	if err != nil {
		return models.User{}
	}

	return userID
}
