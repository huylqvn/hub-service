package migration

import (
	"hub-service/container"
	"hub-service/model"
)

// InitMasterData creates the master data used in this application.
func InitMasterData(container container.Container) (err error) {
	if container.GetConfig().Extension.MasterGenerator {
		rep := container.GetRepository()

		r := model.NewRole("Admin")
		_, err = r.Create(rep)
		if err != nil {
			return err
		}

		hub := model.NewHub("hubtest")
		_, err = hub.Create(rep)
		if err != nil {
			return err
		}

		t := model.NewTeam("teamtest", hub.ID)
		_, err = t.Create(rep)
		if err != nil {
			return err
		}

		a := model.NewUserWithPlainPassword("user1", "pw1", r.ID, t.ID)
		_, err = a.Create(rep)
		if err != nil {
			return err
		}
		a = model.NewUserWithPlainPassword("user2", "pw2", r.ID, t.ID)
		_, err = a.Create(rep)
		if err != nil {
			return err
		}
	}

	return nil
}
