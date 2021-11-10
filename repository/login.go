package repository

import (
	"fmt"
	"payment-monitoring/models"

	"github.com/twinj/uuid"
	"gorm.io/gorm"
)

type LoginRepoStruct struct {
	db *gorm.DB
}

func NewLoginRepoImpl(db *gorm.DB) LoginRepoInterface {
	return LoginRepoStruct{db}
}

func (lr LoginRepoStruct) GetAdminId(name string) (models.Role, error) {
	role := models.Role{}
	err := lr.db.Debug().Where("role = ?", name).First(&role).Error
	if err != nil {
		return models.Role{}, err
	}
	return role, nil
}

func (lr LoginRepoStruct) GetDataWorkUnit(username string, password string) (models.WorkUnitAccount, error) {
	customer := models.WorkUnitAccount{}
	err := lr.db.Debug().Where("username = ? and password = ?", username, password).First(&customer).Error
	if err != nil {
		return models.WorkUnitAccount{}, err
	}
	return customer, nil
}

func (lr LoginRepoStruct) GetDataOfficerAccount(username string, password string, loginAs string) (models.OfficerAccount, error) {
	officer := models.OfficerAccount{}
	err := lr.db.Debug().Where("username = ? and password = ? and role = ?", username, password, loginAs).First(&officer).Error
	if err != nil {
		return models.OfficerAccount{}, err
	}
	return officer, nil
}

func (lr LoginRepoStruct) CreateAuth(username string, roleId string, roleName string, name string) (*models.Auth, error) {
	au := &models.Auth{}
	tx := lr.db.Begin()

	au.AuthUUID = uuid.NewV4().String()
	au.Username = username
	au.RoleID = roleId
	au.RoleName = roleName
	au.Name = name
	err := lr.db.Create(&au).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	fmt.Println("Insert data to database success")
	return au, nil
}

func (lr LoginRepoStruct) DeleteAuth(uuid string) error {
	au := models.Auth{}
	err := lr.db.Debug().Where("auth_uuid = ?", uuid).Delete(&au).Error
	if err != nil {
		return err
	}

	return nil
}
