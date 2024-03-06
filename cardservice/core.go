package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/platon-p/flipside/cardservice/api/controller"
	"github.com/platon-p/flipside/cardservice/api/middleware"
	"github.com/platon-p/flipside/cardservice/api/route"
	"github.com/platon-p/flipside/cardservice/repository"
	"github.com/platon-p/flipside/cardservice/service"
)

type Core struct {
	engine *gin.Engine
	router *route.Router
}

type Config struct {
	DataSource string `env:"DATASOURCE"`
	SignKey    string `env:"JWT_SIGN_KEY"`
}

type CoreConfig struct {
	DataSource string
	SignKey    []byte
}

func LoadConfig() *CoreConfig {
    var config Config
    if err := cleanenv.ReadEnv(&config); err != nil {
        panic(err)
    }
    return &CoreConfig{
        DataSource: config.DataSource,
        SignKey: []byte(config.SignKey),
    }

}

func NewCore() *Core {
	cfg := LoadConfig()
	conn, err := repository.NewConnection(cfg.DataSource)
    if err != nil {
        panic(err)
    }
	cardSetRepository := repository.NewCardSetRepositoryImpl(conn)
    cardRepository := repository.NewCardRepositoryImpl(conn)

	cardSetService := service.NewCardSetService(cardSetRepository)
    cardService := service.NewCardService(cardSetRepository, cardRepository)
	authMiddleware := middleware.NewAuthMiddleware(cfg.SignKey)

	cardSetController := controller.NewCardSetController(cardSetService)
	cardController := controller.NewCardController(cardService)

	cardSetRouter := route.NewCardSetRouter(cardSetController, authMiddleware)
	cardRouter := route.NewCardRouter(cardController, authMiddleware)

	router := route.NewRouter(cardSetRouter, cardRouter)
	engine := gin.Default()

	return &Core{
		engine: engine,
		router: router,
	}
}

func (c *Core) Start() {
	c.router.Setup(&c.engine.RouterGroup)
	c.engine.Run()
}
