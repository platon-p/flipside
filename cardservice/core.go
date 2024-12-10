package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/platon-p/flipside/cardservice/api/controller"
	"github.com/platon-p/flipside/cardservice/api/middleware"
	"github.com/platon-p/flipside/cardservice/api/route"
	"github.com/platon-p/flipside/cardservice/repository"
	"github.com/platon-p/flipside/cardservice/service"
	"github.com/platon-p/flipside/cardservice/service/training"
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
		SignKey:    []byte(config.SignKey),
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
	userRepository := repository.NewUserRepositoryImpl(conn)
	trainingRepository := repository.NewTrainingRepositoryImpl(conn)

	basicTaskChecker := training.NewBasicTaskChecker(trainingRepository, cardRepository)
	checkers := []training.TaskChecker{basicTaskChecker}

	cardSetService := service.NewCardSetService(cardSetRepository)
	cardService := service.NewCardService(cardSetRepository, cardRepository)
	userService := service.NewUserService(userRepository)
	trainingService := training.NewTrainingService(trainingRepository, cardSetRepository, cardRepository, checkers)

	authMiddleware := middleware.NewAuthMiddleware(cfg.SignKey)

	trainingController := controller.NewTrainingController(trainingService)

	cardSetRouter := route.NewCardSetRouter(cardSetService, authMiddleware)
	cardRouter := route.NewCardRouter(cardService, authMiddleware)
	userRouter := route.NewUserRouter(userService)
	trainingRouter := route.NewTrainingRouter(trainingController, authMiddleware)

	router := route.NewRouter(cardSetRouter, cardRouter, userRouter, trainingRouter)
	engine := gin.Default()

	return &Core{
		engine: engine,
		router: router,
	}
}

func (c *Core) Start() {
	c.router.Setup(&c.engine.RouterGroup)
	err := c.engine.Run()
	if err != nil {
		panic(err)
	}
}
