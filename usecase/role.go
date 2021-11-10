package usecase

import (
	"errors"
	"fmt"
	"payment-monitoring/models"
	"payment-monitoring/repository"
	"strings"
)

type RoleUsecaseStruct struct {
	role repository.RoleRepoInterface
}

func NewRoleUsecaseImpl(role repository.RoleRepoInterface) RoleUsecaseInterface {
	return &RoleUsecaseStruct{role: role}
}

func (u RoleUsecaseStruct) AddRoleUsecase(role models.Role) (models.Role, error) {
	role.RoleName = strings.ToLower(role.RoleName)

	checkRole := u.role.GetRoleRepoByName(role.RoleName)
	if checkRole.ID != 0 {
		return models.Role{}, errors.New("Error")
	}

	roleData := u.role.AddRoleRepo(role)

	return roleData, nil
}

func (u RoleUsecaseStruct) GetAllRoleUsecase() ([]models.Role, error) {
	getAll, err := u.role.GetAllRoleRepo()
	if err != nil {
		return nil, err
	}

	return getAll, nil
}

func (u RoleUsecaseStruct) GetRoleUsecase(role models.Role, id string) models.Role {
	roleData := u.role.GetRoleRepo(role, id)

	return roleData
}

func (u RoleUsecaseStruct) UpdateRoleUsecase(role models.Role, id string) (models.Role, error) {
	fmt.Println(role.RoleName)
	roleData := u.role.GetRoleRepo(role, id)
	if roleData.ID == 0 {
		return models.Role{}, errors.New("error")
	}

	roleData, err := u.role.UpdateRoleRepo(role, id)
	if err != nil {
		return models.Role{}, err
	}
	return roleData, nil
}

func (u RoleUsecaseStruct) DeleteRoleUsecase(role models.Role, id string) models.Role {
	roleData := u.role.DeleteRoleRepo(role, id)

	return roleData
}
