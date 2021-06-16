package container

import (
	"first-api/app/http/controllers"
	repoImpl "first-api/app/repository/impl"
	svcImpl "first-api/app/svc/impl"
	"first-api/infra/conn"
)

func Init(g interface{}) {
	db := conn.Db()
	userRepo := repoImpl.NewMySqlUserRepository(db)
	userSvc := svcImpl.NewUserService(userRepo)

	controllers.NewUserController(g, userSvc)
}
