package migration

import (
	"hub-service/container"
	"hub-service/model"
)

// CreateDatabase creates the tables used in this application.
func CreateDatabase(container container.Container) (err error) {
	if container.GetConfig().Database.Migration {
		db := container.GetRepository()

		err = db.DropTableIfExists(&model.User{})
		if err != nil {
			return err
		}

		err = db.DropTableIfExists(&model.Role{})
		if err != nil {
			return err
		}
		err = db.DropTableIfExists(&model.Team{})
		if err != nil {
			return err
		}
		err = db.DropTableIfExists(&model.Hub{})
		if err != nil {
			return err
		}

		err = db.AutoMigrate(&model.Hub{})
		if err != nil {
			return err
		}
		err = db.AutoMigrate(&model.Team{})
		if err != nil {
			return err
		}

		err = db.AutoMigrate(&model.Role{})
		if err != nil {
			return err
		}
		err = db.AutoMigrate(&model.User{})
		if err != nil {
			return err
		}

	}

	return nil
}
