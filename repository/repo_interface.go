package repository

import (
	"payment-monitoring/models"
)

type RoleRepoInterface interface {
	GetAllRoleRepo() ([]models.Role, error)
	AddRoleRepo(role models.Role) models.Role
	GetRoleRepo(role models.Role, id string) models.Role
	UpdateRoleRepo(role models.Role, id string) (models.Role, error)
	DeleteRoleRepo(role models.Role, id string) models.Role
	GetRoleRepoByName(name string) models.Role
}

type AccountRepoInterface interface {
	AddAccountRepo(accOfficer models.OfficerAccount) (models.OfficerAccount, models.WorkUnitAccount, error)
	GetAllOfficerAccountRepo() ([]models.OfficerAccount, error)
	GetAllWorkUnitAccountRepo() ([]models.WorkUnitAccount, error)
	GetAccountOfficerRepo(accOfficer models.OfficerAccount, id string) models.OfficerAccount
	GetAccountWorkUnitRepo(accCustomer models.WorkUnitAccount, id string) models.WorkUnitAccount
	UpdateAccountRepo(accOfficer models.OfficerAccount, id string) (models.OfficerAccount, models.WorkUnitAccount, error)
	DeleteAccountOfficerRepo(role models.OfficerAccount, id string) models.OfficerAccount
	DeleteAccountWorkUnitRepo(role models.WorkUnitAccount, id string) models.WorkUnitAccount
	GetAllAccountOfficer(username string) (int, error)
	GetAllAccountWorkUnit(username string) (int, error)
	AddUserIdRepo(userID models.User) models.User
}

type LoginRepoInterface interface {
	GetAdminId(name string) (models.Role, error)
	GetDataWorkUnit(username string, password string) (models.WorkUnitAccount, error)
	CreateAuth(username string, roleId, roleName string, name string) (*models.Auth, error)
	DeleteAuth(uuid string) error
	GetDataOfficerAccount(username string, password string, loginAs string) (models.OfficerAccount, error)
}

type PaymentRepoInterface interface {
	CreatePaymentRepository(input models.EntityPayment) (models.EntityPayment, error)
	GetReqByIdUnitRepository(id string) ([]models.EntityPayment, error)
	GetReqByIdPaymentRepository(id int) (models.EntityPayment, error)
	UpdateAprovalRepository(approval models.EntityPayment) (models.EntityPayment, error)
}
