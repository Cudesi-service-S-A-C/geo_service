package connection

import (
	"log"

	"github.com/Cudesi-service-S-A-C/geo_service/service/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DSN = utils.ReadEnv("PGSQL_DSN")
var DB *gorm.DB

func Connection() {
	var error error
	DB, error = gorm.Open(postgres.Open(DSN), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "geo_",
			NoLowerCase:   false,
			SingularTable: false,
		},
	})

	if error != nil {
		log.Fatal(error)
	} else {
		log.Println("[Geo Service] - DB Connected")
	}
}
