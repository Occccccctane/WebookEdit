//go:build wireinject

package startup

import (
	"GinStart/Ioc"
	"GinStart/Repository"
	"GinStart/Repository/Cache"
	"GinStart/Repository/Dao"
	"GinStart/Service"
	Handler "GinStart/Web"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func InitWireServer() *gin.Engine {
	wire.Build(
		//第三方依赖
		Ioc.InitDB, InitRedis,
		//数据库交互层
		Dao.NewUserDao,
		//缓存交互层
		Cache.NewUserCache, Cache.NewCodeCache,
		//存储层
		Repository.NewCacheUserRepository, Repository.NewCodeRepository,
		//服务注册
		//将短信服务再封装，以后需要更换业务不需要再更改构建的依赖
		Ioc.InitSMSService,
		Service.NewCodeService,
		Service.NewUserService,
		//Web管理
		Handler.NewUserHandler,
		//引擎，中间件
		Ioc.InitMiddleWare,
		Ioc.InitWebServer,
	)
	return gin.Default()
}