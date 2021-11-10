package repository

import (
	"payment-monitoring/models"

	"gorm.io/gorm"
)

type RoleRepoStruct struct {
	db *gorm.DB
}

func NewRoleRepoImpl(db *gorm.DB) RoleRepoInterface {
	return &RoleRepoStruct{db: db}
}

func (rr RoleRepoStruct) AddRoleRepo(role models.Role) models.Role {
	trx := rr.db.Begin()
	err := trx.Debug().Create(&role).Error
	if err != nil {
		trx.Rollback()
		return models.Role{}
	}

	trx.Commit()

	return role
}

func (rr RoleRepoStruct) GetAllRoleRepo() ([]models.Role, error) {
	var role []models.Role
	err := rr.db.Debug().Find(&role).Error
	if err != nil {
		return nil, err
	}
	return role, nil
}

func (rr RoleRepoStruct) GetRoleRepo(role models.Role, id string) models.Role {
	err := rr.db.Debug().Where("id = ?", id).First(&role).Error
	if err != nil {
		return models.Role{}
	}

	return role
}

func (rr *RoleRepoStruct) GetRoleRepoByName(name string) models.Role {
	var role models.Role
	err := rr.db.Debug().Where("role_name = ?", name).First(&role).Error
	if err != nil {
		return models.Role{}
	}

	return role
}

func (rr RoleRepoStruct) UpdateRoleRepo(role models.Role, id string) (models.Role, error) {
	err := rr.db.Debug().Model(&role).Where("id = ?", id).Updates(&role).Error
	if err != nil {
		return models.Role{}, err
	}

	return role, nil
}

func (rr RoleRepoStruct) DeleteRoleRepo(role models.Role, id string) models.Role {
	err := rr.db.Debug().Where("id = ?", id).Delete(&role).Error
	if err != nil {
		return models.Role{}
	}

	return role
}
