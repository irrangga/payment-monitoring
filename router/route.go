package router

import (
	"payment-monitoring/config/db"
	"payment-monitoring/controller"
	"payment-monitoring/middlewares"
	"payment-monitoring/repository"
	"payment-monitoring/usecase"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {

	router := gin.New()

	DB := db.DB

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.Use(middlewares.CORSMiddleware())

	newRouteRole := router.Group("")

	// CALL DEPENDENCY REPOSITORY IN ABOVE
	loginRepo := repository.NewLoginRepoImpl(DB)
	roleRepo := repository.NewRoleRepoImpl(DB)
	accountRepo := repository.NewAccountRepoImpl(DB)
	paymentRepo := repository.NewPaymentRepoImpl(DB)

	// CALL DEPENDENCY USE CASE IN ABOVE
	loginUsecase := usecase.NewLoginUsecaseImpl(loginRepo)
	roleUsecase := usecase.NewRoleUsecaseImpl(roleRepo)
	accountUsecase := usecase.NewAccountUsecaseImpl(accountRepo)
	paymentUsecase := usecase.NewPaymentUsecaseImpl(paymentRepo)

	// CALL DEPENDENCY CONTROLLER IN ABOVE
	controller.NewLoginControllerImpl(newRouteRole, loginUsecase)
	newRouteRole.Use(middlewares.TokenAuthMiddleware())
	controller.NewRoleControllerImpl(newRouteRole, roleUsecase)
	controller.NewAccountControllerImpl(newRouteRole, accountUsecase)
	controller.NewPaymentControllerImpl(newRouteRole, paymentUsecase)

	return router
}
