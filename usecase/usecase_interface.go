package usecase

import "payment-monitoring/models"

type RoleUsecaseInterface interface {
	GetAllRoleUsecase() ([]models.Role, error)
	AddRoleUsecase(role models.Role) (models.Role, error)
	GetRoleUsecase(role models.Role, id string) models.Role
	UpdateRoleUsecase(role models.Role, id string) (models.Role, error)
	DeleteRoleUsecase(role models.Role, id string) models.Role
}

type AccountUsecaseInterface interface {
	AddAccountUsecase(accOfficer models.OfficerAccount) (models.OfficerAccount, models.WorkUnitAccount, error)
	GetAllOfficerAccountUsecase() ([]models.OfficerAccount, error)
	GetAllWorkUnitAccountUsecase() ([]models.WorkUnitAccount, error)
	GetAccountUsecase(accOfficer models.OfficerAccount, accCustomer models.WorkUnitAccount, id string) (models.OfficerAccount, models.WorkUnitAccount)
	UpdateAccountUsecase(accOfficer models.OfficerAccount, accCustomer models.WorkUnitAccount, id string) (models.OfficerAccount, models.WorkUnitAccount, error)
	DeleteAccountUsecase(accOfficer models.OfficerAccount, accCustomer models.WorkUnitAccount, id string) (models.OfficerAccount, models.WorkUnitAccount)
	AddUserIdUsecase(userID models.User) models.User
}

type LoginUsecaseInterface interface {
	GetDataWorkUnit(username string, password string) (models.WorkUnitAccount, error)
	CreateAuth(user models.OfficerAccount) (*models.Auth, error)
	SignIn(authD models.Auth) (string, error)
	DeleteAuth(uuid string) error
}

type PaymentUsecaseInterface interface {
	CreatePaymentUsecase(input models.InputCreatePayment) (models.EntityPayment, error)
	GetReqByIdUnitUsecase(id string) ([]models.EntityPayment, error)
	GetReqByIdPaymentUsecase(id string) (models.EntityPayment, error)
	UpdateAprovalUsecase(aproval models.EntityPayment, tipeReqRes string) (models.EntityPayment, error)
}
