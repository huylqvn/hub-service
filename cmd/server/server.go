package server

import (
	"github.com/labstack/echo/v4"

	"hub-service/config"
	"hub-service/container"
	"hub-service/logger"
	"hub-service/middleware"
	"hub-service/repository"
	"hub-service/router"
	"hub-service/session"
)

// @title hub-service API
// @version 1.5.1
// @description This is API specification for hub-service project.

// @license.name MIT
// @license.url https://opensource.org/licenses/mit-license.php

// @host localhost:8080
// @BasePath /api
func Init() error {
	e := echo.New()

	conf, env := config.LoadAppConfig()
	logger := logger.InitLogger(env)
	logger.GetZapLogger().Infof("Loaded this configuration : application." + env + ".yml")

	messages := config.LoadMessagesConfig()
	logger.GetZapLogger().Infof("Loaded messages.properties")

	rep := repository.NewRepository(logger, conf)
	sess := session.NewSession(logger, conf)
	container := container.NewContainer(rep, sess, conf, messages, logger, env)

	router.Init(e, container)
	middleware.InitLoggerMiddleware(e, container)
	middleware.InitSessionMiddleware(e, container)

	if err := e.Start(":8080"); err != nil {
		logger.GetZapLogger().Errorf(err.Error())
		return err
	}

	defer rep.Close()
	return nil
}
