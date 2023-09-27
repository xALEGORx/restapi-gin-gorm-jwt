package core

import "restapi/models"

func (api *Api) Migrate() {
	var migrateModels []interface{} = []interface{}{
		&models.User{},
	}

	if err := api.Db.AutoMigrate(migrateModels...); err != nil {
		api.Log.Errorf("failed auto migrate: %s", err.Error())
	} else {
		api.Log.Info("migration successful completed")
	}
}
