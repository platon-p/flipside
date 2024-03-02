package main

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/platon-p/flashside/authservice/api/controller"
	"github.com/platon-p/flashside/authservice/api/route"
	"github.com/platon-p/flashside/authservice/repository"
	"github.com/platon-p/flashside/authservice/service"
	"github.com/platon-p/flashside/authservice/utils"
)

type Core struct {
	router *route.Router
}

type CoreConfig struct {
	dataSource            string
	jwtSignKey            string
	jwtExpiresIn          time.Duration
	refreshTokenExpiresIn time.Duration
}

type ConfigEnv struct {
	DataSource            string `env:"DATASOURCE"`
	JwtSignKey            string `env:"JWT_SIGN_KEY"`
	JwtExpiresIn          string `env:"JWT_EXPIRES_IN"`
	RefreshTokenExpiresIn string `env:"REFRESH_TOKEN_EXPIRES_IN"`
}

func LoadConfig() CoreConfig {
	var cfg ConfigEnv
	cleanenv.ReadEnv(&cfg)
    jwtExpiresIn, err := time.ParseDuration(cfg.JwtExpiresIn)
    if err != nil {
        panic(err)
    }
    refreshTokenExpiresIn, err := time.ParseDuration(cfg.RefreshTokenExpiresIn)
    if err != nil {
        panic(err)
    }
    config := CoreConfig{
    	dataSource:            cfg.DataSource,
    	jwtSignKey:            cfg.JwtSignKey,
    	jwtExpiresIn:          jwtExpiresIn,
    	refreshTokenExpiresIn: refreshTokenExpiresIn,
    }
    return config
}

func NewCore() *Core {
    config := LoadConfig()
	jwtUtility := utils.NewJwtUtility(config.jwtSignKey, config.jwtExpiresIn)
	passwordUtility := utils.NewPasswordUtility()

	conn, err := repository.NewPostgresConnection(config.dataSource)
    if err != nil {
        panic(err)
    }
	userRepository := repository.NewUserRepositoryImpl(conn)
	refreshTokenRepository := repository.NewRefreshTokenRepositoryPostgres(conn)

	refreshTokenService := service.NewRefreshTokenService(refreshTokenRepository, config.refreshTokenExpiresIn)
	authService := service.NewAuthService(jwtUtility, passwordUtility, userRepository, refreshTokenService)
	checkService := service.NewCheckService(userRepository)

	authController := controller.NewAuthController(authService, checkService)
	checkController := controller.NewCheckController(checkService)

	authRouter := route.NewAuthRouter(authController)
	checkRouter := route.NewCheckRouter(checkController)
	router := route.NewRouter(authRouter, checkRouter)
	return &Core{
		router: router,
	}
}

func (c *Core) Start() {
	r := gin.Default()
	c.router.Setup(&r.RouterGroup)
	r.Run()
}
