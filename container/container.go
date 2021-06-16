package container

import (
	"first-api/controllers"
	"first-api/infra/conn"
	repoImpl "first-api/repository/impl"
	svcImpl "first-api/svc/impl"
)

func Init(g interface{}) {
	db := conn.Db()
	userRepo := repoImpl.NewMySqlUserRepository(db)
	userSvc := svcImpl.NewUserService(userRepo)

	controllers.NewUserController(g, userSvc)
}
