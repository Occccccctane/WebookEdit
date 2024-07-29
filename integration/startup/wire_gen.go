// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package startup

import (
	"GinStart/Ioc"
	"GinStart/Repository"
	"GinStart/Repository/Cache"
	"GinStart/Repository/Dao"
	"GinStart/Service"
	"GinStart/Web"
	"github.com/gin-gonic/gin"
)

// Injectors from wire.go:

func InitWireServer() *gin.Engine {
	cmdable := Ioc.InitRedis()
	v := Ioc.InitMiddleWare(cmdable)
	db := Ioc.InitDB()
	userDao := Dao.NewUserDao(db)
	userCache := Cache.NewUserCache(cmdable)
	userRepository := Repository.NewCacheUserRepository(userDao, userCache)
	userService := Service.NewUserService(userRepository)
	codeCache := Cache.NewCodeCache(cmdable)
	codeRepository := Repository.NewCodeRepository(codeCache)
	service := Ioc.InitSMSService()
	codeService := Service.NewCodeService(codeRepository, service)
	userHandler := Handler.NewUserHandler(userService, codeService)
	engine := Ioc.InitWebServer(v, userHandler)
	return engine
}
