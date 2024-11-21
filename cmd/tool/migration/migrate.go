package migration

import (
	"hub-service/config"
	"hub-service/container"
	"hub-service/logger"
	"hub-service/repository"
	"hub-service/session"
)

func Run() error {
	conf, env := config.LoadAppConfig()
	logger := logger.InitLogger(env)
	logger.GetZapLogger().Infof("Loaded this configuration : application." + env + ".yml")

	messages := config.LoadMessagesConfig()

	rep := repository.NewRepository(logger, conf)
	sess := session.NewSession(logger, conf)
	container := container.NewContainer(rep, sess, conf, messages, logger, env)

	err := CreateDatabase(container)
	if err != nil {
		logger.GetZapLogger().Error("Failed to create database", err.Error())
		return err
	}
	err = InitMasterData(container)
	if err != nil {
		logger.GetZapLogger().Error("Failed to init master data", err.Error())
		return err
	}

	logger.GetZapLogger().Infof("Migration completed.")
	return nil
}
