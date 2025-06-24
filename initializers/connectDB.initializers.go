package initializers

import (
	"fmt"
	"log"

	spanner "github.com/googleapis/go-gorm-spanner"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type Config struct {
	ProjectID  string `mapstructure:"PROJECT_ID"`
	InstanceID string `mapstructure:"INSTANCE_ID"`
	DatabaseID string `mapstructure:"DATABASE_ID"`
}

var DB *gorm.DB

func ConnectDB() (err error) {
	config := Config{
		ProjectID:  "lively-tensor-456313-b3",
		InstanceID: "product-catalog",
		DatabaseID: "product-catalog",
	}

	dsn := fmt.Sprintf("projects/%s/instances/%s/databases/%s", config.ProjectID, config.InstanceID, config.DatabaseID)

	DB, err = gorm.Open(spanner.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		log.Fatalf("Failed to connect to Spanner DB: %v", err)
		return
	}
	fmt.Println("Connected to Spanner DB successfully!")
	return
}
