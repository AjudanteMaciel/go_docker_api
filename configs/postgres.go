package configs

import (
	"go_docker_api/schemas"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// order matters
// the first schema will be created first and so on
var SchemasList = []interface{}{
	&schemas.State{},
	&schemas.Role{},
	&schemas.Employee{},
}

func InitializePostgres() (*gorm.DB, error) {

	logger := GetLogger("postgres")

	dsn := "host=localhost user=postgres password=0000 dbname=godb port=5433 sslmode=disable TimeZone=America/Sao_Paulo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		logger.Errorf("Error initializing postgres: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(SchemasList...)
	if err != nil {
		logger.Errorf("Error migrating schemas: %v", err)
		return nil, err
	}

	return db, nil
}

func DropPostgres() error {

	logger := GetLogger("postgres")

	logger.Infof("Dropping all tables...")

	dsn := "host=localhost user=postgres password=0000 dbname=godb port=5433 sslmode=disable TimeZone=America/Sao_Paulo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Errorf("Error initializing postgres: %v", err)
		return err
	}

	err = db.Migrator().DropTable(SchemasList...)
	if err != nil {
		logger.Errorf("Error dropping tables: %v", err)
		return err
	}

	logger.Infof("All tables dropped successfully.")
	return nil
}
